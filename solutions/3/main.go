package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

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
		rs := scanner.Text()
		prio, err := getPriority(rs)
		if err != nil {
			fmt.Printf("Unable to find priority from rucksack: %s\n", rs)
			os.Exit(1)
		}
		total += prio
	}
	return total
}

func getPriority(rucksack string) (int, error) {
	c1 := rucksack[0 : len(rucksack)/2]
	c2 := rucksack[len(rucksack)/2:]

	for _, chr := range c1 {
		if strings.ContainsRune(c2, chr) {
			return getPriorityFromLetter(chr), nil
		}
	}
	return 0, errors.New("Cannot find matching priority")
}

func getPriorityFromLetter(letter rune) int {
	if unicode.IsUpper(letter) {
		return int(letter) - 64 + 26
	}
	return int(letter) - 96
}

func main() {
	fmt.Println(solve("input.txt", 1))
}
