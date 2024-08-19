package poker_test

import (
	"strings"
	"testing"
	"time"

	poker "github.com/DarrelASandbox/playground-go/chris_james/03-build-an-app/time"
)

type SpyBlindAlerter struct {
	alerts []struct {
		scheduleAt time.Duration
		amount     int
	}
}

func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, struct {
		scheduleAt time.Duration
		amount     int
	}{duration, amount})
}

var dummySpyAlerter = &SpyBlindAlerter{}

func TestCLI(t *testing.T) {
	playerStore := &poker.StubPlayerStore{}
	blindAlerter := &SpyBlindAlerter{}

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		if len(dummySpyAlerter.alerts) != 1 {
			t.Fatal("expected a blind alert alert to be scheduled")
		}
	})
}
