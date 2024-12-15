package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
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

func (t *Tile) moveRight() bool {
	if t.t == Wall {
		return false
	}

	coord := Coordinate{x: t.x + 1, y: t.y}
	next, exists := tiles[coord]
	if exists {
		moved := next.moveRight()
		if !moved {
			return false
		}
	}

	delete(tiles, Coordinate{x: t.x, y: t.y})
	t.x = t.x + 1
	tiles[coord] = *t

	if t.t == Robot {
		robotAt = coord
	}

	return true
}

func (t *Tile) moveLeft() bool {
	if t.t == Wall {
		return false
	}

	coord := Coordinate{x: t.x - 1, y: t.y}
	next, exists := tiles[coord]
	if exists {
		moved := next.moveLeft()
		if !moved {
			return false
		}
	}

	delete(tiles, Coordinate{x: t.x, y: t.y})
	t.x = t.x - 1
	tiles[coord] = *t

	if t.t == Robot {
		robotAt = coord
	}

	return true
}

func (t *Tile) moveTop() bool {
	if t.t == Wall {
		return false
	}

	coord := Coordinate{x: t.x, y: t.y - 1}
	next, exists := tiles[coord]
	if exists {
		moved := next.moveTop()
		if !moved {
			return false
		}
	}

	delete(tiles, Coordinate{x: t.x, y: t.y})
	t.y = t.y - 1
	tiles[coord] = *t

	if t.t == Robot {
		robotAt = coord
	}

	return true
}

func (t *Tile) moveBottom() bool {
	if t.t == Wall {
		return false
	}

	coord := Coordinate{x: t.x, y: t.y + 1}
	next, exists := tiles[coord]
	if exists {
		moved := next.moveBottom()
		if !moved {
			return false
		}
	}

	delete(tiles, Coordinate{x: t.x, y: t.y})
	t.y = t.y + 1
	tiles[coord] = *t

	if t.t == Robot {
		robotAt = coord
	}

	return true
}

func runStep2(input string) int {
	return 0
}

func runStep1(input string) int {
	parse(input)

	for _, command := range commands {
		// debug()

		if command == '\n' {
			continue
		}

		guard := tiles[robotAt]
		switch command {
		case '<':
			guard.moveLeft()
		case '>':
			guard.moveRight()
		case '^':
			guard.moveTop()
		case 'v':
			guard.moveBottom()
		}
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

func parse(input string) {
	sections := strings.Split(input, "\n\n")

	tiles = make(map[Coordinate]Tile)

	lines := strings.Split(sections[0], "\n")
	height = len(lines)
	width = len(lines[0])

	for y, line := range lines {
		for x, tile := range line {
			if tile == '.' {
				continue
			}

			coord := Coordinate{x, y}

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
				x: x,
				y: y,
				t: t,
			}
		}
	}
	commands = sections[1]
}

func debug() {
	fmt.Printf("\n")

	for y := range height {
		for x := range width {
			value, exits := tiles[Coordinate{x: x, y: y}]
			if !exits {
				fmt.Printf(".")
				continue
			}

			switch value.t {
			case Robot:
				fmt.Printf("@")

				if value.x != robotAt.x || value.y != robotAt.y {
					panic("WEWEWeE")
				}
			case Box:
				fmt.Printf("O")
			case Wall:
				fmt.Printf("#")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")

}
