// Solution for day11 of the Advent of Code Challenge 2023
package day11

import (
	_ "embed"
	"image"
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

func part1(inputData string) int {
	return findDistancesWithExpansionFactor(inputData, 2)
}

func part2(inputData string) int {
	return findDistancesWithExpansionFactor(inputData, 1_000_000)
}

func findDistancesWithExpansionFactor(inputData string, expansionFactor int) int {
	var u Universe = strings.Split(inputData, "\n")
	eRows := u.findEmptyRows()
	eCols := u.findEmptyColumns()

	galaxies := u.findAllGalaxies()
	dists := u.calculateDistancesWithExpansion(galaxies, expansionFactor, eRows, eCols)

	sum := 0
	for _, v := range dists {
		sum += v
	}
	return sum
}

func (u Universe) calculateDistancesWithExpansion(galaxies []image.Point, expansionFactor int, eRows, eCols []int) []int {
	var distances []int
	pairIndexes := combin.Combinations(len(galaxies), 2)
	for _, pair := range pairIndexes {
		g1 := galaxies[pair[0]]
		g2 := galaxies[pair[1]]

		eRowsCrossed := util.ReduceSlice(eRows, func(eRow int, count int) int {
			if util.EqualOrBetween(g1.Y, g2.Y, eRow) {
				return count + 1
			}
			return count
		})

		eColsCrossed := util.ReduceSlice(eCols, func(eCol int, count int) int {
			if util.EqualOrBetween(g1.X, g2.X, eCol) {
				return count + 1
			}
			return count
		})

		dist := util.Abs(g1.X-g2.X) + util.Abs(g1.Y-g2.Y) +
			(expansionFactor-1)*(eColsCrossed+eRowsCrossed)
		distances = append(distances, dist)
	}

	return distances
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

func (u Universe) findAllGalaxies() []image.Point {
	var galaxies []image.Point
	for y, row := range u {
		for x := range row {
			if string(u[y][x]) == "#" {
				galaxies = append(galaxies, image.Point{X: x, Y: y})
			}
		}
	}
	return galaxies
}
