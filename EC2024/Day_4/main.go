package main

import (
	"fmt"
	"math"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func part_one(filename string) string {
	lines := aocutils.Readfile(filename)
	minimum := 1000000
	for _, line := range lines {
		value, _ := strconv.Atoi(line)
		if value < minimum {
			minimum = value
		}
	}
	answer := 0
	for _, line := range lines {
		value, _ := strconv.Atoi(line)
		answer += (value - minimum)
	}
	return strconv.Itoa(answer)
}

func part_two(filename string) string {
	lines := aocutils.Readfile(filename)
	minimum := 1000000
	for _, line := range lines {
		value, _ := strconv.Atoi(line)
		if value < minimum {
			minimum = value
		}
	}
	answer := 0
	for _, line := range lines {
		value, _ := strconv.Atoi(line)
		answer += (value - minimum)
	}
	return strconv.Itoa(answer)
}

func part_three(filename string) string {
	lines := aocutils.Readfile(filename)
	minimum := 10000000000
	for _, line_one := range lines {
		check, _ := strconv.Atoi(line_one)
		answer := 0
		for _, line_two := range lines {
			value, _ := strconv.Atoi(line_two)
			answer += int(math.Abs(float64(value) - float64(check)))
		}
		if answer < minimum {
			minimum = answer
		}
	}
	return strconv.Itoa(minimum)
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
