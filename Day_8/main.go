package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	count := 0
	for _, line := range lines {
		digits := strings.Split(line, " | ")[1]
		for _, digit := range strings.Fields(digits) {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				count += 1
			}
		}
	}
	return strconv.Itoa(count)
}

func sortString(s string) string {
	// Convert the string to a slice of runes to handle Unicode characters
	runes := []rune(s)

	// Sort the slice of runes
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// Convert the sorted slice of runes back to a string
	return string(runes)
}

func findDifference(b string, a string) string {
	var difference string = ""
	var j int = 0
	var i int = 0

	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
		} else {
			difference += string(b[j])
			j++
		}
	}
	if j < len(b) {
		difference += string(b[j])
	}
	return difference
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := ""
	for _, line := range lines {
		decode := strings.Split(line, " | ")[0]
		// digits := strings.Split(line, " | ")[1]
		positions := make([][]string, 7)
		var one, seven, four string
		for _, digit := range strings.Fields(decode) {
			if len(digit) == 2 {
				one = sortString(digit)
			} else if len(digit) == 3 {
				seven = sortString(digit)
			} else if len(digit) == 4 {
				four = sortString(digit)
			}
		}

		positions[0] = []string{findDifference(seven, one)}
		positions[2], positions[5] = []string{string(one[0]), string(one[1])}, []string{string(one[0]), string(one[1])}
		fmt.Println(positions, four, seven, findDifference(four, seven))
	}
	return result
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
