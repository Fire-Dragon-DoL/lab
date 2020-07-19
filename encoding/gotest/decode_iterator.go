package gotest

import (
	"io"
)

// DecodeIterator decodes each input line
type DecodeIterator struct {
	decoder    *Decoder
	testEvent  *TestEvent
	isFinished bool
	err        error
}

// NewDecodeIterator provides an iterator for decoding
func NewDecodeIterator(rd io.Reader) *DecodeIterator {
	decoder := NewDecoder(rd)
	iterator := DecodeIterator{decoder: decoder}
	return &iterator
}

// Next decodes the next input line
func (iterator *DecodeIterator) Next() bool {
	if iterator.ended() {
		return false
	}

	iterator.testEvent = nil
	testEvent, err := iterator.decoder.Decode()

	switch err {
	case io.EOF:
		iterator.isFinished = true
	case nil:
		iterator.testEvent = testEvent
	default:
		iterator.err = err
	}

	return iterator.needsIteration()
}

// Error returns the last recorded `error`
func (iterator *DecodeIterator) Error() error {
	return iterator.err
}

// Get returns the current decoded data for the current iteration
func (iterator *DecodeIterator) Get() interface{} {
	return iterator.testEvent
}

func (iterator *DecodeIterator) ended() bool {
	return iterator.isFinished || iterator.err != nil
}

func (iterator *DecodeIterator) needsIteration() bool {
	return !iterator.ended()
}
