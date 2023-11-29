package poker_test

import (
	"bytes"
	"io"
	"poker"
	"strings"
	"testing"
	"time"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type SpyGame struct {
	StartCalled bool
	StartedWith int
	BlindAlert  []byte

	FinishedWith string
}

func (s *SpyGame) Start(numberOfPlayers int, alertsDestination io.Writer) {
	s.StartCalled = true
	s.StartedWith = numberOfPlayers
	alertsDestination.Write(s.BlindAlert)
}

func (s *SpyGame) Finish(winner string) {
	s.FinishedWith = winner
}

func userSends(messages ...string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func TestCLI(t *testing.T) {
	t.Run("start game with 5 players and finish with 'Chris' as winner", func(t *testing.T) {
		in := userSends("5", "Chris wins")

		game := &SpyGame{}

		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()

		assertGameStartedWith(t, game, 5)
		assertGameFinishedWith(t, game, "Chris")
	})
	t.Run("start game with 9 players and finish with 'Cleo' as winner", func(t *testing.T) {
		in := userSends("9", "Cleo wins")

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
	t.Run("it prints an error when the winner is declared incorrectly", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("1", "Lloyd is killer")
		game := &SpyGame{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotFinished(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputMsg)
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

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.StartedWith == numberOfPlayers
	})

	if !passed {
		t.Errorf("wanted Start called with %d but got %d", numberOfPlayers, game.StartedWith)
	}
}

func assertGameFinishedWith(t testing.TB, game *SpyGame, winner string) {
	t.Helper()

	passed := retryUntil(500*time.Millisecond, func() bool {
		return game.FinishedWith == winner
	})

	if !passed {
		t.Errorf("expected finish called with %q but got %q", winner, game.FinishedWith)
	}
}

func retryUntil(d time.Duration, f func() bool) bool {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		if f() {
			return true
		}
	}
	return false
}

func assertGameNotStarted(t testing.TB, game *SpyGame) {
	t.Helper()

	if game.StartCalled {
		t.Errorf("game should not have called")
	}
}

func assertGameNotFinished(t testing.TB, game *SpyGame) {
	t.Helper()

	if game.FinishedWith != "" {
		t.Error("game should not have finished")
	}
}
