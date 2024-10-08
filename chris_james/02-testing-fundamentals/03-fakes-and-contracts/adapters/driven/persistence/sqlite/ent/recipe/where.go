// Code generated by ent, DO NOT EDIT.

package recipe

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/adapters/driven/persistence/sqlite/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Recipe {
	return predicate.Recipe(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Recipe {
	return predicate.Recipe(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Recipe {
	return predicate.Recipe(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Recipe {
	return predicate.Recipe(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Recipe {
	return predicate.Recipe(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Recipe {
	return predicate.Recipe(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Recipe {
	return predicate.Recipe(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Recipe {
	return predicate.Recipe(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Recipe {
	return predicate.Recipe(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldEQ(FieldDescription, v))
}

// MealType applies equality check predicate on the "meal_type" field. It's identical to MealTypeEQ.
func MealType(v int) predicate.Recipe {
	return predicate.Recipe(sql.FieldEQ(FieldMealType, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Recipe {
	return predicate.Recipe(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Recipe {
	return predicate.Recipe(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Recipe {
	return predicate.Recipe(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Recipe {
	return predicate.Recipe(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Recipe {
	return predicate.Recipe(sql.FieldContainsFold(FieldDescription, v))
}

// MealTypeEQ applies the EQ predicate on the "meal_type" field.
func MealTypeEQ(v int) predicate.Recipe {
	return predicate.Recipe(sql.FieldEQ(FieldMealType, v))
}

// MealTypeNEQ applies the NEQ predicate on the "meal_type" field.
func MealTypeNEQ(v int) predicate.Recipe {
	return predicate.Recipe(sql.FieldNEQ(FieldMealType, v))
}

// MealTypeIn applies the In predicate on the "meal_type" field.
func MealTypeIn(vs ...int) predicate.Recipe {
	return predicate.Recipe(sql.FieldIn(FieldMealType, vs...))
}

// MealTypeNotIn applies the NotIn predicate on the "meal_type" field.
func MealTypeNotIn(vs ...int) predicate.Recipe {
	return predicate.Recipe(sql.FieldNotIn(FieldMealType, vs...))
}

// MealTypeGT applies the GT predicate on the "meal_type" field.
func MealTypeGT(v int) predicate.Recipe {
	return predicate.Recipe(sql.FieldGT(FieldMealType, v))
}

// MealTypeGTE applies the GTE predicate on the "meal_type" field.
func MealTypeGTE(v int) predicate.Recipe {
	return predicate.Recipe(sql.FieldGTE(FieldMealType, v))
}

// MealTypeLT applies the LT predicate on the "meal_type" field.
func MealTypeLT(v int) predicate.Recipe {
	return predicate.Recipe(sql.FieldLT(FieldMealType, v))
}

// MealTypeLTE applies the LTE predicate on the "meal_type" field.
func MealTypeLTE(v int) predicate.Recipe {
	return predicate.Recipe(sql.FieldLTE(FieldMealType, v))
}

// HasRecipeingredient applies the HasEdge predicate on the "recipeingredient" edge.
func HasRecipeingredient() predicate.Recipe {
	return predicate.Recipe(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RecipeingredientTable, RecipeingredientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRecipeingredientWith applies the HasEdge predicate on the "recipeingredient" edge with a given conditions (other predicates).
func HasRecipeingredientWith(preds ...predicate.RecipeIngredient) predicate.Recipe {
	return predicate.Recipe(func(s *sql.Selector) {
		step := newRecipeingredientStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Recipe) predicate.Recipe {
	return predicate.Recipe(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Recipe) predicate.Recipe {
	return predicate.Recipe(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Recipe) predicate.Recipe {
	return predicate.Recipe(sql.NotPredicates(p))
}
