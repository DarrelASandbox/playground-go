// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/adapters/driven/persistence/sqlite/ent/predicate"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/adapters/driven/persistence/sqlite/ent/recipe"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/adapters/driven/persistence/sqlite/ent/recipeingredient"
)

// RecipeUpdate is the builder for updating Recipe entities.
type RecipeUpdate struct {
	config
	hooks    []Hook
	mutation *RecipeMutation
}

// Where appends a list predicates to the RecipeUpdate builder.
func (ru *RecipeUpdate) Where(ps ...predicate.Recipe) *RecipeUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetName sets the "name" field.
func (ru *RecipeUpdate) SetName(s string) *RecipeUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ru *RecipeUpdate) SetNillableName(s *string) *RecipeUpdate {
	if s != nil {
		ru.SetName(*s)
	}
	return ru
}

// SetDescription sets the "description" field.
func (ru *RecipeUpdate) SetDescription(s string) *RecipeUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ru *RecipeUpdate) SetNillableDescription(s *string) *RecipeUpdate {
	if s != nil {
		ru.SetDescription(*s)
	}
	return ru
}

// SetMealType sets the "meal_type" field.
func (ru *RecipeUpdate) SetMealType(i int) *RecipeUpdate {
	ru.mutation.ResetMealType()
	ru.mutation.SetMealType(i)
	return ru
}

// SetNillableMealType sets the "meal_type" field if the given value is not nil.
func (ru *RecipeUpdate) SetNillableMealType(i *int) *RecipeUpdate {
	if i != nil {
		ru.SetMealType(*i)
	}
	return ru
}

// AddMealType adds i to the "meal_type" field.
func (ru *RecipeUpdate) AddMealType(i int) *RecipeUpdate {
	ru.mutation.AddMealType(i)
	return ru
}

// AddRecipeingredientIDs adds the "recipeingredient" edge to the RecipeIngredient entity by IDs.
func (ru *RecipeUpdate) AddRecipeingredientIDs(ids ...int) *RecipeUpdate {
	ru.mutation.AddRecipeingredientIDs(ids...)
	return ru
}

// AddRecipeingredient adds the "recipeingredient" edges to the RecipeIngredient entity.
func (ru *RecipeUpdate) AddRecipeingredient(r ...*RecipeIngredient) *RecipeUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddRecipeingredientIDs(ids...)
}

// Mutation returns the RecipeMutation object of the builder.
func (ru *RecipeUpdate) Mutation() *RecipeMutation {
	return ru.mutation
}

// ClearRecipeingredient clears all "recipeingredient" edges to the RecipeIngredient entity.
func (ru *RecipeUpdate) ClearRecipeingredient() *RecipeUpdate {
	ru.mutation.ClearRecipeingredient()
	return ru
}

// RemoveRecipeingredientIDs removes the "recipeingredient" edge to RecipeIngredient entities by IDs.
func (ru *RecipeUpdate) RemoveRecipeingredientIDs(ids ...int) *RecipeUpdate {
	ru.mutation.RemoveRecipeingredientIDs(ids...)
	return ru
}

