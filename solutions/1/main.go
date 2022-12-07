package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Unable to read file")
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var calorie_count []int
	current_count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			val, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("invalid int value")
				os.Exit(1)
			}
			current_count += val
		} else {
			calorie_count = append(calorie_count, current_count)
			current_count = 0
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calorie_count)))
	fmt.Println("Part 1", calorie_count[0])
	fmt.Println("Part 2", calorie_count[0]+calorie_count[1]+calorie_count[2])
}
