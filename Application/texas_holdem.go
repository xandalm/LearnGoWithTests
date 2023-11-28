package poker

import (
	"io"
	"time"
)

// TexasHoldem manages a game of poker.
type TexasHoldem struct {
	store   PlayerStore
	alerter BlindAlerter
}

// NewTexasHoldem returns a new game.
func NewTexasHoldem(alerter BlindAlerter, store PlayerStore) *TexasHoldem {
	return &TexasHoldem{
		alerter: alerter,
		store:   store,
	}
}

// Start will schedule blind alerts dependant on the number of players.
func (g *TexasHoldem) Start(numberOfPlayers int, alertsDestination io.Writer) {

	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		g.alerter.ScheduleAlertAt(blindTime, blind, alertsDestination)
		blindTime = blindTime + blindIncrement
	}
}

// Finish ends the game, recording the winner.
func (g *TexasHoldem) Finish(winner string) {
	g.store.RecordWin(winner)
}
