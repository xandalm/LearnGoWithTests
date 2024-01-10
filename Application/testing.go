package poker

import (
	"fmt"
	"io"
	"testing"
	"time"
)

// StubPlayerStore implements PlayerStore for testing purposes.
type StubPlayerStore struct {
	Scores   map[string]int
	WinCalls []string
	League   League
}

// GetPlayerScore returns a score from Scores.
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.Scores[name]
	return score
}

// RecordWin will record a win to WinCalls.
func (s *StubPlayerStore) RecordWin(name string) {
	s.WinCalls = append(s.WinCalls, name)
}

// GetLeague returns League.
func (s *StubPlayerStore) GetLeague() League {
	return s.League
}

// AssertPlayerWin allows you to spy on the store's calls to RecordWin.
func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.WinCalls) != 1 {
		t.Fatalf("expected a win call but didn't get any")
	}

	if store.WinCalls[0] != winner {
		t.Errorf("didn't record correct winner, got %q but want %q", store.WinCalls[0], winner)
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
func (s *SpyBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int, to io.Writer) {
	s.Alerts = append(s.Alerts, ScheduledAlert{duration, amount})
}

type SpyGame struct {
	StartCalled  bool
	StartedWith  int
	FinishedWith string
}

func (s *SpyGame) Start(numberOfPlayers int, alertsDestination io.Writer) {
	s.StartCalled = true
	s.StartedWith = numberOfPlayers
}

func (s *SpyGame) Finish(winner string) {
	s.FinishedWith = winner
}
