package gotest

import (
	"time"
)

// TestEvent represents a `go test` line
type TestEvent struct {
	Time      time.Time // encodes as an RFC3339-format string
	Action    string
	Package   string
	Test      string
	Elapsed   float64 // seconds
	Output    string
	operation *string
}

var unknownOperation = "UNKNOWN"

// Operation returns the operation documented by this event. It's one of RUN, PAUSE, CONT, OUTPUT
func (t *TestEvent) Operation() string {
	if t.operation != nil {
		return *t.operation
	}

	matches := OperationRegexp.FindStringSubmatch(t.Output)
	if matches == nil || len(matches) < 2 {
		t.operation = &unknownOperation
		return *t.operation
	}

	t.operation = &matches[1]

	return *t.operation
}
