package generics

import "testing"

/*
Ideally, we don't want to have to make specific AssertX functions for every type we ever deal with.
We'd like to be able to have one AssertEqual function that works with any type but
does not let you compare apples and oranges.
*/

func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func AssertNotEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got == want {
		if got == want {
			t.Errorf("didn't want %d", got)
		}
	}
}
