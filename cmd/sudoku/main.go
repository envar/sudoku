package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/envar/sudoku"
	"github.com/envar/sudoku/backtracker"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	solverFlag := flag.String("solver", "backtracker", "The sudoku solver to use. Only backtracker is available right now.")
	fileFlag := flag.String("file", "", "The path to file containing sudoku puzzle. Otherwise, read from stdin.")
	flag.Parse()

	r := os.Stdin // default
	if *fileFlag != "" {
		var err error
		r, err = os.Open(*fileFlag)
		if err != nil {
			return err
		}
	}

	var solver sudoku.Solver
	switch *solverFlag {
	case "backtracker":
		solver = backtracker.NewSolver()
	default:
		return fmt.Errorf("unknown solver: %s", *solverFlag)
	}

	b, err := sudoku.NewBoardFromReader(r)
	if err != nil {
		return err
	}

	solution, err := solver.Solve(b)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, solution.String())

	return nil
}
