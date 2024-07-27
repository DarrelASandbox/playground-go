package recipe

import "github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/ingredients"

type MealType int

type Recipe struct {
	Name        string
	MealType    MealType
	Ingredients []ingredients.Ingredient
	Description string
}
