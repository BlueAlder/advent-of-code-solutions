// Solution for day01 of the Advent of Code Challenge 2024
package day01

import (
	_ "embed"
	"slices"
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

func parseInput(inputData string) ([]int, []int) {
	lines := strings.Split(inputData, "\n")
	var left []int
	var right []int
	for _, line := range lines {
		numbers := strings.Fields(line)
		l, _ := strconv.Atoi(numbers[0])
		left = append(left, l)
		r, _ := strconv.Atoi(numbers[1])
		right = append(right, r)
	}
	return left, right
}

func part1(inputData string) int {
	left, right := parseInput(inputData)
	slices.Sort(left)
	slices.Sort(right)

	if len(left) != len(right) {
		util.LogFatal("left and right slices are not the same length")
	}

	total := 0
	for i := 0; i < len(left); i++ {
		total += util.Abs(left[i] - right[i])
	}

	return total
}

func part2(inputData string) int {
	left, right := parseInput(inputData)
	slices.Sort(left)
	slices.Sort(right)

	total := 0
	for _, location := range left {
		occurances := 0
		for _, r := range right {
			if location == r {
				occurances++
			}
			if r > location {
				break
			}
		}
		total += location * occurances
	}
	return total
}
