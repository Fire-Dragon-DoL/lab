// Package controls exposes sample data to control the testing environment
package controls

import (
	"io"
	"strings"
)

// InputNotParsableExample is a non-parsable reader
func InputNotParsableExample() io.Reader {
	return strings.NewReader("Not Parsable Content\n")
}

// OutputStringExample is a `strings.Builder`
func OutputStringExample() *strings.Builder {
	return &strings.Builder{}
}
