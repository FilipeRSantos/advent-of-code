package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/FilipeRSantos/advent-of-code/maths"
)

//go:embed input.txt
var s string

var maze map[Coordinate]int
var playerAt Coordinate
var finishAt Coordinate

type Coordinate struct {
	x, y int
}

func main() {
	args := os.Args[1]

	if args == "1" {
		fmt.Println("Output: ", runStep1(s, 100))
	} else {
		fmt.Println("Output: ", runStep2(s))
	}
}

func doOriginalLap(coordinate, previousCoordinate Coordinate, currentStep int) {
	if _, exists := maze[coordinate]; !exists {
		return
	}

	maze[coordinate] = currentStep

	if coordinate == finishAt {
		return
	}

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

		doOriginalLap(c, coordinate, currentStep+1)
	}
}

// < 1323
func runStep1(input string, savedTime int) int {
	parse(input)
	doOriginalLap(playerAt, playerAt, 1)

	acc := 0
	for k, v := range maze {
		for kk, vv := range maze {
			diff := vv - v - 2
			if diff < savedTime {
				continue
			}

			diffX := maths.Abs(k.x - kk.x)
			diffY := maths.Abs(k.y - kk.y)

			if diffX < 3 && diffY < 3 && diffX+diffY < 3 {
				acc += 1
			}
		}
	}
	return acc
}

func runStep2(input string) int {
	return 0
}

func parse(input string) {
	maze = make(map[Coordinate]int)
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, tile := range line {
			coords := Coordinate{x: x, y: y}
			switch tile {
			case '#':
				continue
			case 'S':
				playerAt = coords
			case 'E':
				finishAt = coords
			}
			maze[coords] = 0
		}
	}
}
