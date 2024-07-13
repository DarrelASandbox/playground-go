package main

import (
	"log"
	"net/http"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/adapters/httpserver"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
