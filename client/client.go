package client

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"go.tradeforge.dev/alpaca/encoder"
	alpacaerrors "go.tradeforge.dev/alpaca/errors"
	"go.tradeforge.dev/alpaca/model"
	"go.tradeforge.dev/alpaca/sse"

	"github.com/go-resty/resty/v2"
)

const clientVersion = "v0.0.0"

const (
	DefaultClientTimeout = 10 * time.Second
)

// Client defines an HTTP client for the REST API.
type Client struct {
	HTTP *resty.Client

	encoder *encoder.Encoder
	logger  *slog.Logger
}

// New returns a new client with the specified API key and config.
func New(
	apiURL string,
	logger *slog.Logger,
) *Client {
	return newClient(apiURL, logger)
}

func newClient(
	apiURL string,
	logger *slog.Logger,
) *Client {
	c := resty.New()
	c.SetBaseURL(apiURL)
	c.SetHeader("User-Agent", fmt.Sprintf("Alpaca client/%v", clientVersion))
	c.SetHeader("Accept", "application/json")

	return &Client{
		HTTP:    c,
		encoder: encoder.New(),
		logger:  logger,
	}
}

func (c *Client) SetBasicAuth(apiKey, apiSecret string) *Client {
	c.HTTP.SetBasicAuth(apiKey, apiSecret)
	return c
}

func (c *Client) SetHeader(key, value string) *Client {
	c.HTTP.SetHeader(key, value)
	return c
}

// Call makes an API call based on the request params and options. The response is automatically unmarshaled.
func (c *Client) Call(ctx context.Context, method, path string, params, response any, opts ...model.RequestOption) error {
	uri, err := c.encoder.EncodeParams(path, params)
	if err != nil {
		return err
	}
	return c.CallURL(ctx, method, uri, response, opts...)
}

// CallURL makes an API call based on a request URI and options. The response is automatically unmarshaled.
func (c *Client) CallURL(ctx context.Context, method, uri string, response any, opts ...model.RequestOption) error {
	options := mergeOptions(opts...)

	c.HTTP.SetTimeout(DefaultClientTimeout)
	req := c.HTTP.R().SetContext(ctx)
	if options.Body != nil {
		b, err := json.Marshal(options.Body)
		if err != nil {
			return fmt.Errorf("failed to marshal body: %w", err)
		}
		req.SetBody(b)
	}
	req.SetQueryParamsFromValues(options.QueryParams)
	req.SetHeaderMultiValues(options.Headers)
	req.SetResult(response).SetError(&alpacaerrors.ResponseError{})
	req.SetHeader("Content-Type", "application/json")

	_, err := c.executeRequest(ctx, req, method, uri, options.Trace)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) executeRequest(
	_ context.Context,
	req *resty.Request,
	method string,
	uri string,
	trace bool,
) (*resty.Response, error) {
	res, err := req.Execute(method, uri)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	if res.IsError() {
		responseError := parseResponseError(res)
		if responseError != nil {
			c.logger.Error(
				"response error",
				slog.String("url", uri),
				slog.Int("status", responseError.StatusCode),
				slog.String("error message", responseError.Message),
			)
		} else {
			c.logger.Error(
				"response error",
				slog.String("url", uri),
				slog.Int("status", res.StatusCode()),
				slog.String("error message", res.Status()),
				slog.String("response", string(res.Body())),
			)
		}
		return res, responseError
	}

	if trace {
		sanitizedHeaders := req.Header
		for k := range sanitizedHeaders {
			if k == "Authorization" {
				sanitizedHeaders[k] = []string{"REDACTED"}
			}
		}
		c.logger.Debug(
			"request",
			slog.String("url", uri),
			slog.Any("request headers", sanitizedHeaders),
			slog.Any("response headers", res.Header()),
		)
	}
	return res, nil
}

type EventStreamHandler func(ctx context.Context, event *sse.Event) error

// Listen to an event data stream.
// This is a blocking call that will continue to read from the stream until the context is canceled
// or the watch is stopped.
//
// NOTE: The event reader should not be shared between multiple listeners, otherwise, there might be unexpected parsing results.
func (c *Client) Listen(ctx context.Context, path string, params any, handler EventStreamHandler, opts ...model.RequestOption) error {
	r, err := c.listenToSSE(ctx, path, params, opts...)
	if err != nil {
		return fmt.Errorf("initializing SSE stream: %w", err)
	}
	defer func() {
		if err := r.Close(); err != nil {
			c.logger.Error("closing stream", slog.Any("error", err))
		}
	}()

	evtChannel, errChannel := make(chan *sse.Event, 1), make(chan error, 1)
	go c.startReadingSSE(r, evtChannel, errChannel)
	for {
		select {
		case <-ctx.Done():
			c.logger.Error("context cancelled", slog.Any("error", ctx.Err()))
			return nil
		case err := <-errChannel:
			c.logger.Error("reading from stream", slog.Any("error", err))
			return err
		case event := <-evtChannel:
			if event.Retry != 0 {
				// TODO: Handle reconnection time.
				c.logger.Debug("received retry event", slog.Int("retry", event.Retry))
			}
			if event.IsComment() {
				c.logger.Debug("received comment", slog.String("comment", event.Comment))
				continue
			}
			if err := handler(ctx, event); err != nil {
				c.logger.Error("handling event", slog.Any("error", err))
				return err
			}
		}
	}
}

