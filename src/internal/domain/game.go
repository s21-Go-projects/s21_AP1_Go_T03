package domain

import "github.com/google/uuid"

type Board [3][3]int

const (
	Empty = 0
	X     = 1
	O     = 2
)

type Game struct {
	ID     uuid.UUID
	Board  Board
	Winner int
}
