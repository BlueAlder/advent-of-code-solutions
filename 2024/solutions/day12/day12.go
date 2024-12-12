// Solution for day12 of the Advent of Code Challenge 2024
package day12

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

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

type fencing struct {
	area      int
	perimeter int
	plantType rune
}

func part1(inputData string) int {
	grid := parseInput(inputData)

	var fences []fencing
	computed := make(sets.Set[image.Point])
	for y, row := range grid {
		for x, plant := range row {
			pt := image.Pt(x, y)
			if _, ok := computed[pt]; ok {
				continue
			}
			next := []image.Point{pt}
			fencing := fencing{plantType: plant}
			for len(next) > 0 {
				curr := next[0]
				next = next[1:]
				if _, ok := computed[curr]; ok {
					continue
				}
				fencing.area++
				neighbours, edges := getNeighbours(grid, curr)
				fencing.perimeter += edges
				next = append(next, neighbours...)
				computed.Add(curr)
			}
			fences = append(fences, fencing)
		}
	}

	return totalPrice(fences)
}

func part2(inputData string) int {
	return 0
}

func getNeighbours(grid []string, pt image.Point) (neighbours []image.Point, edges int) {
	plantType := grid[pt.Y][pt.X]
	directions := []image.Point{
		image.Pt(0, 1),
		image.Pt(0, -1),
		image.Pt(1, 0),
		image.Pt(-1, 0),
	}

	for _, dir := range directions {
		neighbour := pt.Add(dir)
		if neighbour.X < 0 || neighbour.X >= len(grid[0]) || neighbour.Y < 0 || neighbour.Y >= len(grid) || grid[neighbour.Y][neighbour.X] != plantType {
			edges++
			continue
		}
		// if _, ok := visited[neighbour]; ok {
		// 	continue
		// }
		neighbours = append(neighbours, neighbour)

	}
	return
}

func totalPrice(fences []fencing) (total int) {
	for _, fence := range fences {
		total += fence.area * fence.perimeter
	}
	return
}
