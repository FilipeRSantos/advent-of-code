package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed input.txt
var s string

type Coordinates struct {
	x, y int
}

type Map struct {
	frequencies map[rune][]Coordinates
	rows        int
	columns     int
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

func (m *Map) insideBoundaries(coord Coordinates) bool {
	if coord.x < 0 || coord.y < 0 {
		return false
	}

	if coord.x >= m.columns || coord.y >= m.rows {
		return false
	}

	return true
}

func (m *Map) getAntiNodes(part1 bool) map[Coordinates][]rune {
	antiNodes := make(map[Coordinates][]rune, 10)
	for key, antennas := range m.frequencies {
		for i := 0; i < len(antennas); i++ {
			for j := 0; j < len(antennas); j++ {
				if j == i {
					continue
				}

				diffY := antennas[i].y - antennas[j].y
				diffX := antennas[i].x - antennas[j].x

				if diffY < 0 {
					continue
				}

				if diffY == 0 && diffX < 0 {
					continue
				}

				iterations := 1
				for {
					anyInsideBounds := false
					coords := []Coordinates{
						{
							y: antennas[i].y + (diffY * iterations),
							x: antennas[i].x + (diffX * iterations),
						},
						{
							y: antennas[j].y - (diffY * iterations),
							x: antennas[j].x - (diffX * iterations),
						}}

					if iterations == 1 && !part1 {
						coords = append(coords, antennas[j])
						coords = append(coords, antennas[i])
					}

					for _, coord := range coords {
						if !m.insideBoundaries(coord) {
							continue
						}

						anyInsideBounds = true
						value, exists := antiNodes[coord]
						if !exists {
							antiNodes[coord] = []rune{key}
						} else {
							value = append(value, key)
							antiNodes[coord] = value
						}
					}

					if part1 || !anyInsideBounds {
						break
					}

					iterations++
				}

			}
		}
	}

	return antiNodes
}

func runStep1(input string) int {
	state := parse(input)
	antiNodes := state.getAntiNodes(true)

	return len(antiNodes)
}

func runStep2(input string) int {
	state := parse(input)
	antiNodes := state.getAntiNodes(false)

	return len(antiNodes)
}

func parse(input string) Map {
	lines := strings.Split(input, "\n")
	frequencies := make(map[rune][]Coordinates, 10)
	rows := len(lines)
	columns := len(lines[0])

	for y, line := range lines {
		for x, char := range line {
			if char == '.' || char == '#' {
				continue
			}

			value, exists := frequencies[char]
			if !exists {
				frequencies[char] = []Coordinates{{x, y}}
			} else {
				value = append(value, Coordinates{x, y})
				frequencies[char] = value
			}
		}
	}

	return Map{
		frequencies,
		rows,
		columns,
	}
}
