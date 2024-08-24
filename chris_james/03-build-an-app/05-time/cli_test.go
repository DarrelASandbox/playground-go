package poker_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	poker "github.com/DarrelASandbox/playground-go/chris_james/03-build-an-app/time"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdOut = &bytes.Buffer{}

type GameSpy struct {
	StartCalled      bool
	StartCalledWith  int
	FinishedCalled   bool
	FinishCalledWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalledWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishCalledWith = winner
}

type failOnEndReader struct {
	t   *testing.T
	rdr io.Reader
}

func (m failOnEndReader) Read(p []byte) (n int, err error) {
	n, err = m.rdr.Read(p)
	if n == 0 || err == io.EOF {
		m.t.Fatal("Read to the end when you shouldn't have")
	}
	return n, err
}

func TestCLI(t *testing.T) {
	t.Run("start game with 3 players and finish game with 'Chris as winner", func(t *testing.T) {
		in := userSends("3", "Chris wins")
		stdout := &bytes.Buffer{}
		game := &GameSpy{}
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and record 'Cleo' as winner", func(t *testing.T) {
		in := userSends("8", "Cleo wins")
		game := &GameSpy{}
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()
		assertGameStartedWith(t, game, 8)
		assertFinishCalledWith(t, game, "Cleo")
	})

	t.Run("do no read beyond the first newline", func(t *testing.T) {
		in := failOnEndReader{t, strings.NewReader("1\nChris wins\n hello there")}
		game := poker.NewTexasHoldem(dummyBlindAlerter, dummyPlayerStore)
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		in := strings.NewReader("Pies\n")
		stdout := &bytes.Buffer{}
		game := &GameSpy{}
		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()
		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func assertGameStartedWith(t testing.TB, game *GameSpy, numberOfPlayersWanted int) {
	t.Helper()
	if game.StartCalledWith != numberOfPlayersWanted {
		t.Errorf("wanted Start called with %d but got %d", numberOfPlayersWanted, game.StartCalledWith)
	}
}

func assertGameNotStarted(t testing.TB, game *GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func assertFinishCalledWith(t testing.TB, game *GameSpy, winner string) {
	t.Helper()
	if game.FinishCalledWith != winner {
		t.Errorf("expected finish called with %q but got %q", winner, game.FinishCalledWith)
	}
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}
