package sudoku_test

import (
	"os"
	"testing"

	"github.com/envar/sudoku"
	"github.com/envar/sudoku/backtracker"
)

func testLoadBoard(tb testing.TB, board string) sudoku.Board {
	f, err := os.Open("./testdata/" + board + ".txt")
	if err != nil {
		tb.Fatal(err)
	}
	defer f.Close()

	b, err := sudoku.NewBoard(f)
	if err != nil {
		tb.Fatal(err)
	}

	return b
}

func TestNewBoard(t *testing.T) {
	b := testLoadBoard(t, "sample")

	expected := `1 0 3 0 0 6 0 8 0
0 5 0 0 8 0 1 2 0
7 0 9 1 0 3 0 5 6
0 3 0 0 6 7 0 9 0
5 0 7 8 0 0 0 3 0
8 0 1 0 3 0 5 0 7
0 4 0 0 7 8 0 1 0
6 0 8 0 0 2 0 4 0
0 1 2 0 4 5 0 7 8
`

	if b.String() != expected {
		t.Fatalf("expected\n%s\n, got\n%s\n.", expected, b)
	}
}

func TestBacktrackerSolver(t *testing.T) {
	b := testLoadBoard(t, "easy_1")

	solver := backtracker.NewSolver()
	solution, err := solver.Solve(b)
	if err != nil {
		t.Fatalf("expected err to be nil, got %s", err)
	}

	expected := testLoadBoard(t, "easy_1_solution")
	if solution.String() != expected.String() {
		t.Fatalf("expected solution\n%s\n, got\n%s\n", expected, solved)
	}
}

func benchmarkSolver(s sudoku.Solver, boardName string, b *testing.B) {
	board := testLoadBoard(b, boardName)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		if _, err := s.Solve(board); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkBacktrackerEasy(b *testing.B) { benchmarkSolver(backtracker.NewSolver(), "easy_1", b) }
func BenchmarkBacktrackerEvil(b *testing.B) { benchmarkSolver(backtracker.NewSolver(), "evil_1", b) }
