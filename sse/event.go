package sse

import (
	"bytes"
	"time"
)

const (
	FieldNameData  = "data"
	FieldNameRetry = "retry"
)

// Event represents an individual event from the event stream.
// This is the simplest form of an event that only contains the data that is expected to be a single line of
// JSON data, a comment or the retry indicator.
type Event struct {
	timestamp time.Time

	Data    []byte
	Comment string
	Retry   int
}

func (e *Event) IsComment() bool {
	return e.Comment != ""
}

func (e *Event) IsEmpty() bool {
	return len(e.Data) == 0 && e.Comment == "" && e.Retry == 0
}

func NewEvent(data []byte, comment []byte, retry int) *Event {
	comment = bytes.TrimSpace(comment)
	return &Event{
		Data:      data,
		Retry:     retry,
		Comment:   string(comment),
		timestamp: time.Now(),
	}
}
