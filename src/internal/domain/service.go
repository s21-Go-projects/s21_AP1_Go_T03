package domain

type Service interface {
	NextMove(game Game) Game
	Validate(game Game) bool
	CheckWinner(game Game) int
}
