package sse

import (
	"bytes"
	"fmt"
	"strconv"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseEvent(data []byte) (*Event, error) {
	if len(bytes.TrimSpace(data)) == 0 {
		return &Event{}, nil
	}
	if data[0] == ':' {
		return NewEvent(nil, data[1:], 0), nil
	}
	var (
		field = data
		value []byte
	)
	if i := bytes.IndexRune(data, ':'); i != -1 {
		field = data[:i]
		value = data[i+1:]
		if len(value) != 0 && value[0] == ' ' {
			value = value[1:]
		}
	}
	switch string(field) {
	case FieldNameData:
		return NewEvent(value, nil, 0), nil
	case FieldNameRetry:
		i, err := strconv.Atoi(string(value))
		if err != nil {
			return nil, fmt.Errorf("parsing retry field: %w", err)
		}
		return NewEvent(nil, nil, i), nil
	default:
		return nil, fmt.Errorf("unknown or unexpected field: %s", field)
	}
}
