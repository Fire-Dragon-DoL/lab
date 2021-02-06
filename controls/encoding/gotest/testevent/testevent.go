// Package testevent contains controls for gotest.TestEvent
package testevent

import "github.com/Fire-Dragon-DoL/lab/encoding/gotest"

// OperationOutputExample outputs a test event with no operation text
func OperationOutputExample() *gotest.TestEvent {
	return &gotest.TestEvent{
		Action:  "output",
		Package: "command-line-arguments",
		Test:    "Some/Namespace/Test",
		Output:  "Some output",
	}
}

// OperationRunExample outputs a test event with run operation text
func OperationRunExample() *gotest.TestEvent {
	return &gotest.TestEvent{
		Action:  "output",
		Package: "command-line-arguments",
		Test:    "Some/Namespace/Test",
		Output:  "Some output",
	}
}
