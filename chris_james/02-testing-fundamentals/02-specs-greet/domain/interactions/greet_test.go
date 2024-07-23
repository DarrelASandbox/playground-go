package interactions_test

import (
	"testing"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/domain/interactions"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/specifications"
	"github.com/alecthomas/assert/v2"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(interactions.Greet))

	t.Run("default name to world if it's an empty string", func(t *testing.T) {
		assert.Equal(t, "Hello, World", interactions.Greet(""))
	})
}
