// Solution for day12 of the Advent of Code Challenge 2023
package day12

import (
	_ "embed"
	"reflect"
	"regexp"
	"strconv"
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

type Condition string

type SpringRecord struct {
	condition Condition
	groups    []int
}

func part1(inputData string) int {
	lines := strings.Split(inputData, "\n")
	reDig := regexp.MustCompile(`\d+`)

	rs := util.MapSlice(lines, func(s string) SpringRecord {
		parts := strings.Split(s, " ")
		digitStrings, _ := util.MapSliceWithError(reDig.FindAllString(parts[1], -1), strconv.Atoi)
		return SpringRecord{
			condition: Condition(parts[0]),
			groups:    digitStrings,
		}
	})

	total := 0
	for _, row := range rs {
		total += row.calculateArrangements()
	}

	return total
}

var reQuestion = regexp.MustCompile(`\?`)

func (sr SpringRecord) calculateArrangements() (total int) {
	unusedRe := reQuestion.FindAllStringIndex(string(sr.condition), -1)
	unused := util.MapSlice(unusedRe, func(el []int) int { return el[0] })

	currentGrouping := sr.condition.getGroupings()
	toPlace := util.SumIntSlice(sr.groups) - util.SumIntSlice(currentGrouping)

	combinations := combin.Combinations(len(unused), toPlace)
	for _, idxs := range combinations {
		newString := sr.condition
		for _, idx := range idxs {
			newString = newString[:unused[idx]] + "#" + newString[unused[idx]+1:]
		}
		newGrouping := newString.getGroupings()
		if reflect.DeepEqual(newGrouping, sr.groups) {
			total++
		}
	}
	return total
}

func (sr SpringRecord) isPossible() bool {

	groups := sr.condition.getGroupings()
	return util.MaxIntSlice(sr.groups) >= util.MaxIntSlice(groups)
}

var reHashes = regexp.MustCompile("#+")

func (c Condition) getGroupings() []int {
	return util.MapSlice(reHashes.FindAllString(string(c), -1), func(el string) int {
		return len(el)
	})
}

func part2(inputData string) int {
	return 0
}
