// Solution for day02 of the Advent of Code Challenge 2024
package day02

import (
	_ "embed"
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
	reportsRaw := strings.Split(inputData, "\n")

	reportsFielded := util.MapSlice(reportsRaw, func(reportRaw string) []int {
		return util.MapSlice(strings.Fields(reportRaw), func(numStr string) int {
			return util.MustAtoi(numStr)
		})
	})

	total := len(reportsFielded)
	for _, report := range reportsFielded {
		delta := 0
		for i := 0; i < len(report)-1; i++ {
			diff := (report[i] - report[i+1])
			if util.Abs(delta+diff) > util.Abs(delta) && util.Abs(diff) < 4 {
				delta += diff
			} else {
				total--
				break
			}
		}
	}

	// fmt.Print(reportsFielded)
	return total
}

func part2(inputData string) int {
	return 0
}
