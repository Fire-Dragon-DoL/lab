package lab_test

import (
	"testing"

	"github.com/Fire-Dragon-DoL/lab"
	. "github.com/Fire-Dragon-DoL/lab/asserting"
	"github.com/Fire-Dragon-DoL/lab/controls"
	"github.com/Fire-Dragon-DoL/lab/controls/input/actions"
)

func TestTransformJSONActionOutputPass(t *testing.T) {
	t.Run("Transform JSON", func(t *testing.T) {
		t.Run("Action is 'output'", func(t *testing.T) {
			input := actions.OutputPassExample()
			output := controls.OutputStringExample()
			err := lab.Transform(input, output)

			t.Run("No error", func(t *testing.T) {
				Assert(t, err == nil)
			})

			t.Run("Output is written in human case", func(t *testing.T) {})
			t.Run("Output excludes test function name", func(t *testing.T) {})
			t.Run("Output is green", func(t *testing.T) {})
		})
	})
}
