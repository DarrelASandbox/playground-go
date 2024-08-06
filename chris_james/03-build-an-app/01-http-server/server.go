package main

import (
	"fmt"
	"net/http"
)

// PlayerServer currently returns "20" given _any_ requests.
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
