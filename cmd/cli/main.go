package main

import (
	"fmt"
	go_http_application_with_tdd "github.com/crazyhl/go-http-application-with-tdd"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := go_http_application_with_tdd.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	game := go_http_application_with_tdd.CLI{store, os.Stdin}
	game.PlayPoker()
}
