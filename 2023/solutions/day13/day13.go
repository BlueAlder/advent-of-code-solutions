// Solution for day13 of the Advent of Code Challenge 2023
package day13

import (
	_ "embed"
	"slices"
	"strings"

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

func part1(inputData string) int {
	return summarizeReflection(inputData, false)
}

func part2(inputData string) int {
	return summarizeReflection(inputData, true)
}

func summarizeReflection(inputData string, smudge bool) (total int) {
	grids := util.MapSlice(strings.Split(inputData, "\n\n"), func(el string) []string { return strings.Split(el, "\n") })

	for _, grid := range grids {
		rowIdx := findReflection(grid, smudge)
		if rowIdx != -1 {
			total += rowIdx * 100
			continue
		}

		colIdx := findReflection(rotateGrid(grid), smudge)
		if colIdx != -1 {
			total += colIdx
			continue
		}
		panic("no reflection found")
	}
	return
}

func findReflection(grid []string, smudge bool) int {
	for rIdx := 1; rIdx < len(grid); rIdx++ {
		left := grid[:rIdx]
		right := grid[rIdx:]
		min := util.Min(len(left), len(right))

		left = left[len(left)-min:]
		right = right[:min]
		rightCopy := make([]string, len(right))
		copy(rightCopy, right)
		slices.Reverse(rightCopy)

		leftStr := strings.Join(left, "")
		rightStr := strings.Join(rightCopy, "")

		if (smudge && util.HammingDistanceString(leftStr, rightStr) == 1) || (!smudge && leftStr == rightStr) {
			return rIdx
		}
	}
	return -1
}

// Rotates 90 degrees
func rotateGrid(grid []string) []string {
	newRowLength := len(grid[0])
	rotated := make([]string, newRowLength)
	for i := 0; i < newRowLength; i++ {
		for j := range grid {
			rotated[i] += string(grid[len(grid)-1-j][i])
		}
	}
	return rotated
}
