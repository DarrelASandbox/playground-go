package di

import (
	"bytes"
	"fmt"
)

// `fmt.Fprintf` is like `fmt.Printf` but instead takes a `Writer` to send the string to,
// whereas `fmt.Printf` defaults to stdout.
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
