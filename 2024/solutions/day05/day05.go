// Solution for day05 of the Advent of Code Challenge 2024
package day05

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

func parseInput(input string) ([][]int, [][]int) {
	parts := strings.Split(input, "\n\n")
	ordering := parts[0]
	updates := parts[1]

	orderingLines := strings.Split(ordering, "\n")
	orderingString := util.MapSlice(orderingLines, func(s string) []string {
		return strings.Split(s, "|")
	})
	orderPairs := util.MapSlice(orderingString, func(s []string) []int {
		return util.MapSlice(s, func(s string) int {
			return util.MustAtoi(s)
		})
	})

	updateLines := strings.Split(updates, "\n")
	updateString := util.MapSlice(updateLines, func(s string) []string {
		return strings.Split(s, ",")
	})
	updatePages := util.MapSlice(updateString, func(s []string) []int {
		return util.MapSlice(s, func(s string) int {
			return util.MustAtoi(s)
		})
	})

	return orderPairs, updatePages
}

type OrderMap map[int]Ordering

type Ordering struct {
	before []int
	after  []int
}

func part1(inputData string) int {
	ordering, updates := parseInput(inputData)

	orderMap := make(OrderMap)
	for _, order := range ordering {
		if _, ok := orderMap[order[1]]; ok {
			om := orderMap[order[1]]
			om.before = append(orderMap[order[1]].before, order[0])
			orderMap[order[1]] = om
		} else {
			orderMap[order[1]] = Ordering{before: []int{order[0]}, after: []int{}}
		}
		if _, ok := orderMap[order[0]]; ok {
			om := orderMap[order[0]]
			om.after = append(orderMap[order[0]].after, order[1])
			orderMap[order[0]] = om
		} else {
			orderMap[order[0]] = Ordering{before: []int{}, after: []int{order[1]}}
		}
	}

	total := 0
	for _, update := range updates {
		if orderMap.isValidUpdate(update) {
			total += update[len(update)/2]
		}
	}

	return total
}

func part2(inputData string) int {
	return 0
}

func (o OrderMap) isValidUpdate(update []int) bool {
	for i := range update {
		before := update[:i]
		after := update[i+1:]
		if ordering, ok := o[update[i]]; ok {
			if util.DoIntersect(before, ordering.after) || util.DoIntersect(after, ordering.before) {
				return false
			}
		}
	}
	return true
}
