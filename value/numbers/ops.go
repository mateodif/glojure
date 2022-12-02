package numbers

import (
	"fmt"

	"github.com/glojurelang/glojure/value"
)

type (
	ops interface {
		Combine(y ops) ops

		Add(x, y interface{}) interface{}
		AddP(x, y interface{}) interface{}
	}
	int64Ops      struct{}
	bigIntOps     struct{}
	ratioOps      struct{}
	bigDecimalOps struct{}
	float64Ops    struct{}
)

func Ops(x interface{}) ops {
	switch x.(type) {
	case int8:
		return int64Ops{}
	case int16:
		return int64Ops{}
	case int32:
		return int64Ops{}
	case int64:
		return int64Ops{}
	case uint8:
		return int64Ops{}
	case uint16:
		return int64Ops{}
	case uint32:
		return int64Ops{}
	case uint64:
		return int64Ops{}
	case float32:
		return float64Ops{}
	case float64:
		return float64Ops{}
	case *value.Ratio:
		return ratioOps{}
	case *value.BigInt:
		return bigIntOps{}
	case *value.BigDecimal:
		return bigDecimalOps{}
	default:
		panic(fmt.Sprintf("cannot convert %T to Ops", x))
	}
}

func Add(x, y interface{}) interface{} {
	return Ops(x).Combine(Ops(y)).Add(x, y)
}
func AddP(x, y interface{}) interface{} {
	return Ops(x).Combine(Ops(y)).AddP(x, y)
}
func (o int64Ops) Add(x, y interface{}) interface{} {
	return AsInt64(x) + AsInt64(y)
}
func (o int64Ops) AddP(x, y interface{}) interface{} {
	return AsInt64(x) + AsInt64(y)
}
func (o bigIntOps) Add(x, y interface{}) interface{} {
	return AsBigInt(x).Add(AsBigInt(y))
}
func (o bigIntOps) AddP(x, y interface{}) interface{} {
	return AsBigInt(x).AddP(AsBigInt(y))
}
func (o ratioOps) Add(x, y interface{}) interface{} {
	return AsRatio(x).Add(AsRatio(y))
}
func (o ratioOps) AddP(x, y interface{}) interface{} {
	return AsRatio(x).AddP(AsRatio(y))
}
func (o bigDecimalOps) Add(x, y interface{}) interface{} {
	return AsBigDecimal(x).Add(AsBigDecimal(y))
}
func (o bigDecimalOps) AddP(x, y interface{}) interface{} {
	return AsBigDecimal(x).AddP(AsBigDecimal(y))
}
func (o float64Ops) Add(x, y interface{}) interface{} {
	return AsFloat64(x) + AsFloat64(y)
}
func (o float64Ops) AddP(x, y interface{}) interface{} {
	return AsFloat64(x) + AsFloat64(y)
}
func (o int64Ops) Combine(y ops) ops {
	switch y.(type) {
	case int64Ops:
		return o
	case bigIntOps:
		return y
	case ratioOps:
		return y
	case bigDecimalOps:
		return y
	case float64Ops:
		return y
	default:
		panic("cannot combine Ops")
	}
}
func (o bigIntOps) Combine(y ops) ops {
	switch y.(type) {
	case int64Ops:
		return o
	case bigIntOps:
		return o
	case ratioOps:
		return y
	case bigDecimalOps:
		return y
	case float64Ops:
		return y
	default:
		panic("cannot combine Ops")
	}
}
func (o ratioOps) Combine(y ops) ops {
	switch y.(type) {
	case int64Ops:
		return o
	case bigIntOps:
		return o
	case ratioOps:
		return o
	case bigDecimalOps:
		return y
	case float64Ops:
		return y
	default:
		panic("cannot combine Ops")
	}
}
func (o bigDecimalOps) Combine(y ops) ops {
	switch y.(type) {
	case int64Ops:
		return o
	case bigIntOps:
		return o
	case ratioOps:
		return o
	case bigDecimalOps:
		return o
	case float64Ops:
		return y
	default:
		panic("cannot combine Ops")
	}
}
func (o float64Ops) Combine(y ops) ops {
	switch y.(type) {
	case int64Ops:
		return o
	case bigIntOps:
		return o
	case ratioOps:
		return o
	case bigDecimalOps:
		return o
	case float64Ops:
		return o
	default:
		panic("cannot combine Ops")
	}
}
func AsInt64(x interface{}) int64 {
	switch x := x.(type) {
	case int8:
		return AsInt64(x)
	case int16:
		return AsInt64(x)
	case int32:
		return AsInt64(x)
	case int64:
		return x
	case uint8:
		return AsInt64(x)
	case uint16:
		return AsInt64(x)
	case uint32:
		return AsInt64(x)
	case uint64:
		return AsInt64(x)
	case float32:
		return AsInt64(x)
	case float64:
		return AsInt64(x)
	default:
		panic("cannot convert to int64")
	}
}

