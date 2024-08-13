package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
    {"Name": "Cleo", "Wins": 10},
    {"Name": "Chris", "Wins": 33}]`)

		store := FileSystemPlayerStore{database}
		got := store.GetLeague()
		want := []Player{{"Cleo", 10}, {"Chris", 33}}
		assertLeague(t, got, want)

		/*
		   - read again
		   - we want this to pass, but if you run the test it doesn't
		   - the problem is our `Reader` has reached the end so there is nothing more to read.
		   - We need a way to tell it to go back to the start
		*/
		got = store.GetLeague()
		assertLeague(t, got, want)
	})
}
