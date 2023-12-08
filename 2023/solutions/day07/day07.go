// Solution for day07 of the Advent of Code Challenge 2023
package day07

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"

	util "github.com/BlueAlder/advent-of-code-solutions/utils"
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

type Hand string

type Play struct {
	hand Hand
	bid  int
}

func part1(inputData string) int {
	var cardValues = []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}

	lines := strings.Split(inputData, "\n")

	plays := util.MapSlice(lines, func(line string) Play {
		parts := strings.Split(line, " ")
		bid, _ := strconv.Atoi(parts[1])
		return Play{hand: Hand(parts[0]), bid: bid}
	})

	sort.Slice(plays, func(i, j int) bool {
		return plays[j].hand.beats(plays[i].hand, cardValues, false)
	})

	total := 0
	for idx, play := range plays {
		total += play.bid * (idx + 1)
	}

	return total
}

func (h1 Hand) beats(h2 Hand, cardValues []rune, joker bool) bool {
	var val1, val2 int
	if joker {
		val1 = h1.getHandValueJoker()
		val2 = h2.getHandValueJoker()
	} else {
		val1 = h1.getHandValue()
		val2 = h2.getHandValue()
	}

	if val1 == val2 {
		for i := 0; i < len(h1); i++ {
			cVal1 := getCardValue(rune(h1[i]), cardValues)
			cVal2 := getCardValue(rune(h2[i]), cardValues)
			if cVal1 != cVal2 {
				return cVal1 > cVal2
			}
		}
		return false
	}
	return val1 > val2
}

func getCardValue(r rune, cardValues []rune) int {
	for idx, v := range cardValues {
		if v == r {
			return idx
		}
	}
	return -1

}

// five of a kind : 6
// Four of a kind: 5
// full house : 4
// three of a kind : 3
// two pair : 2
// one pair : 1
// high card : 0
func (h Hand) getHandValue() int {
	vals := make(map[rune]int)
	for _, card := range h {
		vals[card]++
	}

	if len(vals) == 1 {
		// Five of a kind
		return 6
	} else if len(vals) == 2 {
		if vals[rune(h[0])] == 4 || vals[rune(h[0])] == 1 {
			// 4 of a kind
			return 5
		}
		// Full house
		return 4
	} else if len(vals) == 3 {
		for _, card := range h {
			if vals[card] == 3 {
				// three of a kind
				return 3
			} else if vals[card] == 2 {
				// Two pair
				return 2
			}
		}
	} else if len(vals) == 4 {
		// one pair
		return 1
	}
	return 0
}

// five of a kind : 6
// Four of a kind: 5
// full house : 4
// three of a kind : 3
// two pair : 2
// one pair : 1
// high card : 0
func (h Hand) getHandValueJoker() int {

	count := make(map[rune]int)
	for _, card := range h {
		count[card]++
	}

	jokers := count['J']

	if len(count) == 1 {
		// Five of a kind
		return 6
	} else if len(count) == 2 {

		if count[rune(h[0])] == 4 || count[rune(h[0])] == 1 {
			// 4 of a kind
			if jokers > 0 {
				return 6
			}
			return 5
		}
		if jokers > 0 {
			return 6
		}
		// Full house
		return 4

	} else if len(count) == 3 {
		for _, card := range h {
			if count[card] == 3 {
				if jokers > 0 {
					return 5
				}
				// three of a kind
				return 3
			} else if count[card] == 2 {
				if jokers == 1 {
					// Full house
					return 4
				} else if jokers == 2 {
					return 5
				}
				// Two pair
				return 2
			}
		}
	} else if len(count) == 4 {
		// one pair
		if jokers > 0 {
			return 3
		}
		return 1
	}

	// 5 diff cards
	if jokers > 0 {
		return 1
	}
	return 0
}

func part2(inputData string) int {

	var cardValues = []rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

	lines := strings.Split(inputData, "\n")

	plays := util.MapSlice(lines, func(line string) Play {
		parts := strings.Split(line, " ")
		bid, _ := strconv.Atoi(parts[1])
		return Play{hand: Hand(parts[0]), bid: bid}
	})

	sort.Slice(plays, func(i, j int) bool {
		return plays[j].hand.beats(plays[i].hand, cardValues, true)
	})

	total := 0
	for idx, play := range plays {
		total += play.bid * (idx + 1)
	}

	return total
}
