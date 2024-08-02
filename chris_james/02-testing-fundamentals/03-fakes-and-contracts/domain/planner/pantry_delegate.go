package planner

import (
	"context"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/ingredients"
)

type PantryDelegate struct {
	delegate           Pantry
	GetIngredientsFunc func(ctx context.Context) (ingredients.Ingredients, error)
	StoreFunc          func(ctx context.Context, i ...ingredients.Ingredient) error
	RemoveFunc         func(ctx context.Context, i ...ingredients.Ingredient) error
}

func NewPantryDelegate(delegate Pantry) *PantryDelegate {
	return &PantryDelegate{delegate: delegate}
}

func (p *PantryDelegate) GetIngredients(ctx context.Context) (ingredients.Ingredients, error) {
	if p.GetIngredientsFunc != nil {
		return p.GetIngredientsFunc(ctx)
	}
	return p.delegate.GetIngredients(ctx)
}

func (p *PantryDelegate) Store(ctx context.Context, ingredients ...ingredients.Ingredient) error {
	if p.StoreFunc != nil {
		return p.StoreFunc(ctx, ingredients...)
	}
	return p.delegate.Store(ctx, ingredients...)
}

func (p *PantryDelegate) Remove(ctx context.Context, ingredients ...ingredients.Ingredient) error {
	if p.RemoveFunc != nil {
		return p.RemoveFunc(ctx, ingredients...)
	}
	return p.delegate.Remove(ctx, ingredients...)
}
