// Solution for day11 of the Advent of Code Challenge 2024
package day11

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

func parseInput(input string) map[int]int {
	stones := util.MapSlice(strings.Fields(input), util.MustAtoi)
	stoneCount := make(map[int]int)
	for _, stone := range stones {
		stoneCount[stone]++
	}
	return stoneCount
}

func part1(inputData string) int {
	stones := parseInput(inputData)
	for i := 0; i < 25; i++ {
		stones = blinkStones(stones)
	}
	return sumStonesCount(stones)
}

func part2(inputData string) int {
	stones := parseInput(inputData)
	for i := 0; i < 75; i++ {
		stones = blinkStones(stones)
	}
	return sumStonesCount(stones)
}

func sumStonesCount(stoneCount map[int]int) (sum int) {
	for _, count := range stoneCount {
		sum += count
	}
	return sum
}

func blinkStones(stoneCount map[int]int) map[int]int {
	updatedStoneCount := make(map[int]int)
	for stone, count := range stoneCount {
		stones := blinkStone(stone)
		for _, s := range stones {
			updatedStoneCount[s] += count
		}
	}
	return updatedStoneCount
}

func blinkStone(stone int) []int {
	strStone := strconv.Itoa(stone)
	if stone == 0 {
		return []int{1}
	} else if len(strStone)%2 == 0 {
		s1 := strStone[:len(strStone)/2]
		s2 := strStone[len(strStone)/2:]
		return []int{util.MustAtoi(s1), util.MustAtoi(s2)}
	} else {
		return []int{stone * 2024}
	}
}
