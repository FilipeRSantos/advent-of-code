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

var maze map[Coordinate]bool
var reindeerAt Coordinate
var finishAt Coordinate
var reindeerDirection int
var visited map[Coordinate][]int
var bestScoreFound int

const (
	Top    = 0
	Bottom = 1
	Left   = 2
	Right  = 3
)

type Coordinate struct {
	x, y int
}

func main() {
	ans1, ans2 := solve(s)

	fmt.Printf("Part1: %d\nPart2: %d\n", ans1, ans2)
}

func walkMaze(coordinate, previousCoordinate Coordinate, currentDirection int, path map[Coordinate]int, scores map[int]map[Coordinate]bool) {

	if coordinate == previousCoordinate {
		bestScoreFound = math.MaxInt32
		visited = make(map[Coordinate][]int)
	}

	turns := 0
	for _, turn := range path {
		turns += turn
	}
	score := turns*1000 + len(path)

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
		if _, ok := maze[coordinate]; !ok {
			return
		}
	}

	if _, ok := path[coordinate]; ok {
		return
	}

	bestVisitedScore, hasVisited := visited[coordinate]
	if hasVisited {
		if bestVisitedScore[currentDirection] < score {
			return
		}
	} else {
		visited[coordinate] = make([]int, 4)
		for i := range 4 {
			visited[coordinate][i] = math.MaxInt32
		}
	}

	visited[coordinate][currentDirection] = score

	coords := []struct {
		coord     Coordinate
		direction int
	}{
		{coord: Coordinate{x: coordinate.x + 1, y: coordinate.y}, direction: Right},
		{coord: Coordinate{x: coordinate.x, y: coordinate.y - 1}, direction: Top},
		{coord: Coordinate{x: coordinate.x - 1, y: coordinate.y}, direction: Left},
		{coord: Coordinate{x: coordinate.x, y: coordinate.y + 1}, direction: Bottom},
	}

	for _, c := range coords {
		if c.coord == previousCoordinate {
			continue
		}

		turn := 0
		currentPath := make(map[Coordinate]int)

		for k, v := range path {
			currentPath[k] = v
		}

		if c.direction != currentDirection {
			turn = 1
		}
		currentPath[coordinate] = turn

		walkMaze(c.coord, coordinate, c.direction, currentPath, scores)
	}
}

func solve(input string) (int, int) {
	parse(input)

	path := make(map[Coordinate]int)
	scores := make(map[int]map[Coordinate]bool)
	walkMaze(reindeerAt, reindeerAt, reindeerDirection, path, scores)

	score := math.MaxInt32

	for x := range scores {
		if x < score {
			score = x
		}
	}

	return score, len(scores[score]) + 1
}

func parse(input string) {
	maze = make(map[Coordinate]bool)
	lines := strings.Split(input, "\n")

	for y, line := range lines {
		for x, tile := range line {
			coords := Coordinate{x: x, y: y}
			switch tile {
			case '#':
				continue
			case 'S':
				reindeerDirection = Right
				reindeerAt = coords
				continue
			case 'E':
				finishAt = coords
			case '.':
				maze[coords] = true
			}
		}
	}
}
