package poker

import (
	"testing"
)

// StubPlayerStore implements PlayerStore for testing purposes.
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

// GetPlayerScore returns a score from Scores.
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

// RecordWin will record a win to WinCalls.
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

// GetLeague returns League.
func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

// AssertPlayerWin allows you to spy on the store's calls to RecordWin.
func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("expected a win call but didn't get any")
	}

	if store.winCalls[0] != winner {
		t.Errorf("didn't record correct winner, got %q but want %q", store.winCalls[0], winner)
	}
}

// todo - the rest of the helpers
