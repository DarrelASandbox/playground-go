package main

import (
	"log"
	"net/http"

	specs_greet "github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet"
)

func main() {
	handler := http.HandlerFunc(specs_greet.Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
