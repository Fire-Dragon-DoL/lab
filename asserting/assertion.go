package asserting

import "testing"

// Assert tests if the result is true
var Assert Assertion = Assertion(assertf)

// Assertion provides a function to assert results, as well as a function to assert "comma ok" tuples and assert and recover from functions triggering panic
type Assertion func(testing.TB, bool, ...interface{})

// PanicMsg asserts that the provided function triggers panic with the provided message
func (assert Assertion) PanicMsg(t testing.TB, do func(), assertMsg func(interface{}) bool) {
	panicked := true

	defer func() {
		err := recover()

		assert(t, panicked, "Panic expected")

		if panicked {
			result := assertMsg(err)
			assert(t, result, "Invalid panic message")
		}
	}()
	do()

	panicked = false
}

// Panic asserts that the provided function triggers panic
func (assert Assertion) Panic(t testing.TB, do func()) {
	assert.PanicMsg(t, do, any)
}

func any(_ interface{}) bool { return true }

func assertf(t testing.TB, result bool, msgArgs ...interface{}) {
	msg := "Assertion failed"
	var args []interface{}

	switch len(msgArgs) {
	case 0:
	case 1:
		msg = msgArgs[0].(string)
	default:
		msg = msgArgs[0].(string)
		args = msgArgs[1:]
	}

	if !result {
		t.Errorf(msg, args...)
	}
}
