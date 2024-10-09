package main

import (
	"fmt"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := ""
	for _, line := range lines {
		fmt.Println(line)
	}
	return result
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := ""
	for _, line := range lines {
		fmt.Println(line)
	}
	return result
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
