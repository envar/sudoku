package backtracker

import (
	"github.com/envar/sudoku"
)

// Solver represent the recursive backtracker sudoku solver.
type Solver struct{}

// NewSolver returns a solver.
func NewSolver() Solver {
	return Solver{}
}

// Solve implements the sudoku.Solver interface.
func (s Solver) Solve(b sudoku.Board) (sudoku.Board, error) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			// skip filled positions
			if b[i][j] != 0 {
				continue
			}

			// test 1..9 and solving
			for d := 1; d <= 9; d++ {
				if b.IsValid(d, i, j) {
					b[i][j] = d
					solved, err := s.Solve(b)
					if err == nil {
						return solved, nil
					}
				}
			}

			return b, sudoku.ErrUnsolvable
		}
	}

	return b, nil
}
