package interactions_test

import (
	"testing"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/domain/interactions"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/specifications"
)

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(t, specifications.CurseAdapter(interactions.Curse))
}
