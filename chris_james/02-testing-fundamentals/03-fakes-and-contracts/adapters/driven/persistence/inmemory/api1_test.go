package inmemory_test

import (
	"testing"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/adapters/driven/persistence/inmemory"
	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/fakes-and-contracts/domain/planner"
)

func TestInMemoryAPI1(t *testing.T) {
	planner.API1Contract{NewAPI1: func() planner.API1 { return inmemory.NewAPI1() }}.Test(t)
}
