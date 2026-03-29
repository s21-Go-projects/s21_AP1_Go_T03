package domain

type Cell int

const (
	Empty Cell = 0
	X     Cell = 1
	O     Cell = 2
)

type Board [][]Cell

type Game struct {
	ID    string
	Board Board
}