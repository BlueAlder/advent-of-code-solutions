// Solution for day01 of the Advent of Code Challenge 2025
package day01

import (
	_ "embed"
	"strconv"
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

func parseInput(inputData string) []int {
	lines := strings.Split(inputData, "\n")
	var result []int
	for _, line := range lines {
		dir := line[0]
		steps, _ := strconv.Atoi(line[1:])
		if dir == 'R' {
			result = append(result, steps)
		} else {
			result = append(result, -steps)
		}
	}
	return result
}

func part1(inputData string) int {
	steps := parseInput(inputData)

	dial := 50
	count := 0
	for _, step := range steps {
		dial = (dial + step) % 100
		if dial == 0 {
			count++
		}
	}
	return count
}

func part2(inputData string) int {
	steps := parseInput(inputData)

	dial := 50
	count := 0
	// how many times does the dial pass 0 on any click
	for _, step := range steps {
		dial += step
		if dial == 0 {
			count++
		}
		for dial >= 100 {
			dial -= 100
			count++
		}
		for dial < 0 {
			dial += 100
			count++
		}
	}
	return count
}
