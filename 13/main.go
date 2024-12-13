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

type Game struct {
	aX, aY, bX, bY, prizeX, prizeY int
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

//  x < 57537

func runStep1(input string) int {
	games := parse(input)
	cost := 0

	for _, game := range games {

		currentX := 0
		currentY := 0
		bPresses := 0
		aPresses := 0
		incrementB := true
		incrementA := false
		decrementB := false

		canReachX := (game.aX*100)+(game.bX*100) >= game.prizeX
		canReachY := (game.aY*100)+(game.bY*100) >= game.prizeY

		for {
			if !canReachX || !canReachY {
				break
			}

			if currentX == game.prizeX && currentY == game.prizeY {
				break
			}

			if aPresses >= 100 && bPresses >= 100 {
				aPresses = 0
				bPresses = 0
				break
			}

			var nextX, nextY int
			if incrementB {
				nextX = currentX + game.bX
				nextY = currentY + game.bY

				if nextX <= game.prizeX && nextY <= game.prizeY && bPresses < 100 {
					bPresses++
					currentX = nextX
					currentY = nextY
				} else {
					incrementB = false
					incrementA = true
				}
				continue
			}

			if incrementA {
				nextX = currentX + game.aX
				nextY = currentY + game.aY

				if nextX <= game.prizeX && nextY <= game.prizeY && aPresses < 100 {
					aPresses++
					currentX = nextX
					currentY = nextY
				} else {
					decrementB = true
					incrementA = false
				}
				continue
			}

			if decrementB && bPresses > 0 {
				currentX -= game.bX
				currentY -= game.bY
				bPresses--

				if (currentX <= game.prizeX && currentY <= game.prizeY) || bPresses == 0 {
					decrementB = false
					incrementA = true
				}

				continue
			}

			if decrementB && bPresses == 0 {
				aPresses = 0
				bPresses = 0
				break
			}

			panic("Should never be here")
		}

		cost += aPresses*3 + bPresses
	}

	return cost
}

func runStep2(input string) int {
	return 0
}

func parse(input string) []Game {
	lines := strings.Split(input, "\n")
	output := make([]Game, (len(lines)+1)/4)

	for i := 0; i < len(lines); i += 4 {

		var index1 = strings.Index(lines[i], "Y+")
		var index2 = strings.Index(lines[i+1], "Y+")
		var index3 = strings.Index(lines[i+2], "Y=")

		output[i/4] = Game{
			aX:     maths.ParseInt(lines[i][12 : index1-2]),
			aY:     maths.ParseInt(lines[i][index1+2:]),
			bX:     maths.ParseInt(lines[i+1][12 : index2-2]),
			bY:     maths.ParseInt(lines[i+1][index2+2:]),
			prizeX: maths.ParseInt(lines[i+2][9 : index3-2]),
			prizeY: maths.ParseInt(lines[i+2][index3+2:]),
		}
	}

	return output
}
