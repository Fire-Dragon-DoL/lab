// Package gotest provides functionality to decode `go test -json` output
package gotest

import "regexp"

// OperationRegexp is a regular expression to extract TestEvent operation from Output
var OperationRegexp *regexp.Regexp

func init() {
	OperationRegexp = regexp.MustCompile(`^\s*={3}\s+(\S+)\s+`)
}
