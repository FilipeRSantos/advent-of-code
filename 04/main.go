package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

const (
	None      = -1
	North     = 0
	NorthEast = 1
	East      = 2
	SouthEast = 3
	South     = 4
	SouthWest = 5
	West      = 6
	NorthWest = 7
)

//go:embed input.txt
var s string

type WordPuzzle struct {
	lines   []string
	rows    int
	columns int
}

func main() {
	var ans int
	args := os.Args[1]

	if args == "1" {
		ans = runStep1(s)
	} else {
		ans = runStep2(s)
	}

	fmt.Println("Output: ", ans)
}

func runStep1(input string) int {
	lines := strings.Split(input, "\n")
	wordPuzzle := &WordPuzzle{
		lines:   lines,
		rows:    len(lines),
		columns: len(lines[0]),
	}

	return wordPuzzle.checkPuzzle(`XMAS`)
}

func runStep2(input string) int {
	return 0
}

func (w *WordPuzzle) checkPuzzle(word string) int {
	matches := 0

	for row, line := range w.lines {
		for column, _ := range line {
			matches += w.matches(row, column, 0, word, None)
		}
	}

	return matches
}

func (w *WordPuzzle) matches(row, column, currentWordIndex int, word string, searchDirection int) int {
	if row < 0 {
		return 0
	}

	if column < 0 {
		return 0
	}

	if row >= w.rows {
		return 0
	}

	if column >= w.columns {
		return 0
	}

	currentLetter := w.lines[row][column]

	if currentLetter != word[currentWordIndex] {
		return 0
	}

	if currentWordIndex == len(word)-1 {
		return 1
	}

	matches := 0

	if searchDirection == North || searchDirection == None {
		matches += w.matches(row-1, column, currentWordIndex+1, word, North)
	}

	if searchDirection == NorthEast || searchDirection == None {
		matches += w.matches(row-1, column+1, currentWordIndex+1, word, NorthEast)
	}

	if searchDirection == East || searchDirection == None {
		matches += w.matches(row, column+1, currentWordIndex+1, word, East)
	}

	if searchDirection == SouthEast || searchDirection == None {
		matches += w.matches(row+1, column+1, currentWordIndex+1, word, SouthEast)
	}

	if searchDirection == South || searchDirection == None {
		matches += w.matches(row+1, column, currentWordIndex+1, word, South)
	}

	if searchDirection == SouthWest || searchDirection == None {
		matches += w.matches(row+1, column-1, currentWordIndex+1, word, SouthWest)
	}

	if searchDirection == West || searchDirection == None {
		matches += w.matches(row, column-1, currentWordIndex+1, word, West)
	}

	if searchDirection == NorthWest || searchDirection == None {
		matches += w.matches(row-1, column-1, currentWordIndex+1, word, NorthWest)
	}

	return matches
}
