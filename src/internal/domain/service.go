package domain

type Service interface {
	NextMove(game Game) Game
	Validate(game Game) (Game, bool)
	CheckWinner(game Game) int
}
