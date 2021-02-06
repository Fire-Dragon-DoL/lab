// Package reporters provides a series of different built-in reporters for `lab` that can be used to modify the way test output is rendered
package reporters

import "github.com/Fire-Dragon-DoL/lab/encoding/gotest"

// Reporter represents any reporter
type Reporter interface {
	Write(*gotest.TestEvent)
}
