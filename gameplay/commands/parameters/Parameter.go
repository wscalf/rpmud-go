package parameters

// Parameter represents an abstract parameter to a command.
type Parameter interface {
	// Name is the parameter name that will be displayed in help text and used to refer to the value from commands.
	Name() string
	// IsDelimiter indicates whether or not this parameter blocks processing of the previous parameter.
	IsDelimiter() bool
	// IsRequired indicates whether or not this parameter is required.
	IsRequired() bool
	// Find locates the next matching text in the input and returns the index of the first character.
	Find(command string) int
	// Consume creates a Value from the input text and returns the value and the remaining, unused text
	Consume(command string) (Value, string)
}
