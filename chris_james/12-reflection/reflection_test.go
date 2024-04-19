package reflection

import (
	"reflect"
	"testing"
)

/*
This test function systematically verifies that the walk function behaves correctly for different structures and values.
By defining expected outcomes and comparing them with actual results,
it ensures that walk meets its specification in handling various types of inputs.
This approach is common in unit testing, where the functionality of individual components is validated in isolation.
*/

/*
We're going to be writing a number of tests where we pass in different values and checking the array of strings that `fn` was called with.

We should refactor our test into a table based test to make this easier to continue testing new scenarios.
*/
func TestWalk(t *testing.T) {

	// A slice of anonymous struct instances, each representing a test case.
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct{ Name string }{"Chris"}, []string{"Chris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
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
