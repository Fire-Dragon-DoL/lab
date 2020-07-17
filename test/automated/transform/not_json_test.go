package lab_test

import (
	"testing"

	"github.com/Fire-Dragon-DoL/lab"
	. "github.com/Fire-Dragon-DoL/lab/asserting"
	"github.com/Fire-Dragon-DoL/lab/controls"
	"github.com/Fire-Dragon-DoL/lab/encoding/gotest"
)

func TestTransformNotJSON(t *testing.T) {
	t.Run("Transform non-JSON text", func(t *testing.T) {
		input := controls.InputNotParsableExample()
		output := controls.OutputStringExample()
		err := lab.Transform(input, output)

		t.Run("Error not parsable", func(t *testing.T) {
			Assert(t, err == gotest.ErrContentFormatInvalid)
		})
	})
}
