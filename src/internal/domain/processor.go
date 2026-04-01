package domain

type Processor interface {
	Process(game Game) (Game, error)
}
