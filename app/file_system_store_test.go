package poker

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {

	// 	t.Run("league from a reader", func(t *testing.T) {
	// 		database, cleanDatabase := createTempFile(t, `[
	// {"Name": "Cleo", "Wins":10},
	// {"Name":"Chris", "Wins": 33}
	// 		]`)

	// 		defer cleanDatabase()
	// 		store, err := NewFileSystemPlayerStore( database )
	// 		assertNoError(t, err)

	// 		got := store.GetLeague()
	// 		want := League{
	// 			{"Cleo", 10},
	// 			{"Chris", 33},
	// 		}
	// 		got = store.GetLeague()

	// 		assertLeague(t, got, want)
	// 	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
{"Name": "Cleo", "Wins":10},
{"Name":"Chris", "Wins": 33}
		]`)

		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 33

		AssertScoreEquals(t, got, want)
	})
	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
{"Name": "Cleo", "Wins":10},
{"Name":"Chris", "Wins": 33}
		]`)

		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34

		AssertScoreEquals(t, got, want)
	})
	t.Run("store wins for a new player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
{"Name": "Cleo", "Wins":10},
{"Name":"Chris", "Wins": 33}
		]`)

		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)

		store.RecordWin("Andrew")

		got := store.GetPlayerScore("Andrew")
		want := 1

		AssertScoreEquals(t, got, want)
	})
	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)
	})
	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
{"Name": "Cleo", "Wins":10},
{"Name":"Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, err := NewFileSystemPlayerStore(database)

		AssertNoError(t, err)

		got := store.GetLeague()

		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		AssertLeague(t, got, want)

		got = store.GetLeague()
		AssertLeague(t, got, want)
	})
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()
	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
