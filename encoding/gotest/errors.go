package gotest

// DecodeError represents any error happening during decoding phase
type DecodeError uint

// Input decoding errors
const (
	_ DecodeError = iota
	ErrContentFormatInvalid
)

// DecodeError provides a string representation of the error code
func (decodeErr DecodeError) Error() string {
	var text string

	switch decodeErr {
	case ErrContentFormatInvalid:
		text = "Input format is invalid"
	default:
		text = "Unknown error"
	}

	return text
}
