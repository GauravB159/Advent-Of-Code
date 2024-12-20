package main

import (
	"fmt"
	"math"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

type Node struct {
	position aocutils.Key
	level    int
}

func bfs(start aocutils.Key, grid *aocutils.Grid) int {
	directions := [4]aocutils.Key{{Row: 1, Col: 0}, {Row: -1, Col: 0}, {Row: 0, Col: 1}, {Row: 0, Col: -1}}
	queue := make([]Node, 0)
	queue = append(queue, Node{position: start, level: 0})
	visited := make(map[aocutils.Key]bool)
	path_count := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if visited[node.position] {
			continue
		}
		if grid.Data[node.position] == 'E' {
			path_count = node.level
			break
		}
		visited[node.position] = true
		for _, direction := range directions {
			next := aocutils.Key{Row: node.position.Row + direction.Row, Col: node.position.Col + direction.Col}
			if grid.Data[next] == '.' || grid.Data[next] == 'E' {
				queue = append(queue, Node{position: next, level: node.level + 1})
			}
		}
	}
	return path_count
}

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	var start aocutils.Key
	for position, value := range grid.Data {
		if value == 'S' {
			start = position
		}
	}
	path := make([]aocutils.Key, 0)
	node := start
	directions := [4]aocutils.Key{{Row: 1, Col: 0}, {Row: -1, Col: 0}, {Row: 0, Col: 1}, {Row: 0, Col: -1}}
	visited := make(map[aocutils.Key]bool, 0)
	path = append(path, start)
	for true {
		if grid.Data[node] == 'E' {
			break
		}
		visited[node] = true
		for _, direction := range directions {
			next_node := aocutils.Key{Row: node.Row + direction.Row, Col: node.Col + direction.Col}
			if visited[next_node] {
				continue
			}
			if grid.Data[next_node] == '.' || grid.Data[next_node] == 'E' {
				path = append(path, next_node)
				node = next_node
				break
			}
		}
	}
	distances := make(map[aocutils.Key]int)
	for idx, node := range path {
		distances[node] = len(path) - 1 - idx
	}
	counts := make(map[int]int)
	type Pair struct {
		start aocutils.Key
		end   aocutils.Key
	}
	count := 0
	answer := 0
	for position, value := range grid.Data {
		if value == '.' || value == 'E' || value == 'S' {
			for second_position, second_value := range grid.Data {
				if second_value == '.' || second_value == 'E' || second_value == 'S' {
					distance := int(math.Abs(float64(second_position.Col)-float64(position.Col)) + math.Abs(float64(second_position.Row)-float64(position.Row)))
					if distance > 1 && distance < 3 {
						count += 1
						pos_dist := distances[position]
						second_pos_dist := distances[second_position]
						if (second_pos_dist-pos_dist)-distance >= 100 {
							answer += 1
							counts[(second_pos_dist-pos_dist)-distance] += 1
						}
					}
				}
			}
		}
	}
	return strconv.Itoa(answer)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	var start aocutils.Key
	for position, value := range grid.Data {
		if value == 'S' {
			start = position
		}
	}
	path := make([]aocutils.Key, 0)
	node := start
	directions := [4]aocutils.Key{{Row: 1, Col: 0}, {Row: -1, Col: 0}, {Row: 0, Col: 1}, {Row: 0, Col: -1}}
	visited := make(map[aocutils.Key]bool, 0)
	path = append(path, start)
	for true {
		if grid.Data[node] == 'E' {
			break
		}
		visited[node] = true
		for _, direction := range directions {
			next_node := aocutils.Key{Row: node.Row + direction.Row, Col: node.Col + direction.Col}
			if visited[next_node] {
				continue
			}
			if grid.Data[next_node] == '.' || grid.Data[next_node] == 'E' {
				path = append(path, next_node)
				node = next_node
				break
			}
		}
	}
	distances := make(map[aocutils.Key]int)
	for idx, node := range path {
		distances[node] = len(path) - 1 - idx
	}
	counts := make(map[int]int)
	type Pair struct {
		start aocutils.Key
		end   aocutils.Key
	}
	count := 0
	answer := 0
	for position, value := range grid.Data {
		if value == '.' || value == 'E' || value == 'S' {
			for second_position, second_value := range grid.Data {
				if second_value == '.' || second_value == 'E' || second_value == 'S' {
					distance := int(math.Abs(float64(second_position.Col)-float64(position.Col)) + math.Abs(float64(second_position.Row)-float64(position.Row)))
					if distance > 1 && distance < 21 {
						count += 1
						pos_dist := distances[position]
						second_pos_dist := distances[second_position]
						if (second_pos_dist-pos_dist)-distance >= 100 {
							answer += 1
							counts[(second_pos_dist-pos_dist)-distance] += 1
						}
					}
				}
			}
		}
	}
	return strconv.Itoa(answer)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
