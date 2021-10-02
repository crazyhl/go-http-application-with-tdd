package go_http_application_with_tdd

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	Database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() League {
	f.Database.Seek(0, 0)
	league, _ := NewLeague(f.Database)
	return league
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
		league = append(league, Player{
			Name: name,
			Wins: 1,
		})
	}

	f.Database.Seek(0, 0)
	json.NewEncoder(f.Database).Encode(league)
}
