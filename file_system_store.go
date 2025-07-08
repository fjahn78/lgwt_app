package main

import (
	"io"
	"encoding/json"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()

	player := league.Find(name)	
	if player != nil {
		player.Wins++
	}

	f.database.Seek(0, io.SeekStart)
	json.NewEncoder(f.database).Encode(league)
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int

	player := f.GetLeague().Find(name)	
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, io.SeekStart)
	league, _ := NewLeague(f.database)
	return league
}
