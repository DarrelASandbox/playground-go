package planner_test

import (
	"context"
	"testing"
	"time"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/adapters/driven/persistence/inmemory"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/planner"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/planner/internal/expect"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/planner/internal/plannertest"
)

type RecipePlannerTest struct {
	CreateDependencies func() (planner.RecipeBook, planner.Pantry, Cleanup)
}

type Cleanup func()

func TestRecipePlanner(t *testing.T) {
	// for local, snappy integration test with a fake (which we can be confident is correct due
	// to it  conforming to the store contract)
	t.Run("with in memory store", func(t *testing.T) {
		RecipePlannerTest{
			CreateDependencies: func() (planner.RecipeBook, planner.Pantry, Cleanup) {
				return inmemory.NewRecipeStore(), inmemory.NewPantry(), func() {
					// nothing to clean up
				}
			},
		}.Test(t)
	})
}

func (r RecipePlannerTest) Test(t *testing.T) {
	t.Run("planning meals", func(t *testing.T) {
		t.Run("happy path, have ingredients for a recipe, schedule it, update pantry", func(t *testing.T) {
			var (
				ctx                          = context.Background()
				lasagna                      = plannertest.RandomRecipe()
				recipeBook, pantry, teardown = r.CreateDependencies()
				sut                          = planner.New(recipeBook, pantry)
			)

			t.Cleanup(teardown)

			expect.NoErr(t, recipeBook.AddRecipes(ctx, lasagna))
			expect.NoErr(t, pantry.Store(ctx, lasagna.Ingredients...))
			expect.NoErr(t, sut.ScheduleMeal(ctx, lasagna, time.Now()))
			remainingIngredients, err := pantry.GetIngredients(ctx)
			expect.NoErr(t, err)
			expect.Equal(t, 0, len(remainingIngredients))
		})

		t.Run("returns a missing ingredients error if you try to schedule a meal without all the ingredients", func(t *testing.T) {})
		t.Run("when recipeBook fails to get ingredients, we get an error", func(t *testing.T) {})
		t.Run("returns the specific ingredients missing if you try to schedule a meal with some missing ingredients", func(t *testing.T) {})
	})

	t.Run("suggesting recipes", func(t *testing.T) {
		t.Run("if we have the ingredients for a recipe we can make it", func(t *testing.T) {})
		t.Run("if we have ingredients for 2 recipes, we can make both", func(t *testing.T) {})
	})
}
