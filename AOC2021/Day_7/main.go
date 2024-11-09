package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	line := aocutils.Readfile(filename)[0]
	positions := strings.Split(line, ",")
	num_positions := make([]int, len(positions))
	max := 0
	for index, position := range positions {
		num_pos, _ := strconv.Atoi(position)
		num_positions[index] = num_pos
		if num_pos > max {
			max = num_pos
		}
	}
	min_sum := 1000000000.0
	for target := 0; target < max; target++ {
		sum := 0.0
		for _, value := range num_positions {
			sum += math.Abs(float64(target - value))
		}
		if sum < min_sum {
			min_sum = sum
		}
	}
	return strconv.Itoa(int(min_sum))
}

func twostar(filename string) string {
	line := aocutils.Readfile(filename)[0]
	positions := strings.Split(line, ",")
	num_positions := make([]int, len(positions))
	max := 0
	for index, position := range positions {
		num_pos, _ := strconv.Atoi(position)
		num_positions[index] = num_pos
		if num_pos > max {
			max = num_pos
		}
	}
	min_sum := 1000000000.0
	for target := 0; target < max; target++ {
		sum := 0.0
		for _, value := range num_positions {
			n := math.Abs(float64(target - value))
			sum += (n * (n + 1)) / 2
		}
		if sum < min_sum {
			min_sum = sum
		}
	}
	return strconv.Itoa(int(min_sum))
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
