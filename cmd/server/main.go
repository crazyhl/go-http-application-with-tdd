package main

import (
	go_http_application_with_tdd "github.com/crazyhl/go-http-application-with-tdd"
	"log"
	"net/http"
)

func main() {
	store := go_http_application_with_tdd.InMemoryPlayerStore{}
	server := &go_http_application_with_tdd.PlayerServer{Store: &store}
	log.Fatal(http.ListenAndServe(":5000", server))
}
