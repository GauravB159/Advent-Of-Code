package main

import (
	"fmt"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	count := 0
	prev := -1
	for _, line := range lines {
		if num, error := strconv.Atoi(line); error == nil {
			if prev != -1 && num > prev {
				count += 1
			}
			prev = num
		}
	}
	return strconv.Itoa(count)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	count := 0
	for index, line := range lines {
		if index+3 > len(lines)-1 {
			break
		}
		num, _ := strconv.Atoi(line)
		future, _ := strconv.Atoi(lines[index+3])
		if num < future {
			count += 1
		}
	}
	return strconv.Itoa(count)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
