package apiquery

import (
	"github.com/anthropics/anthropic-sdk-go/packages/param"
	"reflect"
)

func (e *encoder) newRichFieldTypeEncoder(t reflect.Type, underlyingValueIdx []int) encoderFunc {
	underlying := t.FieldByIndex(underlyingValueIdx)
	primitiveEncoder := e.newPrimitiveTypeEncoder(underlying.Type)
	return func(key string, value reflect.Value) []Pair {
		if fielder, ok := value.Interface().(param.Optional); ok && fielder.IsPresent() {
			return primitiveEncoder(key, value.FieldByIndex(underlyingValueIdx))
		} else if ok && fielder.IsNull() {
			return []Pair{{key, "null"}}
		}
		return nil
	}
}
