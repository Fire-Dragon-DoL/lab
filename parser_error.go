package lab

// ParserError represents any error happening during parsing phase
type ParserError uint

// Input parsing errors
const (
	_ ParserError = iota
	ErrNotParsable
)

// Error provides a string representation of the error code
func (parseErr ParserError) Error() string {
	var text string

	switch parseErr {
	case ErrNotParsable:
		text = "Content not parsable"
	default:
		text = "Unknown parser error"
	}

	return text
}
