package main

import (
	"fmt"
	"github.com/achristie/go-with-tests/app"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	game := poker.NewHoldem(poker.BlindAlerterFunc(poker.StdOutAlerter), db)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)

	fmt.Println("let's play poker!")
	fmt.Println("Type {Name} wins to record a win")
	cli.PlayPoker()
}
