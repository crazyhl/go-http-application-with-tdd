package main

import (
	gohttpapplicationwithtdd "github.com/crazyhl/go-http-application-with-tdd"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}
	store, _ := gohttpapplicationwithtdd.NewFileSystemPlayerStore(db)
	server := gohttpapplicationwithtdd.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
