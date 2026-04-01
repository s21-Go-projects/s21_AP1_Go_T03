package datasource

import (
	//"tictac/internal/domain"

	"tictac/internal/domain"

	"github.com/google/uuid"
)

type Repository interface {
	Save(domain.Game) error
	Get(uuid.UUID) (domain.Game, error)
}
