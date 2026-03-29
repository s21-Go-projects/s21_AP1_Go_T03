package datasource

import "sync"

type Storage struct {
	data sync.Map
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Save(game GameDTO) {
	s.data.Store(game.ID, game)
}

func (s *Storage) Get(id string) (GameDTO, bool) {
	val, ok := s.data.Load(id)
	if !ok {
		return GameDTO{}, false
	}
	return val.(GameDTO), true
}
