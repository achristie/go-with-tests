package main

import (
	"github.com/achristie/go-with-tests/app"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	db, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	server := poker.NewPlayerServer(db)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
