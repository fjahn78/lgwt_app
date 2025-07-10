package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/fjahn78/lgwt_app"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker!")
	fmt.Println("Type '{Name} wins' to record a win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem oopening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store", dbFileName, err)
	}

	game := poker.CLI{store, os.Stdin}
	game.PlayPoker()
}
