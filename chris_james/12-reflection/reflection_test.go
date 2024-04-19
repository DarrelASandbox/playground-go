package reflection

import (
	"reflect"
	"testing"
)

// We want to store a slice of strings (`got`) which stores which strings were passed into `fn` by `walk`.
// Often in previous chapters, we have made dedicated types for this to spy on function/method invocations
// but in this case, we can just pass in an anonymous function for `fn` that closes over `got`.
func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	// Anonymous struct
	x := struct {
		Name string
	}{expected}

	// Call walk with x and the spy and
	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}

	if got[0] != expected {
		t.Errorf("got %q, want %q", got[0], expected)
	}
}

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
	field := val.Field(0)
	fn(field.String())
}
