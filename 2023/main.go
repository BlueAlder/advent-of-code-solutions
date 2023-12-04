package main

import (
	"flag"
	"slices"

	"github.com/BlueAlder/advent-of-code-solutions/solutions"
	util "github.com/BlueAlder/advent-of-code-solutions/utils"
)

var dayNumber int
var partNumber int

var parts = []int{1, 2}

func main() {
	util.LogGood("Advent of Code 2023!")

	flag.IntVar(&dayNumber, "dayNumber", -1, "the challenge number that you would like to run")
	flag.IntVar(&dayNumber, "d", -1, "the challenge number that you would like to run")

	flag.IntVar(&partNumber, "part", 0, "the part you would like to run")
	flag.IntVar(&partNumber, "p", 0, "the part you would like to run")

	flag.Parse()

	if dayNumber < 1 || dayNumber > 25 {
		util.LogFatal("Invalid day number, please set a value between 1-25\n")
	}

	if slices.Contains(parts, partNumber) {
		solutions.Run(dayNumber, partNumber)
	} else {
		solutions.Run(dayNumber, 1)
		solutions.Run(dayNumber, 2)
	}
}
