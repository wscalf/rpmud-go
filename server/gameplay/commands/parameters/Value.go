package parameters

// Value represents a parameter value. It may have a single value or many.
type Value struct {
	values []string
}

// NewMultiValue creates a new Value from the provided string(s)
func NewMultiValue(values []string) Value {
	return Value{values: values}
}

// NewSingleValue creates a new Value from the provided string
func NewSingleValue(value string) Value {
	return Value{values: []string{value}}
}

// IsMulti indicates whether or not there are multiple values, returning true if there are, otherwise false
func (p Value) IsMulti() bool {
	return len(p.values) > 1
}

// Single gets the first string from the Value
func (p Value) Single() string {
	return p.values[0]
}

// Multiple gets a slice of all values. For single-value cases, this is a slice of 1.
func (p Value) Multiple() []string {
	return p.values
}
