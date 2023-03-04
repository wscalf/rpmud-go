package parameters

import (
	"strings"
)

type Delimiter struct {
	text string
}

// Name is the parameter name that will be displayed in help text and used to refer to the value from commands.
func (d Delimiter) Name() string {
	return d.text
}

// IsDelimiter indicates whether or not this parameter blocks processing of the previous parameter.
func (d Delimiter) IsDelimiter() bool {
	return true
}

// IsRequired indicates whether or not this parameter is required.
func (d Delimiter) IsRequired() bool {
	return true
}

// Find locates the next matching text in the input and returns the index of the first character.
func (d Delimiter) Find(command string) int {
	return strings.Index(command, d.text)
}

// Consume creates a Value from the input text and returns the value and the remaining, unused text
func (d Delimiter) Consume(command string) (Value, string) {
	v := NewSingleValue(d.text)
	remaining := strings.TrimSpace(command[len(d.text):])

	return v, remaining
}

func NewDelimiter(text string) Delimiter {
	return Delimiter{
		text: text,
	}
}
