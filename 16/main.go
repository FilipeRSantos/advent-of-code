package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"strings"
)

//go:embed input.txt
var s string

var maze map[Coordinate]bool
var reindeerAt Coordinate
var finishAt Coordinate
var reindeerDirection int
var visited map[Coordinate][]int

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
	var ans int
	args := os.Args[1]

	if args == "1" {
		ans = runStep1(s)
	} else {
		ans = runStep2(s)
	}

	fmt.Println("Output: ", ans)
}

func walkMaze(coordinate, previousCoordinate Coordinate, currentDirection int, path map[Coordinate]int, scores map[int]bool) {

	turns := 0
	for _, turn := range path {
		turns += turn
	}
	score := turns*1000 + len(path)

	if coordinate == finishAt {
		scores[score] = true
		fmt.Printf("%d,", score)
		return
	}

	if coordinate != previousCoordinate {
		if _, ok := maze[coordinate]; !ok {
			return
		}
	} else {
		visited = make(map[Coordinate][]int)
	}

	if _, ok := path[coordinate]; ok {
		return
	}

	bestVisitedScore, hasVisited := visited[coordinate]
	if hasVisited {
		if bestVisitedScore[currentDirection] <= score {
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

func runStep1(input string) int {
	parse(input)

	path := make(map[Coordinate]int)
	scores := make(map[int]bool)
	walkMaze(reindeerAt, reindeerAt, reindeerDirection, path, scores)

	score := math.MaxInt32

	for x := range scores {
		if x < score {
			score = x
		}
	}

	return score
}

func runStep2(input string) int {
	return 0
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
