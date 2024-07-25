package httpserver

import (
	"fmt"
	"net/http"

	"github.com/DarrelASandbox/playground-go/chris_james/02-testing-fundamentals/specs-greet/domain/interactions"
)

const (
	greetPath = "/greet"
	cursePath = "/curse"
)

/*
HTTP handlers should only be responsible for handling HTTP concerns;
any "domain logic" should live outside of the handler.

This allows us to develop domain logic in isolation from HTTP, making it simpler to test and understand.
*/
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(greetPath, replyWith(interactions.Greet))
	mux.HandleFunc(cursePath, replyWith(interactions.Curse))
	return mux
}

func replyWith(f func(name string) (interactions string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprint(w, f(name))
	}
}
