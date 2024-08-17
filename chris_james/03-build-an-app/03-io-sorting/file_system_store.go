package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

/*
Every time someone calls GetLeague() or GetPlayerScore() we are reading the entire file and parsing it into JSON.
We should not have to do that because FileSystemStore is entirely responsible for the state of the league;
it should only need to read the file when the program starts up and only need to update the file when data changes.

We can create a constructor which can do some of this initialization for us and store the league as a value in our FileSystemStore to be used on the reads instead.
*/
type FileSystemPlayerStore struct {
	database *json.Encoder
	league   League
}

func NewFileSystemPlayerStore(file *os.File) (*FileSystemPlayerStore, error) {
	err := initializePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("problem initializing player db file %v", err)
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemPlayerStore{database: json.NewEncoder(&tape{file}), league: league}, nil
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

	/*
		In RecordWin we have the line json.NewEncoder(f.database).Encode(f.league).
		We don't need to create a new encoder every time we write, we can initialize one in our constructor and use that instead.
	*/
	f.database.Encode(f.league)
}

func initializePlayerDBFile(file *os.File) error {
	file.Seek(0, io.SeekStart)
	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, io.SeekStart)
	}

	return nil
}
