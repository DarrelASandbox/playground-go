package reflection

import "reflect"

/*
This code is very unsafe and very naive,
but remember: our goal when we are in "red" (the tests failing) is to write the smallest amount of code possible.
We then write more tests to address our concerns.

We need to use reflection to have a look at x and try and look at its properties.

make some very optimistic assumptions about the value passed in:
  - We look at the first and only field. However, there may be no fields at all, which would cause a panic.
  - We then call String(), which returns the underlying value as a string.
    However, this would be wrong if the field was something other than a string.
*/
func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fn(field.String())
	}
}
