package sqlite_test

import (
	"testing"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/adapters/driven/persistence/sqlite"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/planner"
)

func TestRecipeStore(t *testing.T) {
	planner.RecipeBookContract{
		NewBook: func() planner.RecipeBook { return sqlite.NewRecipeStore(sqlite.NewSQLiteClient()) },
	}.Test(t)
}
