package specs_greet

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
