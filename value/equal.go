package value

type equaler interface {
	Equal(interface{}) bool
}

// Equal returns true if the two values are equal.
func Equal(a, b interface{}) bool {
	if a, ok := a.(equaler); ok {
		return a.Equal(b)
	}
	if b, ok := b.(equaler); ok {
		return b.Equal(a)
	}

	// TODO: equal for numeric types, for sequences, etc.

	return a == b
}