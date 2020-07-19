package lab_test

import (
	"testing"

	"github.com/Fire-Dragon-DoL/lab"
	"github.com/Fire-Dragon-DoL/lab/controls"
	"github.com/Fire-Dragon-DoL/lab/controls/input/actions"
	. "github.com/thehungry-dev/asserting"
)

func TestTransformJSONActionRun(t *testing.T) {
	t.Run("Transform JSON", func(t *testing.T) {
		t.Run("Action is run", func(t *testing.T) {
			input := actions.RunExample()
			output := controls.OutputStringExample()
			err := lab.Transform(input, output)

			t.Run("No error", func(t *testing.T) {
				Assert(t, err == nil)
			})

			t.Run("Output is empty", func(t *testing.T) {
				Assert(t, output.String() == "")
			})
		})
	})
}
