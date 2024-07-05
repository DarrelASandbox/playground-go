package generics

import "testing"

func AssertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

// Provide type parameters by describing the generic type and give it a label
func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		if got == want {
			t.Errorf("didn't want %v", got)
		}
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
