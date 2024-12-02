// Solution for day15 of the Advent of Code Challenge 2023
package day15

import (
	_ "embed"
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

func part1(inputData string) int {
	steps := strings.Split(inputData, ",")
	total := 0
	for _, step := range steps {
		total += hash(step)
	}

	return total
}

type Lens struct {
	label  string
	length int
}

func part2(inputData string) int {
	steps := strings.Split(inputData, ",")
	boxes := make([][]Lens, 256)
	// for i := 0; i < 256; i++ {
	// boxes[i] = make([]In)
	// }

	for _, step := range steps {
		if step[len(step)-1] == '-' {
			label := step[:len(step)-1]
			boxNum := hash(label)
			idxToDelete := slices.IndexFunc(boxes[boxNum], func(l Lens) bool { return l.label == label })
			if idxToDelete != -1 {
				boxes[boxNum] = append(boxes[boxNum][:idxToDelete], boxes[boxNum][idxToDelete+1:]...)
			}
		} else {
			data := strings.Split(step, "=")
			length, _ := strconv.Atoi(data[1])
			newLens := Lens{
				label:  data[0],
				length: length,
			}

			boxNum := hash(newLens.label)
			idxToReplace := slices.IndexFunc(boxes[boxNum], func(l Lens) bool { return l.label == newLens.label })
			if idxToReplace != -1 {
				boxes[boxNum][idxToReplace] = newLens
			} else {
				boxes[boxNum] = append(boxes[boxNum], newLens)
			}

		}

	}

	power := 0
	for i := range boxes {
		for slot, lens := range boxes[i] {
			power += (i + 1) * (slot + 1) * (lens.length)
		}
	}
	return power
}

func hash(s string) int {
	hash := 0
	for _, c := range s {
		hash += int(c)
		hash *= 17
		hash = hash % 256
	}
	return hash
}
