package domain

import "github.com/google/uuid"

type Repository interface {
	Save(Game) error
	Get(uuid.UUID) (Game, error)
}
