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
	graph := make(map[string]map[string]bool)
	for _, line := range lines {
		left := strings.Split(line, "-")[0]
		right := strings.Split(line, "-")[1]
		if _, exists := graph[left]; !exists {
			graph[left] = make(map[string]bool)
		}
		if _, exists := graph[right]; !exists {
			graph[right] = make(map[string]bool)
		}
		graph[left][right] = true
		graph[right][left] = true
	}
	visited := make(map[string]bool)
	for node := range graph {
		for second_node := range graph[node] {
			for third_node := range graph[node] {
				if second_node == third_node {
					continue
				}
				if graph[second_node][third_node] {
					nodes := []string{node, second_node, third_node}
					sort.Strings(nodes)
					key := strings.Join(nodes, "-")
					if node[0] == 't' || second_node[0] == 't' || third_node[0] == 't' {
						visited[key] = true
					}
				}
			}
		}
	}
	return strconv.Itoa(len(visited))
}

func largest_connected_component(graph map[string]map[string]bool, node string, path *[]string, visited map[string]bool) {
	if visited[node] {
		return
	}
	visited[node] = true
	(*path) = append((*path), node)
	for connection := range graph[node] {
		next := true
		for _, prev_node := range *path {
			if _, exists := graph[prev_node][connection]; !exists {
				next = false
			}
		}
		if next {
			largest_connected_component(graph, connection, path, visited)
		}
	}
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	graph := make(map[string]map[string]bool)
	for _, line := range lines {
		left := strings.Split(line, "-")[0]
		right := strings.Split(line, "-")[1]
		if _, exists := graph[left]; !exists {
			graph[left] = make(map[string]bool)
		}
		if _, exists := graph[right]; !exists {
			graph[right] = make(map[string]bool)
		}
		graph[left][right] = true
		graph[right][left] = true
	}
	max_len := 0
	var max_path []string
	for node := range graph {
		path := make([]string, 0)
		visited := make(map[string]bool)
		largest_connected_component(graph, node, &path, visited)
		if len(path) > max_len {
			max_len = len(path)
			max_path = path
		}
	}
	sort.Strings(max_path)
	return strings.Join(max_path, ",")
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
