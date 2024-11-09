package main

import (
	"fmt"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func part_one(filename string) string {
	lines := aocutils.Readfile(filename)
	line := lines[0]
	mapping := map[rune]int{
		'A': 0,
		'B': 1,
		'C': 3,
	}
	sum := 0
	for _, enemy := range line {
		sum += mapping[enemy]
	}
	return strconv.Itoa(sum)
}

func part_two(filename string) string {
	lines := aocutils.Readfile(filename)
	line := lines[0]
	mapping := map[rune]int{
		'A': 0,
		'B': 1,
		'C': 3,
		'D': 5,
		'x': 0,
	}
	sum := 0
	for i := 0; i < len(line); i += 2 {
		enemy := line[i]
		next_enemy := line[i+1]
		extra := 0
		if enemy != 'x' && next_enemy != 'x' {
			extra = 2
		}
		sum += mapping[rune(enemy)] + mapping[rune(next_enemy)] + extra
	}
	return strconv.Itoa(sum)
}

func part_three(filename string) string {
	lines := aocutils.Readfile(filename)
	line := lines[0]
	mapping := map[rune]int{
		'A': 0,
		'B': 1,
		'C': 3,
		'D': 5,
		'x': 0,
	}
	sum := 0
	for i := 0; i < len(line); i += 3 {
		x_count := 0
		enemy := line[i]
		next_enemy := line[i+1]
		next_next_enemy := line[i+2]
		extra := 0
		if enemy == 'x' {
			x_count += 1
		}
		if next_enemy == 'x' {
			x_count += 1
		}
		if next_next_enemy == 'x' {
			x_count += 1
		}
		if x_count == 1 {
			extra = 2
		}
		if x_count == 0 {
			extra = 6
		}
		sum += mapping[rune(enemy)] + mapping[rune(next_enemy)] + mapping[rune(next_next_enemy)] + extra
	}
	return strconv.Itoa(sum)
}

func main() {
	aocutils.Timer("Part 1", part_one, "input_one.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("Part 2", part_two, "input_two.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("Part 3", part_three, "input_three.txt")
}
