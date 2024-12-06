// Solution for day05 of the Advent of Code Challenge 2024
package day05

import (
	_ "embed"
	"slices"
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

type OrderMap map[int][]int

func buildOrderMap(ordering [][]int) OrderMap {
	orderMap := make(OrderMap)
	for _, order := range ordering {
		if om, ok := orderMap[order[1]]; ok {
			om = append(om, order[0])
			orderMap[order[1]] = om
		} else {
			orderMap[order[1]] = []int{order[0]}
		}
	}
	return orderMap
}

func part1(inputData string) int {
	ordering, updates := parseInput(inputData)
	orderMap := buildOrderMap(ordering)

	total := 0
	for _, update := range updates {
		if orderMap.isValidUpdate(update) {
			total += update[len(update)/2]
		}
	}

	return total
}

func part2(inputData string) int {

	ordering, updates := parseInput(inputData)
	orderMap := buildOrderMap(ordering)

	total := 0
	for _, update := range updates {
		if !orderMap.isValidUpdate(update) {
			slices.SortFunc(update, func(a, b int) int {
				if slices.Contains(orderMap[a], b) {
					return -1
				}
				return 1
			})
			total += update[len(update)/2]
		}
	}
	return total
}

func (o OrderMap) isValidUpdate(update []int) bool {
	for i := range update {
		after := update[i+1:]
		if ordering, ok := o[update[i]]; ok {
			if util.DoIntersect(after, ordering) {
				return false
			}
		}
	}
	return true
}
