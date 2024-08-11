// Code generated by ent, DO NOT EDIT.

package ingredient

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the ingredient type in the database.
	Label = "ingredient"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldVegan holds the string denoting the vegan field in the database.
	FieldVegan = "vegan"
	// EdgePantry holds the string denoting the pantry edge name in mutations.
	EdgePantry = "pantry"
	// EdgeRecipeingredient holds the string denoting the recipeingredient edge name in mutations.
	EdgeRecipeingredient = "recipeingredient"
	// Table holds the table name of the ingredient in the database.
	Table = "ingredients"
	// PantryTable is the table that holds the pantry relation/edge.
	PantryTable = "ingredients"
	// PantryInverseTable is the table name for the Pantry entity.
	// It exists in this package in order to avoid circular dependency with the "pantry" package.
	PantryInverseTable = "pantries"
	// PantryColumn is the table column denoting the pantry relation/edge.
	PantryColumn = "pantry_ingredient"
	// RecipeingredientTable is the table that holds the recipeingredient relation/edge. The primary key declared below.
	RecipeingredientTable = "recipe_ingredient_ingredient"
	// RecipeingredientInverseTable is the table name for the RecipeIngredient entity.
	// It exists in this package in order to avoid circular dependency with the "recipeingredient" package.
	RecipeingredientInverseTable = "recipe_ingredients"
)

// Columns holds all SQL columns for ingredient fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldVegan,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "ingredients"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"pantry_ingredient",
}

var (
	// RecipeingredientPrimaryKey and RecipeingredientColumn2 are the table columns denoting the
	// primary key for the recipeingredient relation (M2M).
	RecipeingredientPrimaryKey = []string{"recipe_ingredient_id", "ingredient_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultVegan holds the default value on creation for the "vegan" field.
	DefaultVegan bool
)

// OrderOption defines the ordering options for the Ingredient queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByVegan orders the results by the vegan field.
func ByVegan(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVegan, opts...).ToFunc()
}

// ByPantryField orders the results by pantry field.
func ByPantryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPantryStep(), sql.OrderByField(field, opts...))
	}
}

// ByRecipeingredientCount orders the results by recipeingredient count.
func ByRecipeingredientCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRecipeingredientStep(), opts...)
	}
}

// ByRecipeingredient orders the results by recipeingredient terms.
func ByRecipeingredient(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRecipeingredientStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPantryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PantryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, PantryTable, PantryColumn),
	)
}
func newRecipeingredientStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RecipeingredientInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, RecipeingredientTable, RecipeingredientPrimaryKey...),
	)
}