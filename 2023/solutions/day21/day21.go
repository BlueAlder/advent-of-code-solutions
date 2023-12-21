// Solution for day21 of the Advent of Code Challenge 2023
package day21

import (
	_ "embed"
	"image"
	"strings"

	"github.com/BlueAlder/advent-of-code-solutions/pkg/sets"
	util "github.com/BlueAlder/advent-of-code-solutions/pkg/utils"
)

//go:embed input.txt
var input string

func Solve(part int) int {
	if part == 1 {
		return part1(input)
	} else if part == 2 {
		return part2(input)
	} else {
		util.LogFatal("invalid part number")
		return -1
	}
}

type Grid struct {
	plots []string
	rocks sets.Set[image.Point]
}

func part1(inputData string) int {
	var plots = strings.Split(inputData, "\n")
	rocks := make(sets.Set[image.Point])
	grid := Grid{plots, rocks}
	var s image.Point

	for y, row := range grid.plots {
		for x := range row {
			switch grid.plots[y][x] {
			case '#':
				rocks.Add(image.Point{X: x, Y: y})
			case 'S':
				s = image.Point{X: x, Y: y}
			}
		}
	}

	steps := 0
	stepCount := 64
	possibilities := make(sets.Set[image.Point])
	possibilities.Add(s)
	for steps < stepCount {
		steps++

		newPossibilities := make(sets.Set[image.Point])
		for p := range possibilities {
			newPossibilities.Add(grid.findNeighbours(p)...)
		}
		possibilities = newPossibilities
	}

	return len(possibilities)
}

func (g Grid) findNeighbours(p image.Point) []image.Point {
	var neigbours []image.Point

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, dir := range dirs {
		np := image.Point{X: p.X + dir[0], Y: p.Y + dir[1]}
		if !g.rocks.Has(np) && g.InBounds(np) {
			neigbours = append(neigbours, np)
		}
	}
	return neigbours

}

func (g Grid) InBounds(p image.Point) bool {
	return p.X >= 0 && p.X < len(g.plots[0]) && p.Y >= 0 && p.Y < len(g.plots)
}

func part2(inputData string) int {
	return 0
}
