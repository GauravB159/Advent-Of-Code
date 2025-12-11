package main

import (
	"fmt"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func traverse(node string, target string, graph map[string][]string) int {
	if node == target {
		return 1
	}
	result := 0
	for _, path := range graph[node] {
		result += traverse(path, target, graph)
	}
	return result
}

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	graph := make(map[string][]string)
	for _, line := range lines {
		splits := strings.Split(line, " ")
		input := splits[0]
		graph[input[:len(input)-1]] = splits[1:]
	}
	result := traverse("you", "out", graph)
	return strconv.Itoa(result)
}

func two_traverse(node string, level int, target string, visited map[string]int, cache map[string][]int, graph map[string][]string) []int {
	if node == target {
		cache[node] = []int{1, 0, 0, 0}
		return []int{1, 0, 0, 0}
	}
	if value, exists := cache[node]; exists {
		return value
	}
	result := []int{0, 0, 0, 0}
	visited[node] = level
	for _, path := range graph[node] {
		temp_result := two_traverse(path, level+1, target, visited, cache, graph)
		switch path {
		case "fft":
			result[1] += temp_result[0]
			result[3] += temp_result[2]
		case "dac":
			result[2] += temp_result[0]
			result[3] += temp_result[1]
		default:
			result[0] += temp_result[0]
			result[1] += temp_result[1]
			result[2] += temp_result[2]
			result[3] += temp_result[3]
		}
	}
	delete(visited, node)
	cache[node] = result
	return result
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	graph := make(map[string][]string)
	for _, line := range lines {
		splits := strings.Split(line, " ")
		input := splits[0]
		graph[input[:len(input)-1]] = splits[1:]
	}
	visited := make(map[string]int)
	cache := make(map[string][]int)
	result := two_traverse("svr", 0, "out", visited, cache, graph)
	return strconv.Itoa(result[3])
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
