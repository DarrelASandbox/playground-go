package specs_greet_test

import (
	"testing"

	specs_greet "github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/specifications"
)

/*
cannot use specs_greet.Greet (value of type func(name string) string) as specifications.Greeter value in argument to specifications.GreetSpecification: func(name string) string does not implement specifications.Greeter (missing method Greet)
*/
func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specs_greet.Greet)
}
