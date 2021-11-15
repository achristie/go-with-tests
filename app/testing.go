package poker

import (
	"testing"
	"reflect"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}


func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != winner {
			t.Errorf("did not store the correct winner, got %q want %q", store.winCalls[0], winner)
}
}

func AssertLeague(t testing.TB, got, wanted League) {
	t.Helper()

	if !reflect.DeepEqual(got, wanted) {
		t.Errorf("got %v want %v", got, wanted)
	}
}
func AssertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)

	}
}

func AssertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
func AssertScoreEquals(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
			t.Errorf("got %d want %d", got, want)
	}
}
func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
