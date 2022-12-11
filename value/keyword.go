package value

import (
	"fmt"
	"strings"
)

// Keyword represents a keyword. Syyntactically, a keyword is a symbol
// that starts with a colon and evaluates to itself.
type Keyword struct {
	Section
	Value string
}

func NewKeyword(s string, opts ...Option) *Keyword {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	return &Keyword{
		Section: o.section,
		Value:   s,
	}
}

func (k *Keyword) Namespace() string {
	// Return the namespace of the keyword, or the empty string if it
	// doesn't have one.
	if i := strings.Index(k.Value, "/"); i != -1 {
		return k.Value[:i]
	}
	return ""
}

func (k *Keyword) Name() string {
	// Return the name of the keyword, or the empty string if it
	// doesn't have one.
	if i := strings.Index(k.Value, "/"); i != -1 {
		return k.Value[i+1:]
	}
	return k.Value
}

func (k *Keyword) String() string {
	return ":" + k.Value
}

func (k *Keyword) Equal(v interface{}) bool {
	other, ok := v.(*Keyword)
	if !ok {
		return false
	}
	return k.Value == other.Value
}

func (k *Keyword) Apply(env Environment, args []interface{}) (interface{}, error) {
	if len(args) == 0 || len(args) > 2 {
		return nil, fmt.Errorf("wrong number of args (%v) passed to: %v", len(args), k)
	}
	var defaultVal interface{} = nil
	if len(args) == 2 {
		defaultVal = args[1]
	}

	assoc, ok := args[0].(Associative)
	if !ok {
		return defaultVal, nil
	}
	v, ok := assoc.EntryAt(k)
	if !ok {
		return defaultVal, nil
	}
	return v, nil
}
