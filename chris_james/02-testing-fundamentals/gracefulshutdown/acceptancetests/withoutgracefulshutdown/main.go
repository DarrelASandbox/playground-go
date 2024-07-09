package main

import (
	"log"
	"net/http"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/gracefulshutdown/acceptancetests"
)

func main() {
	server := &http.Server{Addr: ":8081", Handler: http.HandlerFunc(acceptancetests.SlowHandler)}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
