package poker_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	poker "github.com/DarrelASandbox/playground-go/chris_james/03-build-an-app/time"
)

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{at, amount})
}

var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdOut = &bytes.Buffer{}

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nChris wins\n")
		playerStore := &poker.StubPlayerStore{}
		game := poker.NewTexasHoldem(dummyBlindAlerter, playerStore)
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("1\nCleo wins\n")
		playerStore := &poker.StubPlayerStore{}
		game := poker.NewTexasHoldem(dummyBlindAlerter, playerStore)
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

	// t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
	// 	stdout := &bytes.Buffer{}
	// 	in := strings.NewReader("7\n")
	// 	blindAlerter := &poker.SpyBlindAlerter{}
	// 	game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)
	// 	cli := poker.NewCLI(in, stdout, game)
	// 	cli.PlayPoker()

	// 	got := stdout.String()
	// 	want := poker.PlayerPrompt
	// 	if got != want {
	// 		t.Errorf("got %q, want %q", got, want)
	// 	}

	// 	cases := []poker.ScheduledAlert{
	// 		{0 * time.Second, 100},
	// 		{12 * time.Minute, 200},
	// 		{24 * time.Minute, 300},
	// 		{36 * time.Minute, 400},
	// 	}

	// 	for i, want := range cases {
	// 		t.Run(fmt.Sprint(want), func(t *testing.T) {
	// 			if len(blindAlerter.alerts) <= i {
	// 				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
	// 			}

	// 			got := blindAlerter.alerts[i]
	// 			assertScheduledAlert(t, got, want)
	// 		})
	// 	}
	// })
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
