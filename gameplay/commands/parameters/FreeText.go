package parameters

import "strings"

type FreeText struct {
	name     string
	required bool
}

// Name is the parameter name that will be displayed in help text and used to refer to the value from commands.
func (f FreeText) Name() string {
	return f.name
}

// IsDelimiter indicates whether or not this parameter blocks processing of the previous parameter.
func (f FreeText) IsDelimiter() bool {
	return false
}

// IsRequired indicates whether or not this parameter is required.
func (f FreeText) IsRequired() bool {
	return f.required
}

// Find locates the next matching text in the input and returns the index of the first character.
func (f FreeText) Find(command string) int {
	if len(command) > 0 {
		return 0
	} else {
		return -1
	}
}

// Consume creates a Value from the input text and returns the value and the remaining, unused text
func (f FreeText) Consume(command string) (Value, string) {
	v := NewSingleValue(strings.TrimSpace(command))

	return v, ""
}

func NewFreeText(name string, required bool) FreeText {
	return FreeText{name: name, required: required}
}
