// Solution for day09 of the Advent of Code Challenge 2024
package day09

import (
	_ "embed"
	"fmt"
	"slices"
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

type block struct {
	value  int
	length int
	empty  bool
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

func checkSump1(blocks []block) int {
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

	return checkSump1(blocks)
}

// Just to visualise what we are outputting
func printBlocks(blocks []block) {
	for _, b := range blocks {
		if b.empty {
			fmt.Print(strings.Repeat(".", b.length))
		} else {
			fmt.Print(strings.Repeat(strconv.Itoa(b.value), b.length))
		}
	}
	fmt.Println()
}

func part2(inputData string) int {
	var blocks []block
	for i, char := range inputData {
		count := util.MustAtoi(string(char))
		if i%2 == 0 {
			blocks = append(blocks, block{value: i / 2, length: count, empty: false})
		} else {
			blocks = append(blocks, block{value: -1, length: count, empty: true})
		}
	}
	for blockIndex := len(blocks) - 1; blockIndex > 0; blockIndex-- {
		if blocks[blockIndex].empty {
			continue
		}

		for freeIndex := 0; freeIndex < blockIndex; freeIndex++ {
			if !blocks[freeIndex].empty {
				continue
			}
			if blocks[freeIndex].length >= blocks[blockIndex].length {
				remainingSpace := blocks[freeIndex].length - blocks[blockIndex].length
				blocks[freeIndex] = blocks[blockIndex]
				blocks[blockIndex] = block{value: -1, length: blocks[blockIndex].length, empty: true}
				if remainingSpace > 0 {
					blocks = slices.Insert(blocks, freeIndex+1, block{value: -1, length: remainingSpace, empty: true})
				}
				break

			}
		}
	}
	return checkSumP2(blocks)
}

func checkSumP2(blocks []block) int {
	sum := 0
	i := 0
	for _, block := range blocks {
		if !block.empty {
			for j := 0; j < block.length; j++ {
				sum += (block.value * i)
				i++
			}
		} else {
			i += block.length
		}
	}
	return sum
}
