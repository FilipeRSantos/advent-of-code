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

type Coordinate struct {
	x, y int
}

type Map struct {
	trailheads []Coordinate
	finishes   []Coordinate
	rows       int
	columns    int
	heights    map[Coordinate]int
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
	maps := parse(input)

	acc := 0
	for _, coord := range maps.trailheads {
		finishes := make(map[Coordinate]int, 0)
		maps.climb(coord, coord, finishes)

		acc += len(finishes)
	}

	return acc
}

func runStep2(input string) int {
	maps := parse(input)

	acc := 0
	for _, coord := range maps.trailheads {
		finishes := make(map[Coordinate]int, 0)
		maps.climb(coord, coord, finishes)

		for _, x := range finishes {
			acc += x
		}
	}

	return acc
}

func parse(input string) Map {
	lines := strings.Split(input, "\n")

	rows := len(lines)
	columns := len(lines[0])

	trailheads := make([]Coordinate, 0)
	finishes := make([]Coordinate, 0)
	heights := make(map[Coordinate]int, rows*columns)

	for row, line := range lines {
		for column, h := range line {
			var height int
			if h == '.' {
				height = -1
			} else {
				height = maths.ParseInt(string(h))
			}

			coord := Coordinate{x: column, y: row}

			if height == 0 {
				trailheads = append(trailheads, coord)
			} else if height == 9 {
				finishes = append(finishes, coord)
			}

			heights[coord] = height

		}
	}

	return Map{
		trailheads,
		finishes,
		rows,
		columns,
		heights,
	}
}

func (m *Map) withinBounds(coordinate Coordinate) bool {
	return coordinate.x >= 0 && coordinate.x < m.columns && coordinate.y >= 0 && coordinate.y < m.rows
}

func (m *Map) climb(coordinate, previousCoordinate Coordinate, finishes map[Coordinate]int) {

	if !m.withinBounds(coordinate) {
		return
	}

	if m.heights[previousCoordinate] != m.heights[coordinate]-1 && previousCoordinate != coordinate {
		return
	}

	if m.heights[coordinate] == 9 {
		value, exists := finishes[coordinate]

		if !exists {
			finishes[coordinate] = 1
		} else {
			finishes[coordinate] = value + 1
		}

		return
	}

	coords := []Coordinate{
		{x: coordinate.x + 1, y: coordinate.y},
		{x: coordinate.x, y: coordinate.y + 1},
		{x: coordinate.x - 1, y: coordinate.y},
		{x: coordinate.x, y: coordinate.y - 1},
	}

	for _, c := range coords {
		if c == previousCoordinate {
			continue
		}
		m.climb(c, coordinate, finishes)
	}
}
