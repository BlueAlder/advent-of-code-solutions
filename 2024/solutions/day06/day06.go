// Solution for day06 of the Advent of Code Challenge 2024
package day06

import (
	_ "embed"
	"fmt"
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

type GridData struct {
	grid          []string
	startLocation complex128
}

func parseInput(input string) GridData {
	grid := strings.Split(input, "\n")
	for y, line := range grid {
		index := strings.Index(line, "^")
		if index != -1 {
			return GridData{grid, complex(float64(index), float64(y))}
		}
	}
	util.LogFatal("Guard location not found")
	panic(1)
}

func part1(inputData string) int {
	grid := parseInput(inputData)
	guardDirection := 0 - 1i
	visited, _ := grid.walkGrid(guardDirection)
	return len(visited)
}

type positionLocation struct {
	position  complex128
	direction complex128
}

// Obstruction must be on the path that the guard already goes on
// so test all 5030 and see if it creates a loop with a loop being defined as
// returning to the same point with the same direction.
func part2(inputData string) int {
	grid := parseInput(inputData)
	guardDirection := 0 - 1i
	visited, _ := grid.walkGrid(guardDirection)
	visited.Remove(grid.startLocation)

	total := 0
	for _, location := range visited.ToSlice() {
		if location == 1+8i {
			fmt.Println("hello")
		}
		row := grid.grid[int(imag(location))]
		row = row[:int(real(location))] + "#" + row[int(real(location))+1:]
		grid.grid[int(imag(location))] = row

		if grid.walkGridAndFindLoop(guardDirection) {
			total += 1
		}
		row = row[:int(real(location))] + "." + row[int(real(location))+1:]
		grid.grid[int(imag(location))] = row
	}
	return total
}

func (g GridData) walkGrid(guardDirection complex128) (sets.Set[complex128], sets.Set[positionLocation]) {
	guardLocation := g.startLocation
	visited := make(sets.Set[complex128])
	visitedWithDirection := make(sets.Set[positionLocation])
	for {
		visited.Add(guardLocation)
		visitedWithDirection.Add(positionLocation{guardLocation, guardDirection})
		inFront := guardLocation + guardDirection
		if real(inFront) < 0 || imag(inFront) < 0 || real(inFront) >= float64(len(g.grid[0])) || imag(inFront) >= float64(len(g.grid)) {
			break
		}
		if g.grid[int(imag(inFront))][int(real(inFront))] == '#' {
			guardDirection *= 1i
		} else {
			guardLocation = inFront
		}
	}
	return visited, visitedWithDirection
}

func (g GridData) walkGridAndFindLoop(guardDirection complex128) bool {
	guardLocation := g.startLocation
	visitedWithDirection := make(sets.Set[positionLocation])
	guard := positionLocation{guardLocation, guardDirection}
	for {
		visitedWithDirection.Add(guard)
		inFront := guard.position + guard.direction
		if real(inFront) < 0 || imag(inFront) < 0 || real(inFront) >= float64(len(g.grid[0])) || imag(inFront) >= float64(len(g.grid)) {
			return false
		}
		if g.grid[int(imag(inFront))][int(real(inFront))] == '#' {
			guard.direction *= 1i
		} else {
			guard.position = inFront
		}
		if visitedWithDirection.Has(guard) {
			return true
		}
	}
}
