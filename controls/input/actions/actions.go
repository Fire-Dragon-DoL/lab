// Package actions provides controls for input lines based on action
package actions

import (
	"fmt"
	"io"
	"strings"
)

// ActionExample builds a string line as an `io.Reader`
func ActionExample(context string) io.Reader {
	line := fmt.Sprintf("%s\n", context)
	return strings.NewReader(line)
}
