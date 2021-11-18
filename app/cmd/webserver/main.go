package main

import (
	"log"
	"net/http"

	poker "github.com/achristie/go-with-tests/app"
)

const dbFileName = "game.db.json"

func main() {
	db, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	game := poker.NewHoldem(poker.BlindAlerterFunc(poker.Alerter), db)
	server, err := poker.NewPlayerServer(db, game)

	if err != nil {
		log.Fatalf("problem creating player server %v", err)
	}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
