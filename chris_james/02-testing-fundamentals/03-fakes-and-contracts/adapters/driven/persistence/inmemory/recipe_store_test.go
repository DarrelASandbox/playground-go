package inmemory_test

import (
	"testing"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/adapters/driven/persistence/inmemory"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/planner"
)

func TestInMemoryRecipeStore(t *testing.T) {
	planner.RecipeBookContract{
		NewBook: func() planner.RecipeBook { return inmemory.NewRecipeStore() }}.Test(t)
}
