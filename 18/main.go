package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"

	"github.com/FilipeRSantos/advent-of-code/maths"
)

//go:embed input.txt
var s string

var maze map[Coordinate]int
var playerAt Coordinate
var finishAt Coordinate
var visited map[Coordinate]int
var bestScoreFound int
var corruptedQtd int

type Coordinate struct {
	x, y int
}

func main() {
	ans1, ans2 := solve(s, 71, 1024)

	fmt.Printf("Part1: %d\nPart2: %d\n", ans1, ans2)
}

func walkMaze(coordinate, previousCoordinate Coordinate, path map[Coordinate]bool, scores map[int]map[Coordinate]bool) {

	if coordinate == previousCoordinate {
		bestScoreFound = math.MaxInt32
		visited = make(map[Coordinate]int)
	}

	score := len(path)
	if score > bestScoreFound {
		return
	}

	if coordinate == finishAt {
		bestScoreFound = maths.Min(bestScoreFound, score)

		_, exists := scores[score]
		if !exists {
			scores[score] = make(map[Coordinate]bool)
		}

		for c := range path {
			scores[score][c] = true
		}
		return
	}

	if coordinate != previousCoordinate {
		v, ok := maze[coordinate]

		if ok && v < corruptedQtd {
			return
		}
	}

	if _, ok := path[coordinate]; ok {
		return
	}

	bestVisitedScore, hasVisited := visited[coordinate]
	if hasVisited {
		if bestVisitedScore <= score {
			return
		}
	}

	visited[coordinate] = score

	coords := []Coordinate{
		{x: coordinate.x + 1, y: coordinate.y},
		{x: coordinate.x, y: coordinate.y - 1},
		{x: coordinate.x - 1, y: coordinate.y},
		{x: coordinate.x, y: coordinate.y + 1},
	}

	for _, c := range coords {
		if c == previousCoordinate {
			continue
		}

		if c.x < 0 || c.y < 0 {
			continue
		}

		if c.x > finishAt.x || c.y > finishAt.y {
			continue
		}

		currentPath := make(map[Coordinate]bool)

		for k, v := range path {
			currentPath[k] = v
		}

		currentPath[coordinate] = true
		walkMaze(c, coordinate, currentPath, scores)
	}
}

func solve(input string, size, corrupted int) (int, int) {
	parse(input)
	playerAt = Coordinate{0, 0}
	finishAt = Coordinate{size - 1, size - 1}
	corruptedQtd = corrupted

	path := make(map[Coordinate]bool)
	scores := make(map[int]map[Coordinate]bool)
	walkMaze(playerAt, playerAt, path, scores)

	score := math.MaxInt32

	for x := range scores {
		if x < score {
			score = x
		}
	}

	return score, len(scores[score]) + 1
}

func parse(input string) {
	maze = make(map[Coordinate]int)

	for i, group := range strings.Split(input, "\n") {
		values := strings.Split(group, ",")
		coords := Coordinate{x: maths.ParseInt(values[0]), y: maths.ParseInt(values[1])}
		maze[coords] = i
	}
}
