package main

import (
	gohttpapplicationwithtdd "github.com/crazyhl/go-http-application-with-tdd"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := gohttpapplicationwithtdd.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()
	server := gohttpapplicationwithtdd.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