func AsBigInt(x interface{}) *value.BigInt {
	switch x := x.(type) {
	case int8:
		return value.NewBigIntFromInt64(int64(x))
	case int16:
		return value.NewBigIntFromInt64(int64(x))
	case int32:
		return value.NewBigIntFromInt64(int64(x))
	case int64:
		return value.NewBigIntFromInt64(int64(x))
	case uint8:
		return value.NewBigIntFromInt64(int64(x))
	case uint16:
		return value.NewBigIntFromInt64(int64(x))
	case uint32:
		return value.NewBigIntFromInt64(int64(x))
	case uint64:
		return value.NewBigIntFromInt64(int64(x))
	case float32:
		return value.NewBigIntFromInt64(int64(x))
	case float64:
		return value.NewBigIntFromInt64(int64(x))
	default:
		panic("cannot convert to BigInt")
	}
}

func AsRatio(x interface{}) *value.Ratio {
	switch x := x.(type) {
	case int8:
		return value.NewRatio(int64(x), 1)
	case int16:
		return value.NewRatio(int64(x), 1)
	case int32:
		return value.NewRatio(int64(x), 1)
	case int64:
		return value.NewRatio(int64(x), 1)
	case uint8:
		return value.NewRatio(int64(x), 1)
	case uint16:
		return value.NewRatio(int64(x), 1)
	case uint32:
		return value.NewRatio(int64(x), 1)
	case uint64:
		return value.NewRatio(int64(x), 1)
	case float32:
		return value.NewRatio(int64(x), 1)
	case float64:
		return value.NewRatio(int64(x), 1)
	default:
		panic("cannot convert to Ratio")
	}
}

func AsBigDecimal(x interface{}) *value.BigDecimal {
	switch x := x.(type) {
	case int8:
		return value.NewBigDecimalFromFloat64(float64(x))
	case int16:
		return value.NewBigDecimalFromFloat64(float64(x))
	case int32:
		return value.NewBigDecimalFromFloat64(float64(x))
	case int64:
		return value.NewBigDecimalFromFloat64(float64(x))
	case uint8:
		return value.NewBigDecimalFromFloat64(float64(x))
	case uint16:
		return value.NewBigDecimalFromFloat64(float64(x))
	case uint32:
		return value.NewBigDecimalFromFloat64(float64(x))
	case uint64:
		return value.NewBigDecimalFromFloat64(float64(x))
	case float32:
		return value.NewBigDecimalFromFloat64(float64(x))
	case float64:
		return value.NewBigDecimalFromFloat64(float64(x))
	default:
		panic("cannot convert to BigDecimal")
	}
}

func AsFloat64(x interface{}) float64 {
	switch x := x.(type) {
	case int8:
		return AsFloat64(x)
	case int16:
		return AsFloat64(x)
	case int32:
		return AsFloat64(x)
	case int64:
		return AsFloat64(x)
	case uint8:
		return AsFloat64(x)
	case uint16:
		return AsFloat64(x)
	case uint32:
		return AsFloat64(x)
	case uint64:
		return AsFloat64(x)
	case float32:
		return AsFloat64(x)
	case float64:
		return x
	default:
		panic("cannot convert to float64")
	}
}