package errortypes

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
Instead of checking the exact string of the error,
we are doing a type assertion on the error to see if it is a `BadStatusError`.

This reflects our desire for the kind of error clearer.
Assuming the assertion passes we can then check the properties of the error are correct.
*/
type BadStatusError struct {
	URL    string
	Status int
}

func (b BadStatusError) Error() string {
	return fmt.Sprintf("did not get 200 from %s, got %d", b.URL, b.Status)
}

// DumbGetter will get the string body of url if it gets a 200
// Function has become simpler, it's no longer concerned with the intricacies of an error string,
// it just creates a BadStatusError.
func DumbGetter(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("problem fetching from %s, %v", url, err)
	}

	if res.StatusCode != http.StatusOK {
		return "", BadStatusError{URL: url, Status: res.StatusCode}
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body) // ignoring err for brevity
	return string(body), nil
}

/*
Our tests now reflect (and document) what a user of our code could do if
they decided they wanted to do some more sophisticated error handling than just logging.
Just do a type assertion and then you get easy access to the properties of the error.

It is still "just" an error, so if they choose to they can pass it up the call stack or log it like any other error.
*/
func TestDumbGetter(t *testing.T) {
	t.Run("when you don't get a 200, you get a status error", func(t *testing.T) {
		svr := httptest.NewServer(http.HandlerFunc(
			func(res http.ResponseWriter, req *http.Request) {
				res.WriteHeader(http.StatusTeapot)
			}))
		defer svr.Close()

		_, err := DumbGetter(svr.URL)
		if err == nil {
			t.Fatal("expected an error")
		}

		// In this case we are using `errors.As` to try and extract our error into our custom type.
		// It returns a `bool` to denote success and extracts it into `got` for us.
		var got BadStatusError
		isBadStatusError := errors.As(err, &got)
		want := BadStatusError{URL: svr.URL, Status: http.StatusTeapot}
		if !isBadStatusError {
			t.Fatalf("was not a BadStatusError, got %T", err)
		}

		if got != want {
			t.Errorf(`got "%v", want "%v"`, got, want)
		}
	})
}
