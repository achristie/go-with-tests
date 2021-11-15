package poker_test

import (
	"testing"
	"github.com/achristie/go-with-tests/app"
	"fmt"
	"strings"
	"time"
)

var dummySpyAlerter = &SpyBlindAlerter{}

type scheduledAlert struct {
	at time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{at, amount})
}

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})
	t.Run("record kelsey win from user input", func(t *testing.T) {
		in := strings.NewReader("Kelsey wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Kelsey")
	})
	t.Run("it schedules printing of blind values", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &SpyBlindAlerter{}

		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()

		cases := []scheduledAlert{
			{ 0 * time.Minute, 100 },
			{ 10 * time.Minute, 200 },
			{ 20 * time.Minute, 300},
			{ 30 * time.Minute, 400 },
			{ 40 * time.Minute, 500 },
			{ 50 * time.Minute, 750 },
			{ 60 * time.Minute, 1000 },
			{ 70 * time.Minute, 2000 },
			{ 80 * time.Minute, 4000 },
			{ 90 * time.Minute, 8000 },
		}

		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t * testing.T) {

				if len(blindAlerter.alerts) <= 1 {
					t.Fatalf("alert %d was not schedule %v", i, blindAlerter.alerts)
				}

				got := blindAlerter.alerts[i]
				assertScheduledAlert(t, got, want)

			})
		}
	})
}
func assertScheduledAlert(t testing.TB, got, want scheduledAlert) {
	t.Helper()
	if got.amount != want.amount {
		t.Errorf("got amount %d, want %d", got.amount, want.amount)
	}

	if got.at != want.at {
		t.Errorf("got scheduled time of %v, want %v", got.at, want.at)
	}
}