// Subscribe to an SSE event data stream.
// This is a non-blocking call.
//
// NOTE: The event reader should not be shared between multiple listeners, otherwise, there might be unexpected parsing results.
func (c *Client) Subscribe(ctx context.Context, path string, params any, handler EventStreamHandler, opts ...model.RequestOption) error {
	r, err := c.listenToSSE(ctx, path, params, opts...)
	if err != nil {
		return fmt.Errorf("initializing SSE stream: %w", err)
	}
	defer func() {
		if err := r.Close(); err != nil {
			c.logger.Error("closing stream", slog.Any("error", err))
		}
	}()

	evtChannel, errChannel := make(chan *sse.Event, 1), make(chan error, 1)
	go c.startReadingSSE(r, evtChannel, errChannel)
	go func() {
	L:
		for {
			select {
			case <-ctx.Done():
				c.logger.Error("context cancelled", slog.Any("error", ctx.Err()))
				break L
			case err := <-errChannel:
				if errors.Is(err, io.EOF) {
					continue
				}
				c.logger.Error("reading from stream", slog.Any("error", err))
				break L
			case event := <-evtChannel:
				if event.Retry != 0 {
					// TODO: Handle reconnection time.
					c.logger.Debug("received retry event", slog.Int("retry", event.Retry))
				}
				if event.IsComment() {
					c.logger.Debug("received comment", slog.String("comment", event.Comment))
					continue
				}
				if err := handler(ctx, event); err != nil {
					c.logger.Error("handling event", slog.Any("error", err))
					break L
				}
			}
		}
	}()

	return nil
}

func (c *Client) listenToSSE(ctx context.Context, path string, params any, opts ...model.RequestOption) (io.ReadCloser, error) {
	uri, err := c.encoder.EncodeParams(path, params)
	if err != nil {
		return nil, err
	}
	options := mergeOptions(opts...)

	req := c.HTTP.R().SetContext(ctx)
	req.SetQueryParamsFromValues(options.QueryParams)
	req.SetHeaderMultiValues(options.Headers)
	req.SetError(&alpacaerrors.ResponseError{})
	req.SetHeader("Accept", "text/event-stream")
	req.SetHeader("Connection", "keep-alive")
	req.SetHeader("Cache-Control", "no-cache")
	// Not parsing the response enables us to read the raw response body without it
	// getting closed. Hence, allowing the SSE client to keep reading from the stream.
	req.SetDoNotParseResponse(true)

	res, err := c.executeRequest(ctx, req, http.MethodGet, uri, options.Trace)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	return res.RawBody(), nil
}

func (c *Client) startReadingSSE(r io.ReadCloser, evtCh chan<- *sse.Event, errCh chan<- error) {
	parser := sse.NewParser()
	reader := bufio.NewReader(r)

	go func() {
		for {
			l, err := reader.ReadString('\n')
			if err != nil {
				errCh <- fmt.Errorf("reading message: %w", err)
				return
			}
			evt, err := parser.ParseEvent([]byte(l))
			if err != nil {
				errCh <- fmt.Errorf("parsing event: %w", err)
				return
			}
			if evt.IsEmpty() {
				continue
			}
			evtCh <- evt
		}
	}()
}

func mergeOptions(opts ...model.RequestOption) *model.RequestOptions {
	options := &model.RequestOptions{}
	for _, o := range opts {
		o(options)
	}

	return options
}

func parseResponseError(res *resty.Response) *alpacaerrors.ResponseError {
	responseError, ok := alpacaerrors.AsResponseError(res.Error())
	if !ok || responseError == nil {
		return nil
	}
	responseError.RequestID = res.Header().Get("X-Request-ID")
	responseError.StatusCode = res.StatusCode()
	b := struct {
		Message string `json:"message"`
	}{}
	if err := json.Unmarshal(res.Body(), &b); err != nil {
		responseError.Message = string(res.Body())
	} else {
		responseError.Message = b.Message
	}
	return responseError
}
