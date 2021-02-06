package reporters

import "github.com/Fire-Dragon-DoL/lab/encoding/gotest"

// Doc is a reporter that outputs text in a documentation-like format
type Doc struct {
}

// Write acknowledges a test event
func (d *Doc) Write(testEvent *gotest.TestEvent) {

}
