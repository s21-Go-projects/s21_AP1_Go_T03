package domain

type GameService interface {
	NextMove(game Game) Game
	Validate(game Game) error
	IsFinished(game Game) bool
}
