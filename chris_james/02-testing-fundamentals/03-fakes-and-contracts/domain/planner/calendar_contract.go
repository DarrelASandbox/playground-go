package planner

import (
	"context"
	"testing"
	"time"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/planner/internal/expect"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/planner/internal/plannertest"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/recipe"
)

type Calendar interface {
	ScheduleMeal(ctx context.Context, recipe recipe.Recipe, date time.Time) error
	GetSchedule(ctx context.Context) (map[time.Time]recipe.Recipes, error)
}

type CalendarContract struct {
	NewCalendar func() Calendar
}

func (c CalendarContract) Test(t *testing.T) {
	t.Run("it returns what is put in", func(t *testing.T) {
		var (
			ctx         = context.Background()
			someRecipes = plannertest.RandomRecipes()
			tomorrow    = time.Now()
			sut         = c.NewCalendar()
		)

		for _, r := range someRecipes {
			expect.NoErr(t, sut.ScheduleMeal(ctx, r, tomorrow))
		}

		got, err := sut.GetSchedule(ctx)
		expect.NoErr(t, err)
		expect.DeepEqual(t, got[tomorrow], someRecipes)
	})
}
