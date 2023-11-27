package poker

import (
	"fmt"
	"testing"
	"time"
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

// ScheduledAlert holds information about when an alert is scheduled.
type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

// SpyBlindAlerter allows you to spy on ScheduleAlertAt calls.
type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

// ScheduleAlertAt records alerts that have been scheduled.
func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduledAlert{duration, amount})
}
