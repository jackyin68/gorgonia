package tensor

import (
	"reflect"

	"github.com/chewxy/vecf32"
	"github.com/chewxy/vecf64"
	"github.com/pkg/errors"
)

/*
GENERATED FILE. DO NOT EDIT
*/

// Range creates a ranged array with a given type. It panics if the Dtype is not supported or does not represent a naturally orderable type (strings, pointers etc)
// Do note that the range algorithm is very simple, and simply does increments or decrements of 1. This means for floating point types
// you're not able to create a range with a 0.1 increment step, and for complex number types, the imaginary part will always be 0i
func Range(dt Dtype, start, end int) interface{} {
	size := end - start
	incr := true
	if start > end {
		incr = false
		size = start - end
	}

	if size < 0 {
		panic("Cannot create a range that is negative in size")
	}
	switch dt.Kind() {
	case reflect.Int:
		retVal := make([]int, size)
		for i, v := 0, int(start); i < size; i++ {
			retVal[i] = v
			if incr {
				v++
			} else {
				v--
			}
		}
		return retVal
	case reflect.Int8:
		retVal := make([]int8, size)
		for i, v := 0, int8(start); i < size; i++ {
			retVal[i] = v
			if incr {
				v++
			} else {
				v--
			}
		}
		return retVal
	case reflect.Int16:
		retVal := make([]int16, size)
		for i, v := 0, int16(start); i < size; i++ {
			retVal[i] = v
			if incr {
				v++
			} else {
				v--
			}
		}
		return retVal
	case reflect.Int32:
		retVal := make([]int32, size)
		for i, v := 0, int32(start); i < size; i++ {
			retVal[i] = v
			if incr {
				v++
			} else {
				v--
			}
		}
		return retVal
	case reflect.Int64:
		retVal := make([]int64, size)
		for i, v := 0, int64(start); i < size; i++ {
			retVal[i] = v
			if incr {
				v++
			} else {
				v--
			}
		}
		return retVal
	case reflect.Uint:
		retVal := make([]uint, size)
		for i, v := 0, uint(start); i < size; i++ {
			retVal[i] = v
			if incr {
				v++
			} else {
				v--
			}
		}
		return retVal
	case reflect.Uint8:
		retVal := make([]uint8, size)
		for i, v := 0, uint8(start); i < size; i++ {
			retVal[i] = v
			if incr {
				v++
			} else {
				v--
			}
		}
		return retVal
	case reflect.Uint16:
		retVal := make([]uint16, size)
		for i, v := 0, uint16(start); i < size; i++ {
			retVal[i] = v
			if incr {
				v++
			} else {
				v--
			}
		}
		return retVal
	case reflect.Uint32:
		retVal := make([]uint32, size)
		for i, v := 0, uint32(start); i < size; i++ {
			retVal[i] = v
			if incr {
				v++
			} else {
				v--
			}
		}
		return retVal
	case reflect.Uint64:
		retVal := make([]uint64, size)
		for i, v := 0, uint64(start); i < size; i++ {
			retVal[i] = v
			if incr {
				v++
			} else {
				v--
			}
		}
		return retVal
	case reflect.Float32:
		return vecf32.Range(start, end)
	case reflect.Float64:
		return vecf64.Range(start, end)
	case reflect.Complex64:
		retVal := make([]complex64, size)
		for i, v := 0, complex(float32(start), float32(0.0)); i < size; i++ {
			retVal[i] = v
			if incr {
				v++
			} else {
				v--
			}
		}
		return retVal
	case reflect.Complex128:
		retVal := make([]complex128, size)
		for i, v := 0, complex(float64(start), float64(0.0)); i < size; i++ {
			retVal[i] = v
			if incr {
				v++
			} else {
				v--
			}
		}
		return retVal
	default:
		err := errors.Errorf("Unrangeable Type %v", dt)
		panic(err)
	}
}
