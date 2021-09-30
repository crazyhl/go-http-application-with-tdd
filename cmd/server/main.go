package main

import (
	go_http_application_with_tdd "github.com/crazyhl/go-http-application-with-tdd"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(go_http_application_with_tdd.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
