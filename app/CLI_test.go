package poker_test

import (
	"bytes"
	"github.com/achristie/go-with-tests/app"
	"io"
	"strings"
	"testing"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type GameSpy struct {
	StartCalledWith  int
	FinishCalledWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalledWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishCalledWith = winner
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func TestCLI(t *testing.T) {
	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("7")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		got := stdout.String()
		want := "Please enter the number of players: "

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if game.StartCalledWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartCalledWith)
		}
	})
	t.Run("finish game with 'Chris' as winner", func(t *testing.T) {
		in := strings.NewReader("1\nChris wins\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		if game.FinishCalledWith != "Chris" {
			t.Errorf("expected finish called with 'Chris' but got %q", game.FinishCalledWith)
		}
	})
	t.Run("finish game with 'Kelsey' as winner", func(t *testing.T) {
		in := strings.NewReader("1\nKelsey wins\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		if game.FinishCalledWith != "Kelsey" {
			t.Errorf("expected finish called with 'Kelsey' but got %q", game.FinishCalledWith)
		}
	})
	t.Run("it prints an error when a non numeric value is entered and does not start", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		if game.StartCalledWith > 0 {
			t.Errorf("game should not have started")
		}
	})
}

func assertScheduledAlert(t testing.TB, got, want poker.ScheduledAlert) {
	t.Helper()
	if got.Amount != want.Amount {
		t.Errorf("got amount %d, want %d", got.Amount, want.Amount)
	}

	if got.At != want.At {
		t.Errorf("got scheduled time of %v, want %v", got.At, want.At)
	}
}
