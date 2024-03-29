package main

import (
	"fmt"
	"log"
	"os"
	"poker"
)

const dbFileName = "game.db.json"

func main() {

	store, closeFunc, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	fmt.Println("Let's play poker")
	fmt.Println("Type `{Name} wins` to record a win")
	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}
