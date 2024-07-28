package planner

import (
	"context"
	"testing"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/ingredients"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/planner/internal/expect"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/recipe"
)

type RecipeBook interface {
	GetRecipes(context.Context) ([]recipe.Recipe, error)
	AddRecipes(context.Context, ...recipe.Recipe) error
}

type RecipeBookContract struct {
	NewBook func() RecipeBook
}

func (r RecipeBookContract) Test(t *testing.T) {
	t.Run("it returns what is put in", func(t *testing.T) {
		var (
			ctx         = context.Background()
			someRecipes = []recipe.Recipe{
				{Name: "Banana Pancakes",
					Description: "A delicious treat",
					MealType:    recipe.Breakfast,
					Ingredients: []ingredients.Ingredient{{Name: "Bananas", Quantity: 2}},
				},
				{
					Name:        "Pasta",
					Description: "Plain pasta, delicious",
					MealType:    recipe.Lunch,
					Ingredients: []ingredients.Ingredient{{Name: "Bananas", Quantity: 2}},
				},
			}
			sut = r.NewBook()
		)

		expect.NoErr(t, sut.AddRecipes(ctx, someRecipes...))
		got, err := sut.GetRecipes(ctx)
		expect.NoErr(t, err)
		expect.DeepEqual(t, got, someRecipes)
	})
}
