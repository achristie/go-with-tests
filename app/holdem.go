package poker

import (
	"io"
	"os"
	"time"
)

type Holdem struct {
	alerter BlindAlerter
	store   PlayerStore
}

func NewHoldem(alerter BlindAlerter, store PlayerStore) *Holdem {
	return &Holdem{
		alerter: alerter,
		store:   store,
	}
}

func (p *Holdem) Start(numberOfPlayers int, to io.Writer) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 750, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind, os.Stdout)
		blindTime = blindTime + blindIncrement
	}
}

func (p *Holdem) Finish(winner string) {
	p.store.RecordWin(winner)
}
