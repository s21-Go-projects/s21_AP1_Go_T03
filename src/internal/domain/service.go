package domain

import "github.com/google/uuid"

type Service interface {
	NextMove(game Game) (Game, error)
	Validate(prev Game, next Game) bool
	CheckWinner(board Board) int
	Get(id uuid.UUID) (Game, error)
	Save(game Game) error
}
