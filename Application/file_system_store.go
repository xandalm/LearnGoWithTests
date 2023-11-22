package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker // because we need to move to start of the file content (and we need to write into file now)
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	var wins int
	for _, player := range f.GetLeague() {
		if player.Name == name {
			wins = player.Wins
			break
		}
	}
	return wins
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	for i, player := range league {
		if player.Name == name {
			league[i].Wins++
			break
		}
	}
	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}
