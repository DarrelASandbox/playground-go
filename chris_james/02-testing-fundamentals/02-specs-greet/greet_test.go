package specs_greet_test

import (
	"testing"

	specs_greet "github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(specs_greet.Greet))
}
