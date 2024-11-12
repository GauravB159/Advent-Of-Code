package main

import (
	"fmt"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func traverse_one(node string, path string, edges map[string]([]string), paths map[string]int, level int) {
	if node == "@" {
		paths[path] = level
	}
	for _, edge := range edges[node] {
		traverse_one(edge, path+edge, edges, paths, level+1)
	}
}

func part_one(filename string) string {
	lines := aocutils.Readfile(filename)
	edges := make(map[string]([]string))
	paths := make(map[string]int)
	for _, line := range lines {
		parent := strings.Split(line, ":")[0]
		children := strings.Split(line, ":")[1]
		edges[parent] = strings.Split(children, ",")
	}
	root := "RR"
	traverse_one(root, "", edges, paths, 0)
	counts := make(map[int]([]string), 2)
	for key, value := range paths {
		counts[value] = append(counts[value], key)
	}
	answer := ""
	for _, value := range counts {
		if len(value) == 1 {
			answer = value[0]
		}
	}
	return root + answer
}

func traverse_two(node string, path string, edges map[string]([]string), paths map[string]int, level int) {
	if node == "@" {
		paths[path] = level
	}
	for _, edge := range edges[node] {
		traverse_two(edge, path+string(edge[0]), edges, paths, level+1)
	}
}

func part_two(filename string) string {
	lines := aocutils.Readfile(filename)
	edges := make(map[string]([]string))
	paths := make(map[string]int)
	for _, line := range lines {
		parent := strings.Split(line, ":")[0]
		children := strings.Split(line, ":")[1]
		edges[parent] = strings.Split(children, ",")
	}
	root := "RR"
	traverse_two(root, "", edges, paths, 0)
	counts := make(map[int]([]string), 2)
	for key, value := range paths {
		counts[value] = append(counts[value], key)
	}
	answer := ""
	for _, value := range counts {
		if len(value) == 1 {
			answer = value[0]
		}
	}
	return string(root[0]) + answer
}

func traverse_three(node string, path string, edges map[string]([]string), paths map[string]int, level int) {
	if node == "ANT" || node == "BUG" {
		return
	}
	if node == "@" {
		paths[path] = level
	}
	for _, edge := range edges[node] {
		traverse_three(edge, path+string(edge[0]), edges, paths, level+1)
	}
}

func part_three(filename string) string {
	lines := aocutils.Readfile(filename)
	edges := make(map[string]([]string))
	paths := make(map[string]int)
	for _, line := range lines {
		parent := strings.Split(line, ":")[0]
		children := strings.Split(line, ":")[1]
		edges[parent] = strings.Split(children, ",")
	}
	root := "RR"
	traverse_three(root, "", edges, paths, 0)
	counts := make(map[int]([]string), 2)
	for key, value := range paths {
		counts[value] = append(counts[value], key)
	}
	answer := ""
	for _, value := range counts {
		if len(value) == 1 {
			answer = value[0]
		}
	}
	return string(root[0]) + answer
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
