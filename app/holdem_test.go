package poker_test

import (
	"fmt"
	"github.com/achristie/go-with-tests/app"
	"testing"
	"time"
)

func TestGame_Start(t *testing.T) {
	t.Run("it schedules printing of blind values", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}

		game := poker.NewHoldem(blindAlerter, dummyPlayerStore)

		game.Start(5)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Minute, Amount: 100},
			{At: 10 * time.Minute, Amount: 200},
			{At: 20 * time.Minute, Amount: 300},
			{At: 30 * time.Minute, Amount: 400},
			{At: 40 * time.Minute, Amount: 500},
			{At: 50 * time.Minute, Amount: 750},
			{At: 60 * time.Minute, Amount: 1000},
			{At: 70 * time.Minute, Amount: 2000},
			{At: 80 * time.Minute, Amount: 4000},
			{At: 90 * time.Minute, Amount: 8000},
		}

		checkSchedulingCases(cases, t, blindAlerter)

	})
	t.Run("scheduled alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewHoldem(blindAlerter, dummyPlayerStore)

		game.Start(7)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Minute, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})
}

func TestGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := poker.NewHoldem(dummyBlindAlerter, store)
	winner := "Kelsey"

	game.Finish(winner)
	poker.AssertPlayerWin(t, store, winner)
}

func checkSchedulingCases(cases []poker.ScheduledAlert, t *testing.T, blindAlerter *poker.SpyBlindAlerter) {

	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {

			if len(blindAlerter.Alerts) <= 1 {
				t.Fatalf("alert %d was not schedule %v", i, blindAlerter.Alerts)
			}

			got := blindAlerter.Alerts[i]
			assertScheduledAlert(t, got, want)

		})
	}
}
