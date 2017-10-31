# Sudoku Solver

A simple sudoku puzzle solver written in Go. Puzzles are read from stdin or from a file
in the following format:

```
1 _ 3 _ _ 6 _ 8 _
_ 5 _ _ 8 _ 1 2 _
7 _ 9 1 _ 3 _ 5 6
_ 3 _ _ 6 7 _ 9 _
5 _ 7 8 _ _ _ 3 _
8 _ 1 _ 3 _ 5 _ 7
_ 4 _ _ 7 8 _ 1 _
6 _ 8 _ _ 2 _ 4 _
_ 1 2 _ 4 5 _ 7 8
```

The solution is printed to stdout. 

Currently, the following sudoku solving algorithms are implemented.

1. *backtracker* - A brute force recursive backtracking algorithm. In sequence, cells
   are assigned the numbers 1 through 9 and checked if it leads to a valid solution.

# Usage

```
Usage of ./cmd/sudoku/sudoku:
  -file string
        The path to file containing sudoku puzzle. Otherwise, read from stdin.
  -solver string
        The sudoku solver to use. Only backtracker is available right now. (default "backtracker")
```
