package domain

import (
	"errors"
)

type processor struct {
	service Service
}

func NewProcessor(s Service) Processor {
	return &processor{service: s}
}

func (p *processor) Process(g Game) (Game, error) {

	if !p.service.Validate(g) {
		return g, errors.New("invalid move")
	}

	if p.service.CheckWinner(g) != Empty {
		return g, nil
	}

	return p.service.NextMove(g), nil

}
