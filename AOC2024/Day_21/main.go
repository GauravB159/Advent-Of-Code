package main

import (
	"fmt"
	"math"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func find_possible_paths(node aocutils.Key, end aocutils.Key, current_path string, visited map[aocutils.Key]bool, possible_paths *[]string, grid aocutils.Grid) {
	if _, exists := grid.Data[node]; !exists {
		return
	}
	if grid.Data[node] == '!' {
		return
	}
	if node.Row == end.Row && node.Col == end.Col {
		*possible_paths = append(*possible_paths, current_path+"A")
		return
	}
	if visited[node] {
		return
	}
	directions := map[string]aocutils.Key{
		">": {Row: 0, Col: 1},
		"v": {Row: 1, Col: 0},
		"<": {Row: 0, Col: -1},
		"^": {Row: -1, Col: 0},
	}
	visited[node] = true
	for key, direction := range directions {
		next_node := aocutils.Key{Row: node.Row + direction.Row, Col: node.Col + direction.Col}
		find_possible_paths(next_node, end, current_path+key, visited, possible_paths, grid)
	}
	visited[node] = false
	return
}

func dfs(node string, path string, new_path string, level int, min_direction_path_map *map[string][]string, new_paths *[]string) {
	if level == len(path) {
		*new_paths = append(*new_paths, new_path)
		return
	}
	next_node := string(path[level])
	key := node + next_node
	for _, dir_path := range (*min_direction_path_map)[key] {
		dfs(next_node, path, new_path+dir_path, level+1, min_direction_path_map, new_paths)
	}
}

func get_min_paths(grid aocutils.Grid) map[string][]string {
	min_path_map := make(map[string][]string)
	for position, val := range grid.Data {
		if val == '!' {
			continue
		}
		for second_position, second_val := range grid.Data {
			if second_val == '!' {
				continue
			}
			visited := make(map[aocutils.Key]bool)
			var possible_paths []string
			find_possible_paths(position, second_position, "", visited, &possible_paths, grid)
			var min_path_length int = math.MaxInt64
			for _, path := range possible_paths {
				if len(path) < min_path_length {
					min_path_length = len(path)
				}
			}
			var min_paths []string
			for _, path := range possible_paths {
				if len(path) == min_path_length {
					min_paths = append(min_paths, path)
				}
			}
			key := string(rune(val)) + string(rune(second_val))
			min_path_map[key] = min_paths
		}
	}
	return min_path_map
}

func drill(path_map *map[string][]string, min_direction_path_map *map[string][]string) map[string][]string {
	next_level_map := make(map[string][]string)
	for key, paths := range *path_map {
		mapped := make([]string, 0)
		for _, path := range paths {
			start := "A"
			dfs(start, path, "", 0, min_direction_path_map, &mapped)
		}
		var min_path_length int = math.MaxInt64
		for _, path := range mapped {
			if len(path) < min_path_length {
				min_path_length = len(path)
			}
		}
		var min_paths []string
		for _, path := range mapped {
			if len(path) == min_path_length {
				min_paths = append(min_paths, path)
			}
		}
		next_level_map[key] = min_paths
	}
	return next_level_map
}

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	keypad := lines[:4]
	lines = lines[5:]
	directions := lines[:2]
	lines = lines[3:]
	keypad_grid := aocutils.CreateCharacterGrid(keypad)
	min_path_map := get_min_paths(keypad_grid)

	direction_grid := aocutils.CreateCharacterGrid(directions)
	min_direction_path_map := get_min_paths(direction_grid)
	second_level_map := drill(&min_path_map, &min_direction_path_map)
	third_level_map := drill(&second_level_map, &min_direction_path_map)
	result := 0
	for _, line := range lines {
		start := "A"
		sum := 0
		num, _ := strconv.Atoi(line[:3])
		for _, char := range line {
			key := start + string(char)
			sum += len(third_level_map[key][0])
			start = string(char)
		}
		result += num * sum
	}
	return strconv.Itoa(result)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	keypad := lines[:4]
	lines = lines[5:]
	directions := lines[:2]
	lines = lines[3:]
	keypad_grid := aocutils.CreateCharacterGrid(keypad)
	min_path_map := get_min_paths(keypad_grid)

	direction_grid := aocutils.CreateCharacterGrid(directions)
	min_direction_path_map := get_min_paths(direction_grid)
	drill_level := 2
	nth_level_map := min_path_map
	for range drill_level {
		nth_level_map = drill(&nth_level_map, &min_direction_path_map)
	}
	result := 0
	for _, line := range lines {
		start := "A"
		sum := 0
		num, _ := strconv.Atoi(line[:3])
		for _, char := range line {
			key := start + string(char)
			sum += len(nth_level_map[key][0])
			start = string(char)
		}
		result += num * sum
	}
	return strconv.Itoa(result)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	// aocutils.Timer("2 star", twostar, "input.txt")
}
