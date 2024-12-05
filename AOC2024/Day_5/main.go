package main

import (
	"fmt"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := 0
	first := true
	dict := make(map[string]bool)
	checkLines := make([][]string, 0)
	for _, line := range lines {
		if first {
			if line == "" {
				first = false
				continue
			}
			dict[line] = true
		} else {
			checkLines = append(checkLines, strings.Split(line, ","))
		}
	}
	for _, line := range checkLines {
		valid := checkLine(line, dict) == ""
		if valid {
			num, _ := strconv.Atoi(line[len(line)/2])
			result += num
		}
	}
	return strconv.Itoa(result)
}

func checkLine(line []string, dict map[string]bool) string {
	invalid_key := ""
	for i := 0; i < len(line); i++ {
		for j := i + 1; j < len(line); j++ {
			key := fmt.Sprintf("%s|%s", line[j], line[i])
			if _, exists := dict[key]; exists {
				invalid_key = fmt.Sprintf("%d|%d", j, i)
				return invalid_key
			}
		}
	}
	return invalid_key
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := 0
	first := true
	dict := make(map[string]bool)
	checkLines := make([][]string, 0)
	for _, line := range lines {
		if first {
			if line == "" {
				first = false
				continue
			}
			dict[line] = true
		} else {
			checkLines = append(checkLines, strings.Split(line, ","))
		}
	}
	for _, line := range checkLines {
		init_invalid_key := checkLine(line, dict)
		for invalid_key := init_invalid_key; invalid_key != ""; {
			i, _ := strconv.Atoi(strings.Split(invalid_key, "|")[0])
			j, _ := strconv.Atoi(strings.Split(invalid_key, "|")[1])
			line[i], line[j] = line[j], line[i]
			invalid_key = checkLine(line, dict)
		}
		if init_invalid_key != "" {
			num, _ := strconv.Atoi(line[len(line)/2])
			result += num
		}
	}
	return strconv.Itoa(result)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
