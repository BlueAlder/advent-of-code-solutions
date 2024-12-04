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

func parseInput(inputData string) [][]int {
	reportsRaw := strings.Split(inputData, "\n")
	return util.MapSlice(reportsRaw, func(reportRaw string) []int {
		return util.MapSlice(strings.Fields(reportRaw), func(numStr string) int {
			return util.MustAtoi(numStr)
		})
	})
}

func isSafeReport(report []int) bool {
	delta := 0
	for i := 0; i < len(report)-1; i++ {
		diff := (report[i] - report[i+1])
		if util.Abs(delta+diff) > util.Abs(delta) && util.Abs(diff) < 4 && delta*diff >= 0 {
			delta += diff
		} else {
			return false
		}
	}
	return true
}

func part1(inputData string) int {
	reportsFielded := parseInput(inputData)

	total := len(reportsFielded)
	for _, report := range reportsFielded {
		if !isSafeReport(report) {
			total--
		}
	}

	return total
}

func part2(inputData string) int {
	reportsFielded := parseInput(inputData)

	total := 0

ReportLoop:
	for _, report := range reportsFielded {
		if !isSafeReport(report) {
			for i := range report {
				removedReport := util.RemoveIndex(report, i)
				if isSafeReport(removedReport) {
					total++
					continue ReportLoop
				}
			}
		} else {
			total++
		}
	}
	return total
}
