package datasource

import (
	"errors"
	"tictac/internal/domain"

	"github.com/google/uuid"
)

type repo struct {
	storage *Storage
}

func NewRepository(s *Storage) Repository {
	return &repo{storage: s}
}

func (r *repo) Save(g domain.Game) error {
	r.storage.Games.Store(g.ID.String(), g)
	return nil
}

func (r *repo) Get(id uuid.UUID) (domain.Game, error) {
	v, ok := r.storage.Games.Load(id.String())
	if !ok {
		return domain.Game{}, errors.New("not found")
	}
	return v.(domain.Game), nil
}
