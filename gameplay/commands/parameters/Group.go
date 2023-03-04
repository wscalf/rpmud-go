package parameters

type Group struct {
	inner    Parameter
	name     string
	required bool
}

// Name is the parameter name that will be displayed in help text and used to refer to the value from commands.
func (g Group) Name() string {
	return g.name
}

// IsDelimiter indicates whether or not this parameter blocks processing of the previous parameter.
func (g Group) IsDelimiter() bool {
	return false
}

// IsRequired indicates whether or not this parameter is required.
func (g Group) IsRequired() bool {
	return g.required
}

// Find locates the next matching text in the input and returns the index of the first character.
func (g Group) Find(command string) int {
	return g.inner.Find(command)
}

// Consume creates a Value from the input text and returns the value and the remaining, unused text
func (g Group) Consume(command string) (Value, string) {
	remaining := command
	values := make([]string, 0, 4)

	for g.inner.Find(remaining) == 0 {
		var v Value
		v, remaining = g.inner.Consume(remaining)

		values = append(values, v.Single())
	}

	return NewMultiValue(values), remaining
}

func NewNameGroup(name string, required bool) Group {
	return Group{
		inner:    NewName(name, true),
		name:     name,
		required: required,
	}
}
