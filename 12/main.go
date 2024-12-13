package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed input.txt
var s string

const (
	North = 0
	East  = 1
	South = 2
	West  = 3
)

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

func (r *Region) getCost(part1 bool) int {
	if part1 {
		return len(r.plants) * r.getPerimeter()
	}

	return len(r.plants) / 4 * r.getCorners()
}

func (r *Region) getCorners() int {

	edges := r.getEdges()
	corners := make(map[Coordinate][]Coordinate)
	checkedEdges := make(map[Coordinate]bool)

	for k, v := range edges {
		for _, vv := range v {
			value, exists := corners[vv]
			if !exists {
				corners[vv] = []Coordinate{k}
			} else {
				corners[vv] = append(value, k)
			}
		}
	}

	acc := 0
	for _, edge := range edges {
		switch len(edge) {
		case 1:
			continue
		case 2:
			for _, kk := range edge {
				_, exist := checkedEdges[kk]
				if exist {
					acc--
					continue
				} else {
					checkedEdges[kk] = true
				}
			}

			acc++
		default:
			panic("Should never happen")
		}
	}

	for _, edges := range corners {
		if len(edges) == 2 {
			acc++
		}
	}

	return acc
}

func (r *Region) getEdges() map[Coordinate][]Coordinate {
	edges := make(map[Coordinate][]Coordinate)

	for coord := range r.plants {
		coords := []Coordinate{
			{x: coord.x + 1, y: coord.y},
			{x: coord.x, y: coord.y + 1},
			{x: coord.x - 1, y: coord.y},
			{x: coord.x, y: coord.y - 1},
		}

		for _, c := range coords {
			if _, ok := r.plants[c]; !ok {
				value, exists := edges[coord]
				if !exists {
					edges[coord] = []Coordinate{c}
				} else {
					edges[coord] = append(value, c)
				}
			}
		}
	}

	return edges
}

func (r *Region) getPerimeter() int {
	perimeter := 0
	for _, edges := range r.getEdges() {
		perimeter += len(edges)
	}
	return perimeter
}

func runStep1(input string) int {
	maps := parse(input, true)
	regions := maps.getRegions()

	price := 0

	for _, region := range regions {
		price += region.getCost(true)
	}

	return price
}

func runStep2(input string) int {
	maps := parse(input, false)
	regions := maps.getRegions()

	price := 0
	for _, region := range regions {
		price += region.getCost(false)
	}

	return price
}

func parse(input string, part1 bool) Map {

	lines := strings.Split(input, "\n")
	rows := len(lines) * 2
	columns := len(lines[0]) * 2
	coords := make(map[Coordinate]rune, 0)

	for row, line := range lines {
		for column, plant := range line {

			if part1 {
				coords[Coordinate{x: column, y: row}] = plant
				continue
			}

			coords[Coordinate{x: column * 2, y: row * 2}] = plant
			coords[Coordinate{x: column*2 + 1, y: row * 2}] = plant
			coords[Coordinate{x: column * 2, y: row*2 + 1}] = plant
			coords[Coordinate{x: column*2 + 1, y: row*2 + 1}] = plant
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
		if c.x == previousCoordinate.x && c.y == previousCoordinate.y {
			continue
		}
		m.checkRegion(c, coordinate, currRegion, alreadyCheckedRegions)
	}
}
