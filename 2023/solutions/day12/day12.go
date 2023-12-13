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
	condition string
	groups    []int
}

func part1(inputData string) int {
	lines := strings.Split(inputData, "\n")
	reDig := regexp.MustCompile(`\d+`)

	rs := util.MapSlice(lines, func(s string) SpringRecord {
		parts := strings.Split(s, " ")
		digitStrings, _ := util.MapSliceWithError(reDig.FindAllString(parts[1], -1), strconv.Atoi)
		return SpringRecord{
			condition: parts[0],
			groups:    digitStrings,
		}
	})

	total := 0
	for _, row := range rs {
		test := strings.FieldsFunc(row.condition, func(r rune) bool { return r == '.' })
		fmt.Println(len(test))

		// total += calculateArrangements(row.condition, row.groups)
	}

	return total
}

// var reQuestion = regexp.MustCompile(`\?`)

func calculateArrangements(groups []string, groupCounts []int) (total int) {
	// unusedRe := reQuestion.FindAllStringIndex(string(sr.condition), -1)
	// unused := util.MapSlice(unusedRe, func(el []int) int { return el[0] })
	if len(groups) == 0 && len(groupCounts) == 0 {
		return 1
	} else if len(groups) == 0 || len(groupCounts) == 0 {
		return 0
	}

	// firstGroup := groups[0]
	// switch firstGroup[0] {
	// case '.':

	// 	total += calculateArrangements(line[1:], groupCounts)
	// case '#':
	// 	groupLen := groupCounts[0]
	// 	if len(line) < groupLen {
	// 		return 0
	// 	}
	// 	firstGroup := line[:groupCounts[0]]
	// 	firstAfter := line[groupCounts[0]]
	// 	if !strings.Contains(firstGroup, ".") && firstAfter != '#' {
	// 		newString := line[len(firstGroup)+1:]
	// 		total += calculateArrangements(newString, groupCounts[1:])
	// 	}
	// case '?':
	// 	withHash := "#" + line[1:]
	// 	withDot := "." + line[1:]
	// 	total += calculateArrangements(withHash, groupCounts)
	// 	total += calculateArrangements(withDot, groupCounts)
	// }
	return
}

func part2(inputData string) int {
	return 0
}
