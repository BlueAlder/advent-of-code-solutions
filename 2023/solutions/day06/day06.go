// Solution for day06 of the Advent of Code Challenge 2023
package day06

import (
	_ "embed"
	"math"
	"regexp"
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

type Race struct {
	time     int
	distance int
}

func part1() int {
	reNum := regexp.MustCompile(`\d+`)
	digStr := reNum.FindAllString(input, -1)
	digs, _ := util.MapSliceWithError(digStr, strconv.Atoi)

	var races []Race
	for i := 0; i < len(digs)/2; i++ {
		races = append(races, Race{digs[i], digs[(len(digs)/2)+i]})
	}

	total := 1
	for _, race := range races {
		total *= getWinnable(race)
	}
	return total
}

func part2() int {
	reNum := regexp.MustCompile(`\d+`)
	digStr := reNum.FindAllString(input, -1)
	time, _ := strconv.Atoi(strings.Join(digStr[:len(digStr)/2], ""))
	distance, _ := strconv.Atoi(strings.Join(digStr[len(digStr)/2:], ""))
	race := Race{time: time, distance: distance}
	return getWinnable(race)
}

func getWinnable(r Race) int {
	b := float64(r.time)
	c := float64(r.distance)
	disc := math.Sqrt((b * b) - (4 * c))
	// Adding small values here in case we get an exact root since
	// we need to BEAT the record distance
	min := math.Ceil((b-disc)/2 + 0.00001)
	max := math.Floor((b+disc)/2 - 0.00001)
	// +1 here to count the end bound
	return int(max - min + 1)
}
