package reflection

import "reflect"

// walk recursively iterates over elements in a data structure
// and applies the provided function `fn` to all string values found.
func walk(x interface{}, fn func(input string)) {
	val := getValue(x) // Retrieve the reflect.Value of the provided interface{}.

	numberOfValues := 0                  // Initialize count of sub-values we need to iterate over.
	var getField func(int) reflect.Value // Function to retrieve sub-values by index.

	switch val.Kind() {
	case reflect.String:
		fn(val.String()) // If it's a string, apply the function directly.
	case reflect.Struct:
		numberOfValues = val.NumField() // If it's a struct, get the number of fields.
		getField = val.Field            // Function to access fields by index.
	case reflect.Slice:
		numberOfValues = val.Len() // If it's a slice, get the length.
		getField = val.Index       // Function to access elements by index.
	}

	// Recursive iteration over fields or elements if they exist.
	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn) // Apply walk recursively to each sub-value.
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
