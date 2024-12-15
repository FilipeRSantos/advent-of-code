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
var guardPos map[Coordinate]int

type Coordinate struct {
	x, y int
}

type Guard struct {
	start Coordinate
	move  Coordinate
}

func main() {
	var ans int
	args := os.Args[1]

	if args == "1" {
		ans = runStep1(s, 101, 103)
	} else {
		ans = runStep2(s)
	}

	fmt.Println("Output: ", ans)
}

func simulate(guards []Guard, width, height, steps int) int {
	guardPos = make(map[Coordinate]int)
	quadrants := make(map[int]int)

	for i := range 3 {
		quadrants[i+1] = 0
	}

	for _, guard := range guards {
		finalX := guard.move.x*steps + guard.start.x + 1
		finalY := guard.move.y*steps + guard.start.y + 1

		quadrantX := maths.Abs(finalX) / width
		quadrantY := maths.Abs(finalY) / height

		var localPosX int
		var localPosY int
		if finalX < 0 {
			localPosX = width - (-1*finalX - (width * quadrantX))

		} else {
			localPosX = finalX - (width * quadrantX)
		}

		if finalY < 0 {
			localPosY = height - (-1*finalY - (height * quadrantY))
		} else {
			localPosY = finalY - (height * quadrantY)
		}

		if localPosX == 0 {
			localPosX = width
		}

		if localPosY == 0 {
			localPosY = height
		}

		if localPosX < 1 || localPosX > width {
			panic("X should not be outside map bounds")
		}

		if localPosY < 1 || localPosY > height {
			panic("Y should not be outside map bounds")
		}

		var quadrant int

		if localPosX <= width/2 {
			if localPosY <= height/2 {
				quadrant = 1
			} else if localPosY > height/2+1 {
				quadrant = 3
			}
		} else if localPosX > width/2+1 {
			if localPosY <= height/2 {
				quadrant = 2
			} else if localPosY > height/2+1 {
				quadrant = 4
			}
		}

		value := guardPos[Coordinate{x: localPosX, y: localPosY}]
		guardPos[Coordinate{x: localPosX, y: localPosY}] = value + 1

		if quadrant > 0 {
			value = quadrants[quadrant]
			quadrants[quadrant] = value + 1
		}

	}

	acc := 1

	for _, g := range quadrants {
		acc *= g
	}

	return acc
}

func runStep1(input string, width, height int) int {
	guards := parse(input)
	return simulate(guards, width, height, 100)

}

func runStep2(input string) int {
	guards := parse(input)
	steps := 0

	for {
		width := 101
		height := 103
		simulate(guards, width, height, steps)

		fmt.Printf("Steps: %d\n", steps)
		for row := range width {
			for column := range height {
				value := guardPos[Coordinate{x: column + 1, y: row + 1}]

				if value == 0 {
					fmt.Print(".")
				} else {
					fmt.Printf("%d", value)
				}
			}
			fmt.Printf("\n")
		}

		fmt.Printf("\n\n\n\n")

		steps++
	}
}

func parse(input string) []Guard {
	lines := strings.Split(input, "\n")
	output := make([]Guard, len(lines))

	for i, line := range lines {
		values := strings.Split(line, " ")
		pValues := strings.Split(values[0][2:], ",")
		vValues := strings.Split(values[1][2:], ",")

		output[i] = Guard{
			start: Coordinate{x: maths.ParseInt(pValues[0]), y: maths.ParseInt(pValues[1])},
			move:  Coordinate{x: maths.ParseInt(vValues[0]), y: maths.ParseInt(vValues[1])},
		}
	}

	return output
}
