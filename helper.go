package qibo

import (
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

func Int32ToString(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

// IsArgNil check type is null
func IsArgNil(i interface{}) bool {
	// TODO: need to optimize di function
	r := reflect.ValueOf(i)
	switch r.Kind() {
	case reflect.Slice:
		return r.Len() == 0
	case reflect.String:
		return r.String() == ""
	case reflect.Int:
		return r.Int() == 0
	case reflect.Int32:
		return r.Int() == 0
	case reflect.Int64:
		return r.Int() == 0
	case reflect.Float32:
		return r.Float() == 0
	case reflect.Float64:
		return r.Float() == 0
	default:
		return false
	}
}
