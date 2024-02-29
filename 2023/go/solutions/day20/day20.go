// Solution for day20 of the Advent of Code Challenge 2023
package day20

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"

	util "github.com/BlueAlder/advent-of-code-solutions/pkg/utils"
)

//go:embed input.txt
var input string

// Flip Flop %
//  either on or off
// INIT off
// RECEIVES HIGH : nothing
// RECEIVES LOW: STATUS: OFF -> STATUS: ON  SEND: HIGH
// RECEIVES LOW: STATUS: ON -> STATUS: OFF  SEND: LOW

// CONJUCTION &
// remember last pulse from all connected modules
// INIT ALL CONNECTED = LOW
// RECEIVES: UPDATEMEMORY FROM INPUT
// 		IF ALL PREVIOUS HIGH : SEND LOW
// 		ELSE: SEND HIGH

// BROADCAST
// RELAYS SAME PULSE TO ALL DEST

// BUTTON = LOW PULSE

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
	configs := buildConfigStructure(inputData)
	high := 0
	low := 0
	for i := 0; i < 1000; i++ {
		queue := make([]Pulse, 0)
		p1 := Pulse{
			src:  "button",
			dest: "broadcaster",
			high: false,
		}
		queue = append(queue, p1)

		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			if p.high {
				high++
			} else {
				low++
			}

			m, ok := configs[p.dest]
			if ok {
				queue = append(queue, m.receive(p)...)
			}
		}
	}

	return high * low
}

func part2(inputData string) int {
	configs := buildConfigStructure(inputData)
	buttonPresses := 0

	// Find the conjuction to rx
	var rxSrc *Conjuction
	for _, v := range configs {
		if slices.Contains(v.GetDestinations(), "rx") {
			rxSrc = v.(*Conjuction)
		}
	}

	// Find it's sources and find when they all send a high since they all cycle
	// from testing
	sources := make(map[string]int)
	for k := range rxSrc.inputLastValue {
		sources[k] = 0
	}

	for {
		buttonPresses += 1
		for buttonPresses%100000 == 0 {
			fmt.Println(buttonPresses)
		}

		queue := make([]Pulse, 0)
		p1 := Pulse{
			src:  "button",
			dest: "broadcaster",
			high: false,
		}
		queue = append(queue, p1)

		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			// Check cycle of when sources send a high and then multiply them together
			// to find the lowest common denominator
			if count, ok := sources[p.src]; ok && p.dest == rxSrc.label && p.high && count == 0 {
				sources[p.src] = buttonPresses
				total := 1
				for _, v := range sources {
					total *= v
				}
				if total != 0 {
					return total
				}
			}

			m, ok := configs[p.dest]
			if ok {
				queue = append(queue, m.receive(p)...)
			}
		}
	}
}

func buildConfigStructure(inputData string) map[string]Module {
	configsString := strings.Split(inputData, "\n")
	configs := make(map[string]Module)
	conjuctions := make([]*Conjuction, 0)

	for _, config := range configsString {
		parts := strings.Split(config, " -> ")
		destinations := strings.Split(parts[1], ", ")
		label := parts[0]
		if label[0] == '%' {
			label = label[1:]
			configs[label] = &FlipFlop{
				status: false,
				ModuleData: ModuleData{
					label:        label,
					destinations: destinations,
				},
			}

		} else if label[0] == '&' {
			label = label[1:]
			c := &Conjuction{
				inputLastValue: make(map[string]bool),
				ModuleData:     ModuleData{label: label, destinations: destinations},
			}
			configs[label] = c
			conjuctions = append(conjuctions, c)

		} else {
			configs[label] = &Broadcast{
				ModuleData: ModuleData{
					label:        label,
					destinations: destinations,
				},
			}
		}
	}

	for _, c := range conjuctions {
		for _, m := range configs {
			if slices.Contains(m.GetDestinations(), c.label) {
				c.inputLastValue[m.GetLabel()] = false
			}
		}
	}
	return configs
}
