package domain

import (
	"github.com/google/uuid"
)

type Repository interface {
	Save(Game) error
	Get(uuid.UUID) (Game, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Get(id uuid.UUID) (Game, error) {
	return s.repo.Get(id)
}

func (s *service) Save(g Game) error {
	return s.repo.Save(g)
}

func (s *service) Validate(prev Game, next Game) bool {
	changes := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// нельзя менять уже занятые клетки
			if prev.Board[i][j] != next.Board[i][j] {
				// новая клетка должна быть X (игрок)
				if prev.Board[i][j] != Empty || next.Board[i][j] != X {
					return false
				}
				changes++
			}
		}
	}
	// игрок может сделать ровно один ход
	return changes == 1
}

func (s *service) CheckWinner(b Board) int {
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
		a := b[line[0][0]][line[0][1]]
		if a != Empty && a == b[line[1][0]][line[1][1]] && a == b[line[2][0]][line[2][1]] {
			return a
		}
	}
	return Empty
}

// simple minimax omitted depth optimization
func (s *service) NextMove(g Game) (Game, error) {
	if s.CheckWinner(g.Board) != Empty {
		return g, nil
	}

	bestScore := -1000
	var move [2]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.Board[i][j] == Empty {
				g.Board[i][j] = O
				score := minimax(g.Board, false)
				g.Board[i][j] = Empty
				if score > bestScore {
					bestScore = score
					move = [2]int{i, j}
				}
			}
		}
	}

	g.Board[move[0]][move[1]] = O
	return g, nil
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
