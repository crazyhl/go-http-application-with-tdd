package go_http_application_with_tdd

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	league, _ := NewLeague(f.database)
	return league
}
