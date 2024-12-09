// Solution for day09 of the Advent of Code Challenge 2024
package day09

import (
	_ "embed"

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

type block struct {
	value int
	empty bool
}

func parseInput(input string) []block {
	var blocks []block
	for i, char := range input {
		count := util.MustAtoi(string(char))
		var toAdd block
		if i%2 == 0 {
			toAdd = block{value: i / 2, empty: false}
		} else {
			toAdd = block{value: -1, empty: true}
		}
		for cnt := 0; cnt < count; cnt++ {
			blocks = append(blocks, toAdd)
		}
	}
	return blocks
}

func checkSum(blocks []block) int {
	sum := 0
	for i := 0; i < len(blocks); i++ {
		if !blocks[i].empty {
			sum += (blocks[i].value * i)
		}
	}
	return sum
}

func part1(inputData string) int {
	blocks := parseInput(inputData)

	p1 := 0
	p2 := len(blocks) - 1

	for p1 < p2 {
		if !blocks[p1].empty {
			p1++
			continue
		}

		if blocks[p2].empty {
			p2--
			continue
		}
		blocks[p1], blocks[p2] = blocks[p2], blocks[p1]
		p1++
		p2--
	}

	return checkSum(blocks)
}

func part2(inputData string) int {
	return 0
}
