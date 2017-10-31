package sudoku

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ErrUnsolvable represents a sudoku board that is unsolvable in the current state.
var ErrUnsolvable = errors.New("sudoku board has no solution")

// Board represents a 9x9 sudoku board.
type Board [9][9]int

// NewBoardFromReader creates a sudoku board from a reader.
func NewBoardFromReader(in io.Reader) (Board, error) {
	var b Board

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanWords)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if !s.Scan() {
				if err := s.Err(); err != nil {
					return b, err
				}

				// reached EOF
				return b, errors.New("invalid input: too short")
			}

			if s.Text() == "_" {
				continue
			}

			d, err := strconv.Atoi(s.Text())
			if err != nil {
				return b, fmt.Errorf("invalid input: %s", s.Text())
			}

			if d < 1 || d > 9 {
				return b, fmt.Errorf("invalid input: %d", d)
			}

			b[i][j] = d
		}
	}

	return b, nil
}

func (b Board) String() string {
	var buf bytes.Buffer

	for _, row := range b {
		for j, d := range row {
			if j == 0 {
				fmt.Fprintf(&buf, "%d", d)
				continue
			}
			fmt.Fprintf(&buf, "%2d", d)
		}
		(&buf).WriteByte('\n')
	}

	return (&buf).String()
}

// IsValid checks if the value satisfies row, column, and grid constraints.
func (b Board) IsValid(d, m, n int) bool {
	// bounds constraint
	if d < 1 || d > 9 {
		return false
	}

	// row constraint
	for _, a := range b[m] {
		if a == d {
			return false
		}
	}

	// column constraint
	for _, row := range b {
		if row[n] == d {
			return false
		}
	}

	// grid constraint
	for _, row := range b[(m/3)*3 : (m/3+1)*3] {
		for _, a := range row[(n/3)*3 : (n/3+1)*3] {
			if a == d {
				return false
			}
		}
	}

	return true
}

// Solver is the interface that wraps the basic Solve method.
//
// If Solve is given a sudoku board that cannot be solved, ErrUnsolvable
// should be returned.
type Solver interface {
	Solve(b Board) (Board, error)
}
