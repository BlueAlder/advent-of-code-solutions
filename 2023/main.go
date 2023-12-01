package main

import (
	"flag"
	"fmt"

	"github.com/BlueAlder/advent-of-code-solutions/solutions"
	util "github.com/BlueAlder/advent-of-code-solutions/utils"
)

var dayNumber int

func main() {
	fmt.Println("Advent of Code!")

	flag.IntVar(&dayNumber, "dayNumber", -1, "the challenge number that you would like to run")
	flag.IntVar(&dayNumber, "d", -1, "the challenge number that you would like to run")

	flag.Parse()

	if dayNumber < 1 || dayNumber > 25 {
		util.LogFatal("Invalid day number, please set a value between 1-25\n")
	}

	solutions.Run(dayNumber)
}
