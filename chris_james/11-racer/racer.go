package racer

import (
	"net/http"
)

// `select` helps us synchronize processes really easily and clearly.
// You can wait for values to be sent to a channel with `myVar := <-ch`.
// This is a blocking call, as you're waiting for a value.
// `select` allows you to wait on multiple channels.
// The first one to send a value "wins" and the code underneath the `case` is executed.
// We use `ping` in our `select` to set up two channels, one for each of our URLs.
// Whichever one writes to its channel first will have its code executed in the `select`, which results in its URL being returned (and being the winner).
func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

// Why struct{} and not another type like a bool?
// Well, a chan struct{} is the smallest data type available from a memory perspective so we get no allocation versus a bool.
// Since we are closing and not sending anything on the chan, why allocate anything?
func ping(url string) chan struct{} {

	// Notice how we have to use `make` when creating a channel; rather than say `var ch chan struct{}`.
	// When you use var the variable will be initialized with the "zero" value of the type.
	// So for `string` it is `""`, `int` it is 0, etc.

	// For channels the zero value is `nil` and if you try and
	// send to it with `<-` it will block forever because you cannot send to `nil` channels
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
