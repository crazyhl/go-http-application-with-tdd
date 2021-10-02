package go_http_application_with_tdd

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	store map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	var players []Player
	for name, winners := range i.store {
		players = append(players, Player{
			Name: name,
			Wins: winners,
		})
	}

	return players
}
