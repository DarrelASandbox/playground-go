package cancelreader

import (
	"context"
	"strings"
	"testing"
)

/*
--- FAIL: TestContextAwareReader (0.00s)
    --- FAIL: TestContextAwareReader/stops_reading_when_cancelled (0.00s)
        context_aware_reader_test.go:47: expected an error after cancellation but didn't get one
        context_aware_reader_test.go:51: expected 0 bytes to be read after cancellation but 3 were read
*/
func TestContextAwareReader(t *testing.T) {
	// We want to be able to compose an `io.Reader` with a `context.Context`.
	t.Run("behaves like a normal reader", func(t *testing.T) {
		rdr := NewCancellableReader(context.Background(), strings.NewReader("123456"))
		got := make([]byte, 3)
		_, err := rdr.Read(got)
		if err != nil {
			t.Fatal(err)
		}

		assertBufferHas(t, got, "123")
		_, err = rdr.Read(got)
		if err != nil {
			t.Fatal(err)
		}

		assertBufferHas(t, got, "456")
	})

	t.Run("stops reading when cancelled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		rdr := NewCancellableReader(ctx, strings.NewReader("123456"))
		got := make([]byte, 3)
		_, err := rdr.Read(got)
		if err != nil {
			t.Fatal(err)
		}

		assertBufferHas(t, got, "123")
		cancel()
		n, err := rdr.Read(got)
		if err == nil {
			t.Error("expected an error after cancellation but didn't get one")
		}

		if n > 0 {
			t.Errorf("expected 0 bytes to be read after cancellation but %d were read", n)
		}
	})
}

func assertBufferHas(t testing.TB, buf []byte, want string) {
	t.Helper()
	got := string(buf)
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
