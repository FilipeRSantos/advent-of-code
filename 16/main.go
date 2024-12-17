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

const (
	Top    = 0
	Bottom = 1
	Left   = 2
	Right  = 3
)

type Coordinate struct {
	x, y int
}

type Finish struct {
	steps, turns int
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

func walkMaze(coordinate, previousCoordinate Coordinate, currentDirection int, path map[Coordinate]int, finishes map[Finish]bool) {

	if coordinate == finishAt {
		turns := 0
		for _, turn := range path {
			turns += turn
		}
		finishes[Finish{steps: len(path), turns: turns}] = true
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

	coords := []struct {
		coord     Coordinate
		direction int
	}{
		{coord: Coordinate{x: coordinate.x + 1, y: coordinate.y}, direction: Right},
		{coord: Coordinate{x: coordinate.x - 1, y: coordinate.y}, direction: Left},
		{coord: Coordinate{x: coordinate.x, y: coordinate.y - 1}, direction: Top},
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

		walkMaze(c.coord, coordinate, c.direction, currentPath, finishes)
	}
}

func runStep1(input string) int {
	parse(input)

	path := make(map[Coordinate]int)
	finishes := make(map[Finish]bool)
	walkMaze(reindeerAt, reindeerAt, reindeerDirection, path, finishes)

	score := math.MaxInt16

	for finish := range finishes {
		x := finish.turns*1000 + finish.steps

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
	for y, line := range strings.Split(input, "\n") {
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
