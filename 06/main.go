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

const (
	North = 0
	South = 1
	East  = 2
	West  = 3
)

const (
	Empty       = 0
	Obstacle    = -1
	Obstruction = -2
)

type Map struct {
	state               map[Coordinates]int
	guardStartAt        Coordinates
	guardStartDirection int
	guardAt             Coordinates
	guardDirection      int
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

func parse(input string) *Map {
	lines := strings.Split(input, "\n")
	state := make(map[Coordinates]int, len(lines)*len(lines[0]))
	guardDirection := -1
	guardAt := Coordinates{x: -1, y: -1}

	for row, line := range lines {
		for column, value := range line {
			coord := Coordinates{y: row, x: column}

			if value == '^' {
				guardDirection = North
			} else if value == '>' {
				guardDirection = East
			} else if value == 'v' {
				guardDirection = South
			} else if value == '<' {
				guardDirection = West
			}

			if guardDirection != -1 && guardAt.x == -1 {
				guardAt = coord
				state[coord] = 1
				continue
			}

			if value == '#' {
				state[coord] = Obstacle
				continue
			}

			state[coord] = Empty
		}
	}

	return &Map{
		state:               state,
		guardAt:             guardAt,
		guardDirection:      guardDirection,
		guardStartAt:        guardAt,
		guardStartDirection: guardDirection,
	}
}

func (m *Map) resetState() {
	for key, value := range m.state {
		if value > 0 || value == Obstruction {
			m.state[key] = Empty
		}
	}

	m.state[m.guardStartAt] = 1
	m.guardAt = m.guardStartAt
	m.guardDirection = m.guardStartDirection
}

func (m *Map) getNewCoord(direction int) Coordinates {
	switch direction {
	case North:
		return Coordinates{y: m.guardAt.y - 1, x: m.guardAt.x}
	case South:
		return Coordinates{y: m.guardAt.y + 1, x: m.guardAt.x}
	case East:
		return Coordinates{y: m.guardAt.y, x: m.guardAt.x + 1}
	case West:
		return Coordinates{y: m.guardAt.y, x: m.guardAt.x - 1}
	default:
		return Coordinates{}
	}
}

func (m *Map) getStepDirection(direction int) int {
	coord := m.getNewCoord(direction)

	value, exists := m.state[coord]
	if exists && (value == Obstacle || value == Obstruction) {
		var newDirection int
		switch direction {
		case North:
			newDirection = East
		case South:
			newDirection = West
		case East:
			newDirection = South
		case West:
			newDirection = North
		}
		return m.getStepDirection(newDirection)
	}

	m.guardDirection = direction
	return direction
}

func (m *Map) walk() bool {
	m.getStepDirection(m.guardDirection)
	newCoords := m.getNewCoord(m.guardDirection)

	value, exists := m.state[newCoords]
	if !exists {
		return false
	}

	m.guardAt = newCoords
	m.state[newCoords] = value + 1
	return true
}

func (m *Map) isStuckInLoop() bool {
	for {
		if !m.walk() {
			return false
		}

		value := m.state[m.guardAt]
		if value > 3 {
			return true
		}
	}
}

func runStep1(input string) int {
	state := parse(input)

	for {
		if !state.walk() {
			break
		}
	}

	visited := 0
	for _, value := range state.state {
		if value > 0 {
			visited++
		}
	}

	return visited
}

func runStep2(input string) int {
	state := parse(input)

	// First we simulate the path the Guard will take to
	// Only interact with tiles he is already crossing
	for {
		if !state.walk() {
			break
		}
	}

	var visitedCoords []Coordinates
	for key, value := range state.state {
		if value > 0 && key != state.guardStartAt {
			visitedCoords = append(visitedCoords, key)
		}
	}

	obstructions := 0
	for _, coord := range visitedCoords {
		state.state[coord] = Obstruction
		if state.isStuckInLoop() {
			obstructions++
		}

		state.resetState()
	}

	return obstructions
}
