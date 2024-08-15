package main

import (
	"encoding/json"
	"io"
)

/*
Every time someone calls GetLeague() or GetPlayerScore() we are reading the entire file and parsing it into JSON.
We should not have to do that because FileSystemStore is entirely responsible for the state of the league;
it should only need to read the file when the program starts up and only need to update the file when data changes.

We can create a constructor which can do some of this initialization for us and store the league as a value in our FileSystemStore to be used on the reads instead.
*/
type FileSystemPlayerStore struct {
	database io.Writer
	league   League
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, io.SeekStart)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{database: &tape{database}, league: league}
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}

	json.NewEncoder(f.database).Encode(f.league)
}
