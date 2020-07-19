package gotest_test

import (
	"testing"

	"github.com/Fire-Dragon-DoL/lab/controls"
	"github.com/Fire-Dragon-DoL/lab/encoding/gotest"
	. "github.com/thehungry-dev/asserting"
)

func TestDecodeNotJSON(t *testing.T) {
	t.Run("Decode", func(t *testing.T) {
		t.Run("Input is not JSON", func(t *testing.T) {
			input := controls.InputNotParsableExample()
			decoder := gotest.NewDecoder(input)
			_, err := decoder.Decode()

			t.Run("Error", func(t *testing.T) {
				Assert(t, err == gotest.ErrContentFormatInvalid)
			})
		})
	})
}
