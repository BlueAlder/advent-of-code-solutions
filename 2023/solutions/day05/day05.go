// Solution for day05 of the Advent of Code Challenge 2023
package day05

import (
	_ "embed"
	"math"
	"slices"
	"strconv"
	"strings"

	util "github.com/BlueAlder/advent-of-code-solutions/utils"
)

//go:embed input.txt
var input string

func Solve(part int) int {
	if part == 1 {
		return part1()
	} else if part == 2 {
		return part2()
	} else {
		util.LogFatal("invalid part number")
		return -1
	}
}

type Mapping struct {
	source int
	dest   int
	dist   int
}

func part1() int {
	parts := strings.Split(input, "\n\n")
	seeds, _ := util.MapSliceWithError(strings.Split(parts[0], " ")[1:], strconv.Atoi)

	mappingSet := util.MapSlice(parts[1:], convertTextMap)

	locations := make([]int, len(seeds))
	for seedIdx, seed := range seeds {
		currVal := seed

		for _, mapping := range mappingSet {
		MappingLookup:
			for _, mapEntry := range mapping {
				res := mapEntry.getMappedValue(currVal)
				if res != -1 {
					currVal = res
					break MappingLookup
				}
			}
		}
		locations[seedIdx] = currVal
	}

	min := math.MaxInt
	for _, val := range locations {
		if val < min {
			min = val
		}
	}

	return min
}

func part2() int {

	parts := strings.Split(input, "\n\n")
	seeds, _ := util.MapSliceWithError(strings.Split(parts[0], " ")[1:], strconv.Atoi)

	mappingSet := util.MapSlice(parts[1:], convertTextMap)
	slices.Reverse(mappingSet)

	seed := 0
	for {
		currVal := seed
		for _, mapping := range mappingSet {
		MappingLookup:
			for _, mapEntry := range mapping {
				res := mapEntry.getReverseMappedValue(currVal)
				if res != -1 {
					currVal = res
					break MappingLookup
				}
			}
		}
		if isInSeedRange(currVal, seeds) {
			break
		}
		seed++
	}

	return seed
}

func isInSeedRange(val int, seeds []int) bool {
	for i := 0; i < len(seeds); i += 2 {
		if val >= seeds[i] && val < seeds[i]+seeds[i+1] {
			return true
		}
	}
	return false
}

func (m *Mapping) getMappedValue(val int) int {
	if val >= m.source && val < m.source+m.dist {
		return (val - m.source) + m.dest
	}
	return -1
}

func (m *Mapping) getReverseMappedValue(val int) int {
	if val >= m.dest && val < m.dest+m.dist {
		return (val - m.dest) + m.source
	}
	return -1
}

func convertTextMap(text string) []*Mapping {
	lines := strings.Split(text, "\n")[1:]

	return util.MapSlice(lines, func(line string) *Mapping {
		numbers, _ := util.MapSliceWithError(strings.Split(line, " "), strconv.Atoi)
		return &Mapping{source: numbers[1], dest: numbers[0], dist: numbers[2]}
	})
}
