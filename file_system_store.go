package go_http_application_with_tdd

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	Database io.Writer
	League   League
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, 0)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		Database: database,
		League:   league,
	}
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.League
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.GetLeague().Find(name)

	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		f.League = append(f.League, Player{
			Name: name,
			Wins: 1,
		})
	}

	json.NewEncoder(f.Database).Encode(f.League)
}
