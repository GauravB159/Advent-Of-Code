package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func dfs_one(current string, nodes *map[string][]string, visited map[string]bool, count *int, paths *map[string]bool, path string) {
	if current == "end" {
		if (*paths)[path] {
			return
		}
		(*paths)[path] = true
		(*count) += 1
		return
	}
	if unicode.IsLower(rune(current[0])) && visited[current] {
		return
	}
	visited[current] = true
	for _, node := range (*nodes)[current] {
		dfs_one(node, nodes, visited, count, paths, path+"->"+node)
	}
	visited[current] = false
}

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	nodes := make(map[string][]string)
	count := 0
	paths := make(map[string]bool)
	for _, line := range lines {
		edge := strings.Split(line, "-")
		node := edge[0]
		connect := edge[1]
		if _, exists := nodes[connect]; !exists {
			nodes[connect] = make([]string, 0)
		}
		if _, exists := nodes[node]; !exists {
			nodes[node] = make([]string, 0)
		}
		nodes[node] = append(nodes[node], connect)
		nodes[connect] = append(nodes[connect], node)
	}
	dfs_one("start", &nodes, make(map[string]bool, 0), &count, &paths, "start")
	return strconv.Itoa(count)
}

func dfs_two(current string, nodes *map[string][]string, visited map[string]int, count *int, paths *map[string]bool, path string) {
	if current == "end" {
		if (*paths)[path] {
			return
		}
		(*paths)[path] = true
		(*count) += 1
		return
	}
	hasTwo := false
	for node := range visited {
		if visited[node] > 1 && unicode.IsLower(rune(node[0])) {
			hasTwo = true
		}
	}
	if unicode.IsLower(rune(current[0])) && visited[current] > 0 && hasTwo {
		return
	}
	visited[current] += 1
	for _, node := range (*nodes)[current] {
		if node == "start" {
			continue
		}
		dfs_two(node, nodes, visited, count, paths, path+","+node)
	}
	visited[current] -= 1
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	nodes := make(map[string][]string)
	count := 0
	paths := make(map[string]bool)
	for _, line := range lines {
		edge := strings.Split(line, "-")
		node := edge[0]
		connect := edge[1]
		if _, exists := nodes[connect]; !exists {
			nodes[connect] = make([]string, 0)
		}
		if _, exists := nodes[node]; !exists {
			nodes[node] = make([]string, 0)
		}
		nodes[node] = append(nodes[node], connect)
		nodes[connect] = append(nodes[connect], node)
	}
	dfs_two("start", &nodes, make(map[string]int, 0), &count, &paths, "start")
	fmt.Println(count)
	return ""
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
