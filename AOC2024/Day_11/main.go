package main

import (
	"fmt"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	blinks := strings.Fields(lines[0])
	count := 0
	for range 25 {
		new_blinks := make([]string, 0)
		for _, line := range blinks {
			if line == "0" {
				new_blinks = append(new_blinks, "1")
			} else if len(line)%2 == 0 {
				new_val, _ := strconv.Atoi(line[:len(line)/2])
				new_blinks = append(new_blinks, strconv.Itoa(new_val))
				new_val_two, _ := strconv.Atoi(line[len(line)/2:])
				new_blinks = append(new_blinks, strconv.Itoa(new_val_two))
			} else {
				val, _ := strconv.Atoi(line)
				new_val := val * 2024
				new_blinks = append(new_blinks, strconv.Itoa(new_val))
			}
		}
		blinks = new_blinks
		count = len(blinks)
	}

	return strconv.Itoa(count)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	blinks := strings.Fields(lines[0])
	numbers := make(map[string]int, 0)
	for _, val := range blinks {
		numbers[val] += 1
	}
	for range 75 {
		new_numbers := make(map[string]int, 0)
		for line := range numbers {
			if line == "0" {
				new_numbers["1"] += numbers["0"]
			} else if len(line)%2 == 0 {
				new_val, _ := strconv.Atoi(line[:len(line)/2])
				new_numbers[strconv.Itoa(new_val)] += numbers[line]
				new_val_two, _ := strconv.Atoi(line[len(line)/2:])
				new_numbers[strconv.Itoa(new_val_two)] += numbers[line]
			} else {
				val, _ := strconv.Atoi(line)
				new_val := val * 2024
				new_numbers[strconv.Itoa(new_val)] += numbers[line]
			}
		}
		numbers = new_numbers
	}
	count := 0
	for _, val := range numbers {
		count += val
	}
	return strconv.Itoa(count)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
