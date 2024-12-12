package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed input.txt
var s string

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

type Coordinate struct {
	x, y int
}

type Region struct {
	plantType rune
	plants    map[Coordinate]bool
}

type Map struct {
	rows    int
	columns int
	coords  map[Coordinate]rune
}

func (r *Region) getCost() int {
	return len(r.plants) * r.getPerimeter()
}

func (r *Region) getPerimeter() int {
	perimeter := 0
	for coord := range r.plants {
		coords := []Coordinate{
			{x: coord.x + 1, y: coord.y},
			{x: coord.x, y: coord.y + 1},
			{x: coord.x - 1, y: coord.y},
			{x: coord.x, y: coord.y - 1},
		}

		for _, c := range coords {
			if _, exists := r.plants[c]; !exists {
				perimeter++
			}
		}
	}
	return perimeter
}

func runStep1(input string) int {
	maps := parse(input)
	regions := maps.getRegions()

	price := 0

	for _, region := range regions {
		price += region.getCost()
	}

	return price
}

func runStep2(input string) int {
	return 0
}

func parse(input string) Map {

	lines := strings.Split(input, "\n")
	rows := len(lines)
	columns := len(lines[0])
	coords := make(map[Coordinate]rune, 0)

	for row, line := range lines {
		for column, plant := range line {
			coords[Coordinate{x: column, y: row}] = plant
		}
	}

	return Map{
		rows,
		columns,
		coords,
	}
}

func (m *Map) getRegions() []Region {
	coordsChecked := make(map[Coordinate]bool)
	regions := make([]Region, 0)
	for coord, plantType := range m.coords {
		currentRegion := make(map[Coordinate]bool)
		m.checkRegion(coord, coord, currentRegion, coordsChecked)

		if len(currentRegion) == 0 {
			continue
		}

		regions = append(regions, Region{
			plantType: plantType,
			plants:    currentRegion,
		})
	}
	return regions
}

func (m *Map) withinBounds(coordinate Coordinate) bool {
	return coordinate.x >= 0 && coordinate.x < m.columns && coordinate.y >= 0 && coordinate.y < m.rows
}

func (m *Map) checkRegion(coordinate, previousCoordinate Coordinate, currRegion map[Coordinate]bool, alreadyCheckedRegions map[Coordinate]bool) {

	if _, exists := alreadyCheckedRegions[coordinate]; exists {
		return
	}

	if !m.withinBounds(coordinate) {
		return
	}

	if m.coords[previousCoordinate] != m.coords[coordinate] && previousCoordinate != coordinate {
		return
	}

	currRegion[coordinate] = true
	alreadyCheckedRegions[coordinate] = true

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
		m.checkRegion(c, coordinate, currRegion, alreadyCheckedRegions)
	}
}
