package di

import (
	"fmt"
	"io"
	"os"
)

// `fmt.Fprintf` is like `fmt.Printf` but instead takes a `Writer` to send the string to,
// whereas `fmt.Printf` defaults to stdout.
// `fmt.Fprintf` allows you to pass in an `io.Writer` which we know both `os.Stdout` and `bytes.Buffer` implement.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func di() {
	Greet(os.Stdout, "Ellery")
}
