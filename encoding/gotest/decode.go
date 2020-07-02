package gotest

import (
	"encoding/json"
	"io"
)

// Decoder reads and decodes `go test` output
type Decoder struct {
	jsonDecoder *json.Decoder
}

// NewDecoder starts a `Decoder` for the provided `io.Reader`
func NewDecoder(rd io.Reader) *Decoder {
	decoder := &Decoder{}
	decoder.jsonDecoder = json.NewDecoder(rd)
	return decoder
}

// Decode returns the next `TestEvent` in the `go test` output
func (decoder *Decoder) Decode() (*TestEvent, error) {
	testEvent := &TestEvent{}
	err := decoder.jsonDecoder.Decode(testEvent)
	return testEvent, err
}

// // Decode converts `go test` output into in-memory format
// func Decode(data []byte) (*TestEvent, error) {
// 	testEvent := &TestEvent{}
// 	err := json.Unmarshal(data, testEvent)
// 	return testEvent, err
// }

// // DecodeString converts `go test` line into in-memory format
// func DecodeString(line string) (*TestEvent, error) {
// 	testEvent := &TestEvent{}
// 	data := []byte(line)
// 	err := json.Unmarshal(data, testEvent)
// 	return testEvent, err
// }
