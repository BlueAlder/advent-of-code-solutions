package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var points = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
	"A": 1,
	"B": 2,
	"C": 3,
}

func solve(inputFileName string, part int) int {
	file, err := os.Open(inputFileName)

	if err != nil {
		fmt.Printf("Unable to open file %s\n", inputFileName)
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := 0
	for scanner.Scan() {
		enemy, play := func() (string, string) {
			x := strings.Split(scanner.Text(), " ")
			return x[0], x[1]
		}()

		if part == 1 {
			total += part1(enemy, play)
		} else if part == 2 {
			total += part2(enemy, play)
		}
	}
	return total
}

func part1(enemy string, play string) int {
	total := 0
	sum := points[enemy] + points[play]
	if points[enemy] == points[play] {
		total += 3
	} else if (sum == 3 && play == "Y") || (sum == 4 && play == "X") || (sum == 5 && play == "Z") {
		total += 6
	}
	total += points[play]
	return total
}

func part2(play string, result string) int {
	if result == "X" {
		return ((((points[play] - 2) % 3) + 3) % 3) + 1
	} else if result == "Y" {
		return points[play] + 3
	} else if result == "Z" {
		return ((((points[play] + 3) % 3) + 3) % 3) + 1 + 6
	}
	return 0
}

func main() {
	fmt.Println("Part 1:", solve("input.txt", 1))
	fmt.Println("Part 2:", solve("input.txt", 2))
}
