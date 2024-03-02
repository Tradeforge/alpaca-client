package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"go.tradeforge.dev/alpaca/encoder"
	"go.tradeforge.dev/alpaca/model"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

const clientVersion = "v0.0.0"

const (
	DefaultRetryCount    = 3
	DefaultClientTimeout = 10 * time.Second
)

type EventReader interface {
	Listen(ctx context.Context, stream io.Reader) (<-chan Event, <-chan error)
}

type Event interface {
	GetData() []byte
	GetTimestamp() time.Time
}

// Client defines an HTTP client for the Polygon REST API.
type Client struct {
	HTTP *resty.Client

	encoder     *encoder.Encoder
	eventReader EventReader

	logger *zap.Logger
}

// New returns a new client with the specified API key and config.
func New(
	apiURL string,
	reader EventReader,
	logger *zap.Logger,
) *Client {
	return newClient(apiURL, reader, logger)
}

func newClient(
	apiURL string,
	reader EventReader,
	logger *zap.Logger,
) *Client {
	c := resty.New()

	c.SetBaseURL(apiURL)
	c.SetRetryCount(DefaultRetryCount)
	c.SetHeader("User-Agent", fmt.Sprintf("Alpaca client/%v", clientVersion))
	c.SetHeader("Accept", "application/json")

	return &Client{
		HTTP:        c,
		encoder:     encoder.New(),
		eventReader: reader,
		logger:      logger,
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
	req.SetResult(response).SetError(&model.ResponseError{})
	req.SetHeader("Content-Type", "application/json")

	res, err := req.Execute(method, uri)
	if err != nil {
		c.logger.Error(
			err.Error(),
			zap.Any("response", res))
		return fmt.Errorf("failed to execute request: %w", err)
	} else if res.IsError() {
		c.logger.Error(
			res.String(),
			zap.Any("response", res))
		responseError := parseResponseError(res)
		return responseError
	}

	if options.Trace {
		sanitizedHeaders := req.Header
		for k := range sanitizedHeaders {
			if k == "Authorization" {
				sanitizedHeaders[k] = []string{"REDACTED"}
			}
		}
		c.logger.Debug(
			"request",
			zap.String("url", uri),
			zap.Any("request headers", sanitizedHeaders),
			zap.Any("response headers", res.Header()),
		)
	}

	return nil
}

type EventStreamHandler func(ctx context.Context, event Event) error

// Listen to an event data stream.
// This is a blocking call that will continue to read from the stream until the context is canceled
// or the watch is stopped.
func (c *Client) Listen(ctx context.Context, path string, params any, handler EventStreamHandler, opts ...model.RequestOption) error {
	uri, err := c.encoder.EncodeParams(path, params)
	if err != nil {
		return err
	}

	options := mergeOptions(opts...)

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
	req.SetError(&model.ResponseError{})
	req.SetHeader("Accept", "text/event-stream")
	req.SetHeader("Connection", "keep-alive")
	req.SetHeader("Cache-Control", "no-cache")
	// Not parsing the response enables us to read the raw response body without it
	// getting closed. Hence allowing the SSE client to keep reading from the stream.
	req.SetDoNotParseResponse(true)

	res, err := req.Execute(http.MethodGet, uri)
	if err != nil {
		c.logger.Error(
			err.Error(),
			zap.Any("response", res))
		return fmt.Errorf("failed to execute request: %w", err)
	} else if res.IsError() {
		c.logger.Error(
			res.String(),
			zap.Any("response", res))
		responseError := parseResponseError(res)
		return responseError
	}

	if options.Trace {
		sanitizedHeaders := req.Header
		for k := range sanitizedHeaders {
			if k == "Authorization" {
				sanitizedHeaders[k] = []string{"REDACTED"}
			}
		}
		c.logger.Debug(
			"request",
			zap.String("url", uri),
			zap.Any("request headers", sanitizedHeaders),
			zap.Any("response headers", res.Header()),
		)
	}
	defer res.RawBody().Close()

	evtChannel, errChannel := c.eventReader.Listen(ctx, res.RawBody())
	var evt Event
	for {
		select {
		case err = <-errChannel:
			return err
		case evt = <-evtChannel:
			if err := handler(ctx, evt); err != nil {
				c.logger.Error("handling event", zap.Error(err), zap.Any("event", evt))
				return err
			}
		}
	}
}

func mergeOptions(opts ...model.RequestOption) *model.RequestOptions {
	options := &model.RequestOptions{}
	for _, o := range opts {
		o(options)
	}

	return options
}

func parseResponseError(res *resty.Response) *model.ResponseError {
	if res == nil {
		return nil
	}
	responseError := res.Error().(*model.ResponseError)
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
