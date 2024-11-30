package main

import (
	"fmt"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func part_one(filename string) string {
	lines := aocutils.Readfile(filename)
	for _, line := range lines {
		key := strings.Split(line, ":")[0]
		remain := strings.Split(line, ":")[1]
		order := strings.Split(remain, ",")
		start := 10
		total := 0
		for i := 0; i < 10; i++ {
			symbol := order[i%len(order)]
			if symbol == "+" {
				start += 1
			} else if symbol == "-" {
				start -= 1
			}
			total += start
		}
		fmt.Println(key, total)
	}
	return strconv.Itoa(0)
}

func part_two(filename string) string {
	// lines := aocutils.Readfile(filename)
	return strconv.Itoa(0)
}

func part_three(filename string) string {
	// lines := aocutils.Readfile(filename)
	return strconv.Itoa(0)
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
