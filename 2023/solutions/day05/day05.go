// Solution for day05 of the Advent of Code Challenge 2023
package day05

import (
	_ "embed"
	"math"
	"strconv"
	"strings"

	util "github.com/BlueAlder/advent-of-code-solutions/utils"
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

type Mapping struct {
	source Interval
	dest   Interval
}

type Interval struct {
	start int
	dist  int
}

func part1(inputData string) int {
	parts := strings.Split(inputData, "\n\n")
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

func part2(inputData string) int {
	parts := strings.Split(inputData, "\n\n")
	seeds, _ := util.MapSliceWithError(strings.Split(parts[0], " ")[1:], strconv.Atoi)

	mappingSet := util.MapSlice(parts[1:], convertTextMap)

	min := math.MaxInt
	for i := 0; i < len(seeds); i += 2 {
		r := Interval{seeds[i], seeds[i+1]}
		ranges := []Interval{r}

		for _, mappings := range mappingSet {
			var mapped []Interval
			for _, source := range ranges {
				mapped = append(mapped, mapRangeToRanges(source, mappings)...)
			}
			ranges = mapped
		}
		for _, i := range ranges {
			if i.start < min {
				min = i.start
			}
		}
	}

	return min
}

func mapRangeToRanges(i Interval, ms []Mapping) (mapped []Interval) {
	for _, m := range ms {
		if !m.source.containsValue(i.start) {

			if m.source.containsValue(i.getEndValue()) {
				// Left Split [1]
				left := Interval{i.start, m.source.start - i.start}
				inside := Interval{m.getMappedValue(m.source.start), i.getEndValue() - m.source.start + 1}
				mapped = append(mapped, inside)
				mapped = append(mapped, mapRangeToRanges(left, ms)...)
				return

			} else if m.source.start > i.getEndValue() {
				// Full Left [2]
				continue
			} else if m.source.start > i.start {
				// Overlaps whole interval [3]
				left := Interval{i.start, m.source.start - i.start}
				middle := Interval{m.getMappedValue(m.source.start), m.source.dist}
				right := Interval{m.source.getEndValue() + 1, i.getEndValue() - m.source.getEndValue() + 1}

				mapped = append(mapped, middle)
				mapped = append(mapped, mapRangeToRanges(left, ms)...)
				mapped = append(mapped, mapRangeToRanges(right, ms)...)
			} else {
				// Full Right [5] nop
				continue
			}

		} else {
			if m.source.containsValue(i.getEndValue()) {
				// Middle of mapping [4]
				mapped = append(mapped, Interval{m.getMappedValue(i.start), i.dist})
				return
			} else {
				// Right Split [6]
				right := Interval{m.source.getEndValue() + 1, i.getEndValue() - m.source.getEndValue()}
				inside := Interval{m.getMappedValue(i.start), m.source.getEndValue() - i.start + 1}
				mapped = append(mapped, inside)
				mapped = append(mapped, mapRangeToRanges(right, ms)...)
				return
			}
		}
	}
	return []Interval{i}
}

func (m Mapping) getMappedValue(val int) int {
	if val >= m.source.start && val < m.source.start+m.source.dist {
		return (val - m.source.start) + m.dest.start
	}
	return -1
}

func (i Interval) getEndValue() int {
	return i.start + i.dist - 1
}

func (i Interval) containsValue(val int) bool {
	return val >= i.start && val <= i.getEndValue()
}

func convertTextMap(text string) []Mapping {
	lines := strings.Split(text, "\n")[1:]

	return util.MapSlice(lines, func(line string) Mapping {
		numbers, _ := util.MapSliceWithError(strings.Split(line, " "), strconv.Atoi)
		return Mapping{source: Interval{numbers[1], numbers[2]}, dest: Interval{numbers[0], numbers[2]}}
	})
}
