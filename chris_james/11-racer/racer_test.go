package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// `httptest.NewServer` takes an `http.HandlerFunc` which we are sending in via an anonymous function.
// `http.HandlerFunc` is a type that looks like this: `type HandlerFunc func(ResponseWriter, *Request)`.
// We are wrapping it in an httptest.NewServer which makes it easier to use with testing,
// as it finds an open port to listen on and then you can close it when you're done with your test.
func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	slowServer.Close()
	fastServer.Close()
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
