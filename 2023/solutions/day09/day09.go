// Solution for day09 of the Advent of Code Challenge 2023
package day09

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

func part1(inputData string) (total int) {
	nums := parseInput(inputData)
	for _, line := range nums {
		total += getNextValue(line)
	}
	return
}

func part2(inputData string) (total int) {
	nums := parseInput(inputData)
	for _, line := range nums {
		total += getPrevValue(line)
	}
	return
}

func getNextValue(line []int) int {
	if util.SliceEvery(line, func(el int) bool { return el == 0 }) {
		return 0
	}
	ds := getDeltaArray(line)
	d := getNextValue(ds)
	return line[len(line)-1] + d
}

func getPrevValue(line []int) int {
	if util.SliceEvery(line, func(el int) bool { return el == 0 }) {
		return 0
	}
	ds := getDeltaArray(line)
	d := getPrevValue(ds)
	return line[0] - d
}

func getDeltaArray(line []int) []int {
	delta := make([]int, len(line)-1)
	for i := 0; i < len(line)-1; i++ {
		delta[i] = line[i+1] - line[i]
	}
	return delta
}

func parseInput(inputData string) [][]int {
	lines := strings.Split(inputData, "\n")
	return util.MapSlice(lines, func(line string) []int {
		numString := strings.Split(line, " ")
		numArr, _ := util.MapSliceWithError(numString, strconv.Atoi)
		return numArr
	})
}
