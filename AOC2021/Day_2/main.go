package main

import (
	"fmt"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	vertical := 0
	horizontal := 0
	for _, line := range lines {
		fields := strings.Split(line, " ")
		direction := fields[0]
		value, _ := strconv.Atoi(fields[1])
		if direction == "forward" {
			horizontal += value
		} else if direction == "down" {
			vertical += value
		} else {
			vertical -= value
		}
	}
	return strconv.Itoa(vertical * horizontal)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	aim := 0
	vertical := 0
	horizontal := 0
	for _, line := range lines {
		fields := strings.Split(line, " ")
		direction := fields[0]
		value, _ := strconv.Atoi(fields[1])
		if direction == "forward" {
			horizontal += value
			vertical += value * aim
		} else if direction == "down" {
			aim += value
		} else {
			aim -= value
		}
	}
	return strconv.Itoa(vertical * horizontal)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
