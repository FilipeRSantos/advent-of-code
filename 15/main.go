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
var tiles map[Coordinate]Tile
var robotAt Coordinate
var commands string
var width, height int

const (
	Box   = 0
	Wall  = 1
	Robot = 2
)

const (
	Top    = 0
	Bottom = 1
	Left   = 2
	Right  = 3
)

type Coordinate struct {
	x, y int
}

type Tile struct {
	x, y, t int
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

func (t *Tile) move(part1 bool, direction int) bool {
	if t.t == Wall {
		return false
	}

	to, coords := t.getCoordsToCheck(part1, direction)

	for i, coord := range coords {
		next, exists := tiles[coord]
		if exists {
			moved := next.move(part1, direction)
			if !moved {
				return false
			}
		}

		if i == len(coords)-1 {
			delete(tiles, Coordinate{x: t.x, y: t.y})
			switch direction {
			case Top:
				t.y = t.y - 1
			case Bottom:
				t.y = t.y + 1
			case Left:
				t.x = t.x - 1
			case Right:
				t.x = t.x + 1
			}
			tiles[to] = *t
			if t.t == Robot {
				robotAt = to
			}
		}
	}

	return true
}

func (t *Tile) getCoordsToCheck(part1 bool, direction int) (Coordinate, []Coordinate) {
	var to Coordinate
	var coords []Coordinate

	switch direction {
	case Top:
		to = Coordinate{x: t.x, y: t.y - 1}
	case Bottom:
		to = Coordinate{x: t.x, y: t.y + 1}
	case Left:
		to = Coordinate{x: t.x - 1, y: t.y}
	case Right:
		to = Coordinate{x: t.x + 1, y: t.y}
	}

	if part1 {
		return to, []Coordinate{to}
	}

	switch direction {
	case Top:
		if t.t != Robot {
			coords = []Coordinate{{x: t.x - 1, y: t.y - 1}, {x: t.x + 1, y: t.y - 1}, to}
		} else {
			coords = []Coordinate{{x: t.x - 1, y: t.y - 1}, to}
		}
	case Bottom:
		if t.t != Robot {
			coords = []Coordinate{{x: t.x - 1, y: t.y + 1}, {x: t.x + 1, y: t.y + 1}, to}
		} else {
			coords = []Coordinate{{x: t.x - 1, y: t.y + 1}, to}
		}
	case Left:
		coords = []Coordinate{{x: t.x - 2, y: t.y}, to}
	case Right:
		coords = []Coordinate{{x: t.x + 2, y: t.y}, to}
	}

	return to, coords
}

func (t *Tile) canMove(direction int) bool {
	if t.t == Wall {
		return false
	}

	canMoveFlag := 0
	_, coords := t.getCoordsToCheck(false, direction)

	for _, coord := range coords {
		next, exists := tiles[coord]
		if exists {
			if next.canMove(direction) {
				canMoveFlag++
			}
		} else {
			canMoveFlag++
		}
	}

	return canMoveFlag == len(coords)
}

func runStep2(input string) int {
	parse(input, false)

	for _, command := range commands {
		debug(false, command)

		if command == '\n' {
			continue
		}

		var direction int
		guard := tiles[robotAt]
		switch command {
		case '<':
			direction = Left
		case '>':
			direction = Right
		case '^':
			direction = Top
		case 'v':
			direction = Bottom
		}
		if guard.canMove(direction) {
			guard.move(false, direction)
		}
	}
	debug(false, ' ')

	acc := 0
	for _, tile := range tiles {
		if tile.t != Box {
			continue
		}

		distanceFromBottom := height - tile.y
		distanceFromRight := width - tile.x + 1

		acc += 100*maths.Min(distanceFromBottom, tile.y) + maths.Min(tile.x, distanceFromRight)
	}

	return acc
}

func runStep1(input string) int {
	parse(input, true)

	for _, command := range commands {
		if command == '\n' {
			continue
		}

		var direction int
		guard := tiles[robotAt]
		switch command {
		case '<':
			direction = Left
		case '>':
			direction = Right
		case '^':
			direction = Top
		case 'v':
			direction = Bottom
		}
		guard.move(true, direction)
	}

	acc := 0
	for _, tile := range tiles {
		if tile.t != Box {
			continue
		}

		acc += 100*tile.y + tile.x
	}

	return acc
}

func parse(input string, part1 bool) {
	sections := strings.Split(input, "\n\n")

	tiles = make(map[Coordinate]Tile)

	lines := strings.Split(sections[0], "\n")
	height = len(lines)
	width = len(lines[0])

	if !part1 {
		width *= 2
	}

	for y, line := range lines {
		for x, tile := range line {
			if tile == '.' {
				continue
			}

			xx := x

			if !part1 {
				xx *= 2
			}

			coord := Coordinate{xx, y}

			var t int
			if tile == '#' {
				t = Wall
			}

			if tile == 'O' {
				t = Box
			}

			if tile == '@' {
				t = Robot
				robotAt = coord
			}

			tiles[coord] = Tile{
				x: coord.x,
				y: coord.y,
				t: t,
			}
		}
	}
	commands = sections[1]
}

func debug(part1 bool, command rune) {
	fmt.Printf("\n%s\n", string(command))
	lastSymbol := -1
	for y := range height {
		for x := range width {
			value, exits := tiles[Coordinate{x: x, y: y}]
			if !exits {
				if lastSymbol == -1 {
					fmt.Printf(".")
				}

				if !part1 {
					if lastSymbol == Wall {
						fmt.Printf("#")
					}

					if lastSymbol == Box {
						fmt.Printf("]")
					}

					lastSymbol = -1
				}
				continue
			}

			switch value.t {
			case Robot:
				fmt.Printf("@")
				lastSymbol = -1

				if value.x != robotAt.x || value.y != robotAt.y {
					panic("WEWEWeE")
				}
			case Box:
				if part1 {
					fmt.Printf("O")
				} else {
					lastSymbol = Box
					fmt.Printf("[")
				}
			case Wall:
				if !part1 {
					lastSymbol = Wall
				}
				fmt.Printf("#")
			}
		}
		lastSymbol = -1
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
