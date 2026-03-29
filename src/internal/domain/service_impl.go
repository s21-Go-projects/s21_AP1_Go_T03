package domain

import "errors"

type gameService struct{}

func NewGameService() GameService {
	return &gameService{}
}

func (s *gameService) NextMove(game Game) Game {
	// простейший ход (не полный минимакс для краткости)
	for i := range game.Board {
		for j := range game.Board[i] {
			if game.Board[i][j] == Empty {
				game.Board[i][j] = O
				return game
			}
		}
	}
	return game
}

func (s *gameService) Validate(game Game) error {
	countX, countO := 0, 0

	for _, row := range game.Board {
		for _, c := range row {
			if c == X {
				countX++
			}
			if c == O {
				countO++
			}
		}
	}

	if countO > countX {
		return errors.New("invalid board state")
	}

	return nil
}

func (s *gameService) IsFinished(game Game) bool {
	// проверка строк
	for i := 0; i < 3; i++ {
		if game.Board[i][0] != Empty &&
			game.Board[i][0] == game.Board[i][1] &&
			game.Board[i][1] == game.Board[i][2] {
			return true
		}
	}

	return false
}
