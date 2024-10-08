package planner

import (
	"context"
	"fmt"
	"time"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/ingredients"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/recipe"
)

type Planner struct {
	recipeBook RecipeBook
	pantry     Pantry
}

type ErrorMissingIngredients struct {
	MissingIngredients ingredients.Ingredients
}

func New(recipes RecipeBook, ingredientStore Pantry) *Planner {
	return &Planner{recipeBook: recipes, pantry: ingredientStore}
}

func (p Planner) ScheduleMeal(ctx context.Context, r recipe.Recipe, _ time.Time) error {
	availableIngredients, err := p.pantry.GetIngredients(ctx)
	if err != nil {
		return err
	}

	if hasIngredients, missing := haveIngredients(availableIngredients, r); !hasIngredients {
		return ErrorMissingIngredients{
			MissingIngredients: missing,
		}
	}

	return p.pantry.Remove(ctx, r.Ingredients...)
}

// returns slice of missing ingredients
func haveIngredients(availableIngredients ingredients.Ingredients, recipe recipe.Recipe) (hasIngredients bool, missing ingredients.Ingredients) {
	for _, ingredient := range recipe.Ingredients {
		if !availableIngredients.Has(ingredient) {
			missing = append(missing, ingredient)
		}
	}

	if len(missing) > 0 {
		return false, missing
	}

	return true, nil
}

func (p Planner) SuggestRecipes(ctx context.Context) (recipe.Recipes, error) {
	availableIngredients, err := p.pantry.GetIngredients(ctx)
	if err != nil {
		return nil, err
	}

	recipes, err := p.recipeBook.GetRecipes(ctx)
	if err != nil {
		return nil, err
	}

	var suggestions []recipe.Recipe
	for _, r := range recipes {
		if hasIngredients, _ := haveIngredients(availableIngredients, r); hasIngredients {
			suggestions = append(suggestions, r)
		}
	}

	return suggestions, nil
}

func (e ErrorMissingIngredients) Error() string {
	return fmt.Sprintf("missing ingredients: %v", e.MissingIngredients)
}
