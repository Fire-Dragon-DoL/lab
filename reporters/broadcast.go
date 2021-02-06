package reporters

import "github.com/Fire-Dragon-DoL/lab/encoding/gotest"

// Broadcast is a reporter that register test events on multiple reporters
type Broadcast []Reporter

// Write acknowledges a test event
func (reporters Broadcast) Write(testEvent *gotest.TestEvent) {
	for _, reporter := range reporters {
		reporter.Write(testEvent)
	}
}
