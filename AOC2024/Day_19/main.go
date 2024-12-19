package main

import (
	"fmt"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func dfs_one(remain string, available *map[string]bool, memoize *map[string]bool) bool {
	if val, exists := (*memoize)[remain]; exists {
		return val
	}
	for i := 0; i < len(remain); i++ {
		left := remain[:i]
		right := remain[i:]
		if (*available)[left] && (*available)[right] {
			(*memoize)[remain] = true
			return true
		} else if (*available)[left] && dfs_one(right, available, memoize) {
			(*memoize)[remain] = true
			return true
		} else if (*available)[right] && dfs_one(left, available, memoize) {
			(*memoize)[remain] = true
			return true
		}
	}
	(*memoize)[remain] = false
	return false
}

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	available := strings.Split(lines[0], ", ")
	available_map := make(map[string]bool)
	for _, val := range available {
		available_map[val] = true
	}
	count := 0
	memoize := make(map[string]bool)
	for _, line := range lines[2:] {
		valid := dfs_one(line, &available_map, &memoize)
		if valid {
			count += 1
		}
	}
	return strconv.Itoa(count)
}

func dfs_two(remain string, available *map[string]bool, memoize *map[string]int, path string) int {
	if val, exist := (*memoize)[remain]; exist {
		return val
	}
	if len(remain) == 0 {
		return 1
	}
	if len(remain) == 1 {
		if (*available)[remain] {
			return 1
		} else {
			return 0
		}
	}
	if (*available)[remain] {
		(*memoize)[remain] += 1
	}
	for i := 1; i < len(remain); i++ {
		left := remain[:i]
		right := remain[i:]
		if (*available)[left] {
			val_right := dfs_two(right, available, memoize, path)
			(*memoize)[remain] += val_right
		}
	}
	return (*memoize)[remain]
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	available := strings.Split(lines[0], ", ")
	available_map := make(map[string]bool)
	for _, val := range available {
		available_map[val] = true
	}
	count := 0
	memoize := make(map[string]int)
	for _, line := range lines[2:] {
		valid := dfs_two(line, &available_map, &memoize, "")
		count += valid
	}
	return strconv.Itoa(count)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
