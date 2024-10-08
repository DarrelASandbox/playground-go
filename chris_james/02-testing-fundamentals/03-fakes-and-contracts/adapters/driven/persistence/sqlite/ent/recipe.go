// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/adapters/driven/persistence/sqlite/ent/recipe"
)

// Recipe is the model entity for the Recipe schema.
type Recipe struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// MealType holds the value of the "meal_type" field.
	MealType int `json:"meal_type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RecipeQuery when eager-loading is set.
	Edges        RecipeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RecipeEdges holds the relations/edges for other nodes in the graph.
type RecipeEdges struct {
	// Recipeingredient holds the value of the recipeingredient edge.
	Recipeingredient []*RecipeIngredient `json:"recipeingredient,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RecipeingredientOrErr returns the Recipeingredient value or an error if the edge
// was not loaded in eager-loading.
func (e RecipeEdges) RecipeingredientOrErr() ([]*RecipeIngredient, error) {
	if e.loadedTypes[0] {
		return e.Recipeingredient, nil
	}
	return nil, &NotLoadedError{edge: "recipeingredient"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Recipe) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case recipe.FieldID, recipe.FieldMealType:
			values[i] = new(sql.NullInt64)
		case recipe.FieldName, recipe.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Recipe fields.
func (r *Recipe) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case recipe.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case recipe.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case recipe.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = value.String
			}
		case recipe.FieldMealType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field meal_type", values[i])
			} else if value.Valid {
				r.MealType = int(value.Int64)
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Recipe.
// This includes values selected through modifiers, order, etc.
func (r *Recipe) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryRecipeingredient queries the "recipeingredient" edge of the Recipe entity.
func (r *Recipe) QueryRecipeingredient() *RecipeIngredientQuery {
	return NewRecipeClient(r.config).QueryRecipeingredient(r)
}

// Update returns a builder for updating this Recipe.
// Note that you need to call Recipe.Unwrap() before calling this method if this Recipe
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Recipe) Update() *RecipeUpdateOne {
	return NewRecipeClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Recipe entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Recipe) Unwrap() *Recipe {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Recipe is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Recipe) String() string {
	var builder strings.Builder
	builder.WriteString("Recipe(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("name=")
	builder.WriteString(r.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(r.Description)
	builder.WriteString(", ")
	builder.WriteString("meal_type=")
	builder.WriteString(fmt.Sprintf("%v", r.MealType))
	builder.WriteByte(')')
	return builder.String()
}

// Recipes is a parsable slice of Recipe.
type Recipes []*Recipe
