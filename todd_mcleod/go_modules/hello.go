// Creating a new module
package hello

import "rsc.io/quote" // Adding a dependency

func Hello() string {
	println(quote.Hello())
	return "Hello, world."
}
