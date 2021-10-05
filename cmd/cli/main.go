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

	store, close, err := go_http_application_with_tdd.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	game := go_http_application_with_tdd.NewCLI(store, os.Stdin, go_http_application_with_tdd.BlindAlerterFunc(go_http_application_with_tdd.StdOutAlerter))
	game.PlayPoker()
}
