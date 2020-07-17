package lab

import (
	"bufio"
	"io"
	"strings"

	"github.com/Fire-Dragon-DoL/lab/encoding/gotest"
)

// type Event string

// type Parser interface {
// 	Parse(line string) (Event, error)
// }

// type Reporter interface {
// 	Write(event Event) error
// 	WriteSummary() error
// }

// type Transformer interface {
// 	Transform(input io.Reader, output io.Writer)
// }

// Transform converts the text output of `go test -json` into a human-readable, colored format
func Transform(input io.Reader, output io.Writer) error {
	// linesScanner := file.NewLinesScanner(input)

	// for linesScanner.Next() {
	// 	line :=
	// }
	buf := bufio.NewReader(input)
	lineBuf, _, _ := buf.ReadLine()
	var entireLine strings.Builder
	entireLine.Write(lineBuf)
	line := entireLine.String()

	if strings.Contains(line, "run") {
		return nil
	}

	if strings.Contains(line, "output") {
		return nil
	}

	if strings.Contains(line, "pass") {
		return nil
	}

	return gotest.ErrContentFormatInvalid
}

// type jsonTransformer struct {
// 	Parser    Parser
// 	Reporter  Reporter
// 	Something string
// }

// // Transform converts the test input into an enchanced version
// func (transformer *jsonTransformer) Transform(input io.Reader, output io.Writer) {
// 	// transformer.Parser.Parse()
// }
