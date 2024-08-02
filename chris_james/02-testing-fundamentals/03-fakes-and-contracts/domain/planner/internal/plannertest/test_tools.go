package plannertest

import (
	"math/rand"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/ingredients"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/recipe"
	"github.com/google/uuid"
)

func RandomIngredients() []ingredients.Ingredient {
	return []ingredients.Ingredient{RandomIngredient(), RandomIngredient(), RandomIngredient()}
}

func RandomIngredient() ingredients.Ingredient {
	return ingredients.Ingredient{Name: uuid.New().String(), Quantity: uint(rand.Intn(10)) + 1}
}

func RandomRecipe() recipe.Recipe {
	return recipe.Recipe{Name: uuid.New().String(), Ingredients: RandomIngredients()}
}

func RandomRecipes() []recipe.Recipe {
	return []recipe.Recipe{RandomRecipe(), RandomRecipe(), RandomRecipe()}
}
