package gotest_test

import (
	"testing"

	. "github.com/Fire-Dragon-DoL/lab/asserting"
	"github.com/Fire-Dragon-DoL/lab/controls/input/actions"
	"github.com/Fire-Dragon-DoL/lab/encoding/gotest"
)

func TestDecodeIterator(t *testing.T) {
	t.Run("Decode", func(t *testing.T) {
		input := actions.OutputPassExample()
		iterator := gotest.NewDecodeIterator(input)
		iterations := 0
		var decodeErr error
		var decodedData interface{}

		for iterator.Next() {
			decodedData, decodeErr = iterator.Get()
			iterations += 1
		}

		t.Run("Iterates twice", func(t *testing.T) {
			Assert(t, iterations == 2)
		})

		t.Run("Has no errors", func(t *testing.T) {
			Assert(t, decodeErr == nil)
		})

		_, ok := decodedData.(*gotest.TestEvent)

		t.Run("Data is a TestEvent", func(t *testing.T) {
			Assert(t, ok)
		})
	})
}
