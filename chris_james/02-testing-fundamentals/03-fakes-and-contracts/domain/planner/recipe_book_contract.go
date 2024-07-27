package planner

import (
	"context"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/recipe"
)

type RecipeBook interface {
	GetRecipes(context.Context) ([]recipe.Recipe, error)
	AddRecipes(context.Context, ...recipe.Recipe) error
}
