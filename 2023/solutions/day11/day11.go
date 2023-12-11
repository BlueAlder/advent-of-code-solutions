// Solution for day11 of the Advent of Code Challenge 2023
package day11

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

	util "github.com/BlueAlder/advent-of-code-solutions/pkg/utils"
	"gonum.org/v1/gonum/stat/combin"
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

type Universe []string

type Point struct {
	x int
	y int
}

func part1(inputData string) int {
	var u Universe = strings.Split(inputData, "\n")
	eRows := u.findEmptyRows()
	eCols := u.findEmptyColumns()

	fmt.Println(eRows)
	fmt.Println(eCols)

	u = u.expandUniverse(eRows, eCols)
	galaxies := u.findAllGalaxies()
	dists := u.calculateDistances(galaxies)

	sum := 0
	for _, v := range dists {
		sum += v
	}

	return sum
}

func (u Universe) findEmptyRows() []int {
	var emptyRows []int
	for idx, row := range u {
		if !strings.Contains(row, "#") {
			emptyRows = append(emptyRows, idx)
		}
	}
	return emptyRows
}

func (u Universe) findEmptyColumns() []int {
	var emptyColumns []int

Column:
	for i := 0; i < len(u[0]); i++ {
		for _, row := range u {
			if row[i] != '.' {
				continue Column
			}
		}
		emptyColumns = append(emptyColumns, i)
	}
	return emptyColumns
}

func (u Universe) expandUniverse(rowIdxs, colIdxs []int) Universe {
	// Add cols
	for added, colIdx := range colIdxs {
		for i, row := range u {
			updatedRow := row[:colIdx+added] + "." + row[colIdx+added:]
			u[i] = updatedRow
		}
	}

	// Add rows

	rowLen := len(u[0])
	emptyRow := strings.Repeat(".", rowLen)
	for added, rowIdx := range rowIdxs {
		u = slices.Insert(u, rowIdx+added, emptyRow)
	}

	for _, line := range u {
		fmt.Println(line)
	}

	return u

}

func (u Universe) findAllGalaxies() []Point {
	var galaxies []Point
	for y, row := range u {
		for x := range row {
			if string(u[y][x]) == "#" {
				galaxies = append(galaxies, Point{x: x, y: y})
			}
		}
	}
	return galaxies
}

func (u Universe) calculateDistances(galaxies []Point) []int {
	var distances []int
	pairIndexes := combin.Combinations(len(galaxies), 2)
	for _, pair := range pairIndexes {
		g1 := galaxies[pair[0]]
		g2 := galaxies[pair[1]]
		dist := Abs(g1.x-g2.x) + Abs(g1.y-g2.y)
		distances = append(distances, dist)
	}

	return distances
}

func part2(inputData string) int {
	return 0
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
