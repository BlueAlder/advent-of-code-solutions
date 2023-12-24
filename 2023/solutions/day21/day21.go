// Solution for day21 of the Advent of Code Challenge 2023
package day21

import (
	_ "embed"
	"fmt"
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
	start image.Point
}

func part1(inputData string) int {
	grid := parseInput(inputData)
	return grid.findUniquePlots(64, false)
}

func part2(inputData string) int {
	grid := parseInput(inputData)

	for i := 1; i < 0; i++ {
		fmt.Printf("Steps: %d Plots: %d\n", i, grid.findUniquePlotsInfinite(i))
	}
	fmt.Printf("Steps: %d Plots: %d\n", 5000, grid.findUniquePlotsInfinite(5000))
	return grid.findUniquePlotsInfinite(500)
}

func parseInput(inputData string) Grid {
	var plots = strings.Split(inputData, "\n")
	rocks := make(sets.Set[image.Point])
	var s image.Point

	for y, row := range plots {
		for x := range row {
			switch plots[y][x] {
			case '#':
				rocks.Add(image.Point{X: x, Y: y})
			case 'S':
				s = image.Point{X: x, Y: y}
				plots[y] = plots[y][:x] + "." + plots[y][x+1:]
			}
		}
	}
	return Grid{plots, rocks, s}
}

func (g Grid) findUniquePlots(numSteps int, infinite bool) int {
	steps := 0
	possibilities := make(sets.Set[image.Point])
	possibilities.Add(g.start)
	for steps < numSteps {
		steps++

		newPossibilities := make(sets.Set[image.Point])
		for p := range possibilities {
			if infinite {
				newPossibilities.Add(g.findNeighboursInfinite(p)...)
			} else {
				newPossibilities.Add(g.findNeighbours(p)...)
			}
		}
		possibilities = newPossibilities
	}
	return len(possibilities)
}

func (g Grid) findUniquePlotsInfinite(numSteps int) int {
	steps := 0
	possibilities := make(sets.Set[image.Point])
	possibilities.Add(g.start)

	toVisit := make(sets.Set[image.Point])
	toVisit.Add(g.start)
	for steps < numSteps {
		steps++

		nextVisit := make(sets.Set[image.Point])
		for p := range toVisit {
			for _, n := range g.findNeighboursInfinite(p) {
				if !possibilities.Has(n) {
					nextVisit.Add(n)
					possibilities.Add(n)
				}
			}
		}
		toVisit = nextVisit
	}
	return len(possibilities) / 2
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

func (g Grid) findNeighboursInfinite(p image.Point) []image.Point {
	var neigbours []image.Point

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, dir := range dirs {
		np := image.Point{X: p.X + dir[0], Y: p.Y + dir[1]}
		if !g.rocks.Has(image.Point{X: util.Mod(np.X, len(g.plots[0])), Y: util.Mod(np.Y, len(g.plots))}) {
			neigbours = append(neigbours, np)
		}
	}
	return neigbours
}

func (g Grid) InBounds(p image.Point) bool {
	return p.X >= 0 && p.X < len(g.plots[0]) && p.Y >= 0 && p.Y < len(g.plots)
}
