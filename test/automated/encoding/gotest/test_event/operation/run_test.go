package gotest_test

import (
	"testing"

	"github.com/Fire-Dragon-DoL/lab/controls/encoding/gotest/testevent"
	. "github.com/thehungry-dev/asserting"
)

func TestGoTestOperationOutput(t *testing.T) {
	t.Run("TestEvent", func(t *testing.T) {
		t.Run("Operation", func(t *testing.T) {
			t.Run("Action starts with `=== RUN`", func(t *testing.T) {
				testEvent := testevent.OperationRunExample()
				operation := testEvent.Operation()

				t.Run("Is RUN", func(t *testing.T) {
					Assert(t, operation == "RUN")
				})
			})
		})
	})
}
