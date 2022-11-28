package value

import (
	"fmt"
	"strconv"
)

type printOptions struct {
	printReadably bool
}

type PrintOption func(*printOptions)

func PrintReadably() PrintOption {
	return func(o *printOptions) {
		o.printReadably = true
	}
}

func ToString(v interface{}, opts ...PrintOption) string {
	options := printOptions{}
	for _, opt := range opts {
		opt(&options)
	}

	// if v is a Stringer, use its String method
	if s, ok := v.(fmt.Stringer); ok {
		return s.String()
	}

	switch v := v.(type) {
	case nil:
		return "nil"
	case string:
		if options.printReadably {
			return v
		}
		// NB: java does not support \x escape sequences, but go does.  this
		// results in a difference in the output of the string from Clojure
		// if such characters make it into the string. We will escape them
		// but Clojure on the JVM will not.
		return strconv.Quote(v)
	case bool:
		if v {
			return "true"
		}
		return "false"
	case float64:
		if v == float64(int64(v)) {
			return fmt.Sprintf("%d.0", int64(v))
		}
		return strconv.FormatFloat(v, 'f', -1, 64)
	case uint64, uint32, uint16, uint8, uint, int64, int32, int16, int8, int:
		return fmt.Sprintf("%d", v)
	}

	return fmt.Sprintf("%T", v)
}