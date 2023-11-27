package poker_test

import (
	"bytes"
	"poker"
	"strings"
	"testing"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type SpyGame struct {
	StartCalled  bool
	StartedWith  int
	FinishedWith string
}

func (s *SpyGame) Start(numberOfPlayers int) {
	s.StartCalled = true
	s.StartedWith = numberOfPlayers
}

func (s *SpyGame) Finish(winner string) {
	s.FinishedWith = winner
}

func TestCLI(t *testing.T) {
	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")

		game := &SpyGame{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 7)
	})
	t.Run("start game with 5 players and finish with 'Chris' as winner", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins\n")

		game := &SpyGame{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertGameStartedWith(t, game, 5)
		assertGameFinishedWith(t, game, "Chris")
	})
	t.Run("start game with 9 players and finish with 'Cleo' as winner", func(t *testing.T) {
		in := strings.NewReader("9\nCleo wins\n")

		game := &SpyGame{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertGameStartedWith(t, game, 9)
		assertGameFinishedWith(t, game, "Cleo")
	})
	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &SpyGame{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})
}

func assertMessagesSentToUser(t testing.TB, stdout *bytes.Buffer, messages ...string) {
	t.Helper()

	got := stdout.String()
	want := strings.Join(messages, "")

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertGameStartedWith(t testing.TB, game *SpyGame, numberOfPlayers int) {
	t.Helper()

	if game.StartedWith != numberOfPlayers {
		t.Errorf("wanted Start called with %d but got %d", numberOfPlayers, game.StartedWith)
	}
}

func assertGameFinishedWith(t testing.TB, game *SpyGame, winner string) {
	t.Helper()

	if game.FinishedWith != winner {
		t.Errorf("expected finish called with %q but got %q", winner, game.FinishedWith)
	}
}

func assertGameNotStarted(t testing.TB, game *SpyGame) {
	t.Helper()

	if game.StartCalled {
		t.Errorf("game should not have called")
	}
}
