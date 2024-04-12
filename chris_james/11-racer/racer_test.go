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
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))

	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

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
