package main

import (
	"fmt"
	"math"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	pointer := 50
	count := 0
	for _, line := range lines {
		direction := string(line[0])
		number, _ := strconv.Atoi(line[1:])
		if direction == "R" {
			pointer = pointer + number
		} else {
			pointer = pointer - number
		}
		pointer = pointer % 100
		if pointer == 0 {
			count += 1
		}
	}
	return strconv.Itoa(count)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	pointer := 50
	count := 0
	for _, line := range lines {
		direction := string(line[0])
		number, _ := strconv.Atoi(line[1:])
		beforebefore := pointer
		if direction == "R" {
			pointer = pointer + number
		} else {
			pointer = pointer - number
		}
		before := pointer
		if pointer < 0 {
			count += int(math.Abs(float64(pointer))) / 100
			if beforebefore != 0 {
				count += 1
			}
			pointer = (100 - int(math.Abs(float64(pointer)))%100) % 100
		} else if pointer >= 100 {
			count += int(math.Abs(float64(pointer))) / 100
			pointer = pointer % 100
		} else if pointer == 0 {
			count += 1
		}
		fmt.Println(line, before, pointer, count)
	}
	return strconv.Itoa(count)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
