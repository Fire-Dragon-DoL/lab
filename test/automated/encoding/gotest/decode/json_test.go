package gotest_test

import (
	"testing"

	. "github.com/Fire-Dragon-DoL/lab/asserting"
	"github.com/Fire-Dragon-DoL/lab/controls/input/actions"
	"github.com/Fire-Dragon-DoL/lab/encoding/gotest"
)

func TestDecodeJSON(t *testing.T) {
	t.Run("Decode", func(t *testing.T) {
		t.Run("Input is JSON", func(t *testing.T) {
			input := actions.OutputPassExample()
			decoder := gotest.NewDecoder(input)
			testEvent, err := decoder.Decode()

			t.Run("Success", func(t *testing.T) {
				Assert(t, err == nil)
			})

			t.Run("Action is 'output'", func(t *testing.T) {
				Assert(t, testEvent.Action == "output")
			})
		})
	})
}
