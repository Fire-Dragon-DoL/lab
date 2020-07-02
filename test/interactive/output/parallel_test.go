package interactive_test

import "testing"

func TestInteractive(t *testing.T) {
	t.Parallel()
	t.Run("One", func(t *testing.T) {})
	t.Run("Two", func(t *testing.T) {})

	t.Run("Nested 3", func(t *testing.T) {
		t.Run("Nested 3 One", func(t *testing.T) {})
		t.Run("Nested 3 Two", func(t *testing.T) {
			t.Error("Foo")
		})
	})
}
