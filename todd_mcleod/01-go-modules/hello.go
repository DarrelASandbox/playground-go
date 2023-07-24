// Creating a new module
package hello

// @rsc probably needs to update their app engine server
// https://github.com/golang/go/issues/40452
// https://github.com/rsc/quote/issues/4#issuecomment-665420354
import "rsc.io/quote" // Adding a dependency

func Hello() string {
	println(quote.Hello())
	return "Hello, world."
}
