package main

import (
	gohttpapplicationwithtdd "github.com/crazyhl/go-http-application-with-tdd"
	"log"
	"net/http"
)

func main() {
	store := gohttpapplicationwithtdd.NewInMemoryPlayerStore()
	server := gohttpapplicationwithtdd.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
