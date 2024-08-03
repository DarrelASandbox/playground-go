package plannertest

import (
	"testing"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/planner/internal/expect"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/recipe"
)

func AssertDoesntHaveRecipe(t *testing.T, recipes recipe.Recipes, expected recipe.Recipe) {
	t.Helper()
	_, found := recipes.FindByName(expected.Name)
	expect.False(t, found)
}

func AssertHasRecipe(t *testing.T, recipes recipe.Recipes, expected recipe.Recipe) {
	t.Helper()
	_, found := recipes.FindByName(expected.Name)
	expect.True(t, found)
}
