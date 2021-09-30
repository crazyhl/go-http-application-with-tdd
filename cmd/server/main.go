package main

import (
	gohttpapplicationwithtdd "github.com/crazyhl/go-http-application-with-tdd"
	"log"
	"net/http"
)

func main() {
	server := &gohttpapplicationwithtdd.PlayerServer{Store: gohttpapplicationwithtdd.NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
