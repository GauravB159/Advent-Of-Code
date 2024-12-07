package main

import (
	"fmt"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func dfs_one(current int, index int, target int, values []string) bool {
	if current > target {
		return false
	}
	if current == target && index == len(values) {
		return true
	}
	if index >= len(values) {
		return false
	}
	currVal, _ := strconv.Atoi(values[index])
	return (dfs_one(current+currVal, index+1, target, values) || dfs_one(current*currVal, index+1, target, values))
}

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := 0
	for _, line := range lines {
		target, _ := strconv.Atoi(strings.Split(line, ": ")[0])
		values := strings.Fields(strings.Split(line, ": ")[1])
		start, _ := strconv.Atoi(values[0])
		possible := dfs_one(start, 1, target, values)
		if possible {
			result += target
		}
	}
	return strconv.Itoa(result)
}

func concat(current int, next int) int {
	cval := strconv.Itoa(current) + strconv.Itoa(next)
	icval, _ := strconv.Atoi(cval)
	return icval
}

func dfs_two(current int, index int, target int, values []string) bool {
	if current > target {
		return false
	}
	if current == target && index == len(values) {
		return true
	}
	if index >= len(values) {
		return false
	}
	currVal, _ := strconv.Atoi(values[index])
	return dfs_two(current+currVal, index+1, target, values) || dfs_two(current*currVal, index+1, target, values) || dfs_two(concat(current, currVal), index+1, target, values)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := 0
	for _, line := range lines {
		target, _ := strconv.Atoi(strings.Split(line, ": ")[0])
		values := strings.Fields(strings.Split(line, ": ")[1])
		start, _ := strconv.Atoi(values[0])
		possible := dfs_two(start, 1, target, values)
		if possible {
			result += target
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
