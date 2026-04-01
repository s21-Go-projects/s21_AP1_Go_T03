package domain

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Validate(game Game) (Game, bool) {
	prev, err := s.repo.Get(game.ID)
	if err != nil {
		s.repo.Save(game)
		return game, true
	}
	changes := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// нельзя менять уже занятые клетки
			if prev.Board[i][j] != game.Board[i][j] {
				// новая клетка должна быть X (игрок)
				if prev.Board[i][j] != Empty || game.Board[i][j] != X {
					return prev, false
				}
				changes++
			}
		}
	}
	if changes != 1 {
		return prev, false
	}
	return game, true
}

func (s *service) CheckWinner(game Game) int {
	lines := [][][2]int{
		{{0, 0}, {0, 1}, {0, 2}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {1, 0}, {2, 0}},
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},
		{{0, 0}, {1, 1}, {2, 2}},
		{{0, 2}, {1, 1}, {2, 0}},
	}

	for _, line := range lines {
		a := game.Board[line[0][0]][line[0][1]]
		if a != Empty && a == game.Board[line[1][0]][line[1][1]] && a == game.Board[line[2][0]][line[2][1]] {
			game.Winner = a
			s.repo.Save(game)
			return a
		}
	}
	return Empty
}

func (s *service) NextMove(game Game) Game {
	s.repo.Save(game)
	bestScore := -1000
	var move [2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if game.Board[i][j] == Empty {
				game.Board[i][j] = O
				score := minimax(game.Board, false)
				game.Board[i][j] = Empty
				if score > bestScore {
					bestScore = score
					move = [2]int{i, j}
				}
			}
		}
	}
	game.Board[move[0]][move[1]] = O
	s.repo.Save(game)
	a := s.CheckWinner(game)
	if a != Empty {
		game.Winner = a
		return game
	}

	return game
}

func minimax(board Board, isMax bool) int {
	winner := check(board)
	if winner == O {
		return 1
	}
	if winner == X {
		return -1
	}

	full := true
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == Empty {
				full = false
			}
		}
	}
	if full {
		return 0
	}

	if isMax {
		best := -1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == Empty {
					board[i][j] = O
					val := minimax(board, false)
					board[i][j] = Empty
					if val > best {
						best = val
					}
				}
			}
		}
		return best
	}

	best := 1000
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == Empty {
				board[i][j] = X
				val := minimax(board, true)
				board[i][j] = Empty
				if val < best {
					best = val
				}
			}
		}
	}
	return best
}

func check(b Board) int {
	lines := [][][2]int{
		{{0, 0}, {0, 1}, {0, 2}}, {{1, 0}, {1, 1}, {1, 2}}, {{2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {1, 0}, {2, 0}}, {{0, 1}, {1, 1}, {2, 1}}, {{0, 2}, {1, 2}, {2, 2}},
		{{0, 0}, {1, 1}, {2, 2}}, {{0, 2}, {1, 1}, {2, 0}},
	}
	for _, l := range lines {
		a := b[l[0][0]][l[0][1]]
		if a != Empty && a == b[l[1][0]][l[1][1]] && a == b[l[2][0]][l[2][1]] {
			return a
		}
	}
	return Empty
}
