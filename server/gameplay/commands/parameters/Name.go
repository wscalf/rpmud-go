package parameters

import "strings"

type Name struct {
	name     string
	required bool
}

// Name is the parameter name that will be displayed in help text and used to refer to the value from commands.
func (n Name) Name() string {
	return n.name
}

// IsDelimiter indicates whether or not this parameter blocks processing of the previous parameter.
func (n Name) IsDelimiter() bool {
	return false
}

// IsRequired indicates whether or not this parameter is required.
func (n Name) IsRequired() bool {
	return n.required
}

// Find locates the next matching text in the input and returns the index of the first character.
func (n Name) Find(command string) int {
	start, _ := find(command)

	if start > -1 {
		return 0
	} else {
		return start
	}
}

// Consume creates a Value from the input text and returns the value and the remaining, unused text
func (n Name) Consume(command string) (Value, string) {
	start, end := find(command)

	match := command[start:end]
	v := NewSingleValue(strings.Trim(match, `"`))
	return v, strings.TrimSpace(command[end:])
}

func NewName(name string, required bool) Name {
	return Name{name: name, required: required}
}

func find(text string) (int, int) {
	if len(text) == 0 {
		return -1, -1
	}

	if strings.HasPrefix(text, `"`) {
		start := 0
		end := strings.Index(text[1:], `"`)
		inclusiveEnd := end + 2 //Account for skipping first quote and including second
		if end > -1 {
			return start, inclusiveEnd
		} else {
			return -1, -1
		}
	} else {
		if end := strings.Index(text, " "); end != -1 {
			return 0, end
		} else {
			return 0, len(text)
		}
	}
}
