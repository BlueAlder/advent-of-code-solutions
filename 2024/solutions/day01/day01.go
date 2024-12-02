// Solution for day01 of the Advent of Code Challenge 2024
package day01

import (
	_ "embed"
	
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
	return 0
}

func part2(inputData string) int {
	return 0
}

