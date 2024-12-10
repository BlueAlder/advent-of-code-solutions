// Solution for day10 of the Advent of Code Challenge 2024
package day10

import (
	_ "embed"
	"image"
	"strings"

	"github.com/BlueAlder/advent-of-code-solutions/common/sets"
	util "github.com/BlueAlder/advent-of-code-solutions/common/utils"
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

type topographicalMap struct {
	trailheads []image.Point
	grid       [][]int
}

func parseInput(input string) topographicalMap {
	tpMap := topographicalMap{
		trailheads: []image.Point{},
		grid:       [][]int{},
	}
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		tpMap.grid = append(tpMap.grid, []int{})
		for x, char := range line {
			if char == '0' {
				tpMap.trailheads = append(tpMap.trailheads, image.Pt(x, y))
			}
			tpMap.grid[y] = append(tpMap.grid[y], util.MustAtoi(string(char)))
		}
	}
	return tpMap
}

func part1(inputData string) int {
	tm := parseInput(inputData)

	sum := 0
	for _, trailhead := range tm.trailheads {
		sum += tm.findPath(trailhead)
	}
	return sum
}

func part2(inputData string) int {
	return 0
}

func (tm *topographicalMap) findPath(start image.Point) int {
	visted := make(sets.Set[image.Point])
	next := []image.Point{start}
	count := 0
	// Basic depth first search where the next point can only be within 1 unit of the current point
	for len(next) > 0 {
		current := next[len(next)-1]
		next = next[:len(next)-1]
		if visted.Has(current) {
			continue
		}
		visted.Add(current)
		if tm.grid[current.Y][current.X] == 9 {
			count++
			continue
		}
		for _, neighbor := range tm.getNeighbors(current) {
			if !visted.Has(neighbor) {
				next = append(next, neighbor)
			}
		}
	}
	return count
}

func (tm *topographicalMap) getNeighbors(p image.Point) []image.Point {
	neighbors := []image.Point{}
	for _, offset := range []image.Point{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	} {
		neighbor := p.Add(offset)
		if neighbor.X >= 0 && neighbor.X < len(tm.grid[0]) &&
			neighbor.Y >= 0 && neighbor.Y < len(tm.grid) &&
			tm.grid[neighbor.Y][neighbor.X]-tm.grid[p.Y][p.X] == 1 {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}
