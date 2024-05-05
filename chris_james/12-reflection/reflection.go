package reflection

import "reflect"

// walk recursively iterates over elements in a data structure
// and applies the provided function `fn` to all string values found.
func walk(x interface{}, fn func(input string)) {
	val := getValue(x) // Retrieve the reflect.Value of the provided interface{}.

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String()) // If it's a string, apply the function directly.

	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}

	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}

	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}

	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

// getValue extracts the underlying value of `x` if `x` is a pointer,
// or simply returns the reflect.Value of `x` if not a pointer.
func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x) // Convert interface to reflect.Value.

	// Handle pointers by retrieving the value they point to.
	if val.Kind() == reflect.Pointer {
		val = val.Elem() // Extract underlying value from pointer.
	}

	return val // Return the possibly dereferenced value.
}
