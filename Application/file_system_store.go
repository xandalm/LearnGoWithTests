package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker // because we need to move to start of the file content
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}
