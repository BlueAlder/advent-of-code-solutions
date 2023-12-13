// Solution for day12 of the Advent of Code Challenge 2023
package day12

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	util "github.com/BlueAlder/advent-of-code-solutions/pkg/utils"
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

type SpringRecord struct {
	line        string
	groupCounts []int
}

func part1(inputData string) int {
	lines := strings.Split(inputData, "\n")
	reDig := regexp.MustCompile(`\d+`)

	rs := util.MapSlice(lines, func(s string) SpringRecord {
		parts := strings.Split(s, " ")
		digitStrings, _ := util.MapSliceWithError(reDig.FindAllString(parts[1], -1), strconv.Atoi)
		return SpringRecord{
			line:        parts[0] + ".",
			groupCounts: digitStrings,
		}
	})

	total := 0
	for _, row := range rs {
		total += calculateArrangements(row.line, row.groupCounts)
	}

	return total
}

func part2(inputData string) int {
	lines := strings.Split(inputData, "\n")
	reDig := regexp.MustCompile(`\d+`)

	rs := util.MapSlice(lines, func(s string) SpringRecord {
		parts := strings.Split(s, " ")
		digitsRepeated := repeatStringSeperated(parts[1], ",", 5)
		groupCounts, _ := util.MapSliceWithError(reDig.FindAllString(digitsRepeated, -1), strconv.Atoi)
		return SpringRecord{
			line:        repeatStringSeperated(parts[0], "?", 5) + ".",
			groupCounts: groupCounts,
		}
	})

	total := 0

	for _, row := range rs {
		total += calculateArrangements(row.line, row.groupCounts)
	}

	return total
}

func repeatStringSeperated(s string, seperator string, num int) string {
	sSlice := make([]string, num)
	for i := range sSlice {
		sSlice[i] = s
	}
	return strings.Join(sSlice, seperator)
}

var cache = make(map[string]int)

func calculateArrangements(line string, groupCounts []int) (total int) {
	cacheKey := line + fmt.Sprint(groupCounts)
	if v, ok := cache[cacheKey]; ok {
		return v
	}
	calcArrange := func(line string, groupCounts []int) (total int) {
		if line == "" {
			if len(groupCounts) > 0 {
				return 0
			}
			return 1
		}

		if len(groupCounts) == 0 {
			if strings.Contains(line, "#") {
				return 0
			}
		}

		switch line[0] {
		case '.':
			total += calculateArrangements(line[1:], groupCounts)
		case '#':
			nextGroup := groupCounts[0]
			if nextGroup >= len(line) {
				return 0
			}
			firstGroup := line[:nextGroup]
			firstAfter := line[nextGroup]
			if !strings.Contains(firstGroup, ".") && firstAfter != '#' {
				newString := line[len(firstGroup)+1:]
				total += calculateArrangements(newString, groupCounts[1:])
			} else {
				return 0
			}
		case '?':
			withDot := "." + line[1:]
			total += calculateArrangements(withDot, groupCounts)
			if len(groupCounts) > 0 {
				withHash := "#" + line[1:]
				total += calculateArrangements(withHash, groupCounts)
			}
		}
		return
	}
	res := calcArrange(line, groupCounts)
	cache[cacheKey] = res
	return res
}
