package reflection

import "testing"

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
}

func walk(x interface{}, fn func(input string)) {}
