package recipe

import "github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/ingredients"

type MealType int

const (
	Breakfast MealType = iota
	Lunch
	Dinner
)

type Recipe struct {
	Name        string
	MealType    MealType
	Ingredients []ingredients.Ingredient
	Description string
}

type Recipes []Recipe

func (m MealType) String() string {
	return [...]string{"Breakfast", "Lunch", "Dinner"}[m]
}

func (r Recipes) FindByName(name string) (Recipe, bool) {
	for _, recipe := range r {
		if recipe.Name == name {
			return recipe, true
		}
	}
	return Recipe{}, false
}
