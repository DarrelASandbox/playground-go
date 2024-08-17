package main

import (
	"log"
	"net/http"
	"os"

	poker "github.com/DarrelASandbox/playground-go/chris_james/03-build-an-app/cli-pkg-struct"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	/*
		We cannot parse the league because the file is empty.
		We weren't getting errors before because we always just ignored them.
	*/
	store, err := poker.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	server := poker.NewPlayerServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
