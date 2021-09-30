package main

import (
	go_http_application_with_tdd "github.com/crazyhl/go-http-application-with-tdd"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &go_http_application_with_tdd.PlayerServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}