// RemoveRecipeingredient removes "recipeingredient" edges to RecipeIngredient entities.
func (ru *RecipeUpdate) RemoveRecipeingredient(r ...*RecipeIngredient) *RecipeUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveRecipeingredientIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RecipeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RecipeUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RecipeUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RecipeUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RecipeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(recipe.Table, recipe.Columns, sqlgraph.NewFieldSpec(recipe.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(recipe.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.SetField(recipe.FieldDescription, field.TypeString, value)
	}
	if value, ok := ru.mutation.MealType(); ok {
		_spec.SetField(recipe.FieldMealType, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedMealType(); ok {
		_spec.AddField(recipe.FieldMealType, field.TypeInt, value)
	}
	if ru.mutation.RecipeingredientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.RecipeingredientTable,
			Columns: []string{recipe.RecipeingredientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipeingredient.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedRecipeingredientIDs(); len(nodes) > 0 && !ru.mutation.RecipeingredientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.RecipeingredientTable,
			Columns: []string{recipe.RecipeingredientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipeingredient.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RecipeingredientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.RecipeingredientTable,
			Columns: []string{recipe.RecipeingredientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipeingredient.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recipe.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RecipeUpdateOne is the builder for updating a single Recipe entity.
type RecipeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RecipeMutation
}

// SetName sets the "name" field.
func (ruo *RecipeUpdateOne) SetName(s string) *RecipeUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ruo *RecipeUpdateOne) SetNillableName(s *string) *RecipeUpdateOne {
	if s != nil {
		ruo.SetName(*s)
	}
	return ruo
}

// SetDescription sets the "description" field.
func (ruo *RecipeUpdateOne) SetDescription(s string) *RecipeUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ruo *RecipeUpdateOne) SetNillableDescription(s *string) *RecipeUpdateOne {
	if s != nil {
		ruo.SetDescription(*s)
	}
	return ruo
}

// SetMealType sets the "meal_type" field.
func (ruo *RecipeUpdateOne) SetMealType(i int) *RecipeUpdateOne {
	ruo.mutation.ResetMealType()
	ruo.mutation.SetMealType(i)
	return ruo
}

// SetNillableMealType sets the "meal_type" field if the given value is not nil.
func (ruo *RecipeUpdateOne) SetNillableMealType(i *int) *RecipeUpdateOne {
	if i != nil {
		ruo.SetMealType(*i)
	}
	return ruo
}

// AddMealType adds i to the "meal_type" field.
func (ruo *RecipeUpdateOne) AddMealType(i int) *RecipeUpdateOne {
	ruo.mutation.AddMealType(i)
	return ruo
}

// AddRecipeingredientIDs adds the "recipeingredient" edge to the RecipeIngredient entity by IDs.
func (ruo *RecipeUpdateOne) AddRecipeingredientIDs(ids ...int) *RecipeUpdateOne {
	ruo.mutation.AddRecipeingredientIDs(ids...)
	return ruo
}

// AddRecipeingredient adds the "recipeingredient" edges to the RecipeIngredient entity.
func (ruo *RecipeUpdateOne) AddRecipeingredient(r ...*RecipeIngredient) *RecipeUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddRecipeingredientIDs(ids...)
}

// Mutation returns the RecipeMutation object of the builder.
func (ruo *RecipeUpdateOne) Mutation() *RecipeMutation {
	return ruo.mutation
}

// ClearRecipeingredient clears all "recipeingredient" edges to the RecipeIngredient entity.
func (ruo *RecipeUpdateOne) ClearRecipeingredient() *RecipeUpdateOne {
	ruo.mutation.ClearRecipeingredient()
	return ruo
}

// RemoveRecipeingredientIDs removes the "recipeingredient" edge to RecipeIngredient entities by IDs.
func (ruo *RecipeUpdateOne) RemoveRecipeingredientIDs(ids ...int) *RecipeUpdateOne {
	ruo.mutation.RemoveRecipeingredientIDs(ids...)
	return ruo
}

// RemoveRecipeingredient removes "recipeingredient" edges to RecipeIngredient entities.
func (ruo *RecipeUpdateOne) RemoveRecipeingredient(r ...*RecipeIngredient) *RecipeUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveRecipeingredientIDs(ids...)
}

// Where appends a list predicates to the RecipeUpdate builder.
func (ruo *RecipeUpdateOne) Where(ps ...predicate.Recipe) *RecipeUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RecipeUpdateOne) Select(field string, fields ...string) *RecipeUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Recipe entity.
func (ruo *RecipeUpdateOne) Save(ctx context.Context) (*Recipe, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RecipeUpdateOne) SaveX(ctx context.Context) *Recipe {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RecipeUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RecipeUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RecipeUpdateOne) sqlSave(ctx context.Context) (_node *Recipe, err error) {
	_spec := sqlgraph.NewUpdateSpec(recipe.Table, recipe.Columns, sqlgraph.NewFieldSpec(recipe.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Recipe.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, recipe.FieldID)
		for _, f := range fields {
			if !recipe.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != recipe.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(recipe.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.SetField(recipe.FieldDescription, field.TypeString, value)
	}
	if value, ok := ruo.mutation.MealType(); ok {
		_spec.SetField(recipe.FieldMealType, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedMealType(); ok {
		_spec.AddField(recipe.FieldMealType, field.TypeInt, value)
	}
	if ruo.mutation.RecipeingredientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.RecipeingredientTable,
			Columns: []string{recipe.RecipeingredientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipeingredient.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedRecipeingredientIDs(); len(nodes) > 0 && !ruo.mutation.RecipeingredientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.RecipeingredientTable,
			Columns: []string{recipe.RecipeingredientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipeingredient.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RecipeingredientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   recipe.RecipeingredientTable,
			Columns: []string{recipe.RecipeingredientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recipeingredient.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Recipe{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recipe.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}