package main

import (
	"fmt"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func connected_component(point aocutils.Key, check rune, grid *aocutils.Grid, visited *map[aocutils.Key]bool, points *[]aocutils.Key) {
	if value, exists := (*visited)[point]; exists && value {
		return
	}
	if rune((*grid).Data[point]) != check {
		return
	}
	*points = append(*points, point)
	(*visited)[point] = true
	connected_component(aocutils.Key{Row: point.Row + 1, Col: point.Col}, check, grid, visited, points)
	connected_component(aocutils.Key{Row: point.Row - 1, Col: point.Col}, check, grid, visited, points)
	connected_component(aocutils.Key{Row: point.Row, Col: point.Col + 1}, check, grid, visited, points)
	connected_component(aocutils.Key{Row: point.Row, Col: point.Col - 1}, check, grid, visited, points)
}

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	visited := make(map[aocutils.Key]bool, 0)
	price := 0
	for key := range grid.Data {
		points := make([]aocutils.Key, 0)
		if value, exists := visited[key]; exists && value {
			continue
		}
		check := rune(grid.Data[key])
		connected_component(key, check, &grid, &visited, &points)
		perimeter := 0
		area := len(points)
		var directions [4]aocutils.Key = [4]aocutils.Key{{Row: 1, Col: 0}, {Row: -1, Col: 0}, {Row: 0, Col: 1}, {Row: 0, Col: -1}}
		for _, point := range points {
			neighbor_count := 0
			for _, direction := range directions {
				new_point := aocutils.Key{Row: point.Row + direction.Row, Col: point.Col + direction.Col}
				if grid.Data[new_point] == grid.Data[point] {
					neighbor_count += 1
				}
			}
			perimeter += (4 - neighbor_count)
		}
		price += perimeter * area
	}
	return strconv.Itoa(price)
}

type MidpointKey struct {
	Row float32
	Col float32
}

func edge_dfs(point MidpointKey, direction int, edges *map[MidpointKey]map[int]MidpointKey, visited *map[MidpointKey]bool, visited_edges_with_directions *map[directionWithKey]bool, path *[]directionWithKey) {
	if (*visited_edges_with_directions)[directionWithKey{direction: direction, key: point}] {
		return
	}
	*path = append(*path, directionWithKey{direction: direction, key: point})
	(*visited)[point] = true
	(*visited_edges_with_directions)[directionWithKey{direction: direction, key: point}] = true
	direction_change := map[int]int{
		2: 1,
		1: 2,
		3: 0,
		0: 3,
	}
	fmt.Println(point, direction, (*edges)[point])
	var next_direction int
	var next_point MidpointKey
	for key, value := range (*edges)[point] {
		next_point = value
		next_direction = key
		if direction_change[direction] == next_direction {
			break
		}
	}
	edge_dfs(next_point, next_direction, edges, visited, visited_edges_with_directions, path)
}

type directionWithKey struct {
	direction int
	key       MidpointKey
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	visited := make(map[aocutils.Key]bool, 0)
	price := 0
	for key := range grid.Data {
		points := make([]aocutils.Key, 0)
		if value, exists := visited[key]; exists && value {
			continue
		}
		check := rune(grid.Data[key])
		connected_component(key, check, &grid, &visited, &points)
		area := len(points)
		var directions [4]aocutils.Key = [4]aocutils.Key{{Row: 1, Col: 0}, {Row: -1, Col: 0}, {Row: 0, Col: 1}, {Row: 0, Col: -1}}
		edges := make(map[MidpointKey]map[int]MidpointKey, 0)
		direction_change := map[int]int{
			0: 2,
			1: 3,
			2: 1,
			3: 0,
		}
		for _, point := range points {
			for idx, direction := range directions {
				new_point := aocutils.Key{Row: point.Row + direction.Row, Col: point.Col + direction.Col}
				if grid.Data[new_point] != grid.Data[point] {
					midpoint := MidpointKey{Row: (float32(point.Row) + float32(new_point.Row)) / 2, Col: (float32(point.Col) + float32(new_point.Col)) / 2}
					if point.Row == new_point.Row {
						point_one := MidpointKey{Row: midpoint.Row - 0.5, Col: midpoint.Col}
						point_two := MidpointKey{Row: midpoint.Row + 0.5, Col: midpoint.Col}
						if idx == 3 {
							if len(edges[point_one]) == 0 {
								edges[point_one] = make(map[int]MidpointKey)
							}
							edges[point_one][direction_change[idx]] = point_two
						} else if idx == 2 {
							if len(edges[point_two]) == 0 {
								edges[point_two] = make(map[int]MidpointKey)
							}
							edges[point_two][direction_change[idx]] = point_one
						}
					} else {
						point_one := MidpointKey{Row: midpoint.Row, Col: midpoint.Col - 0.5}
						point_two := MidpointKey{Row: midpoint.Row, Col: midpoint.Col + 0.5}
						if idx == 0 {
							if len(edges[point_one]) == 0 {
								edges[point_one] = make(map[int]MidpointKey)
							}
							edges[point_one][direction_change[idx]] = point_two
						} else if idx == 1 {
							if len(edges[point_two]) == 0 {
								edges[point_two] = make(map[int]MidpointKey)
							}
							edges[point_two][direction_change[idx]] = point_one
						}
					}
					// fmt.Println(edges, point, new_point, idx, "HERE")
				}
			}
		}
		visited_edges := make(map[MidpointKey]bool, 0)
		visited_edges_with_directions := make(map[directionWithKey]bool, 0)
		count := 0
		for point := range edges {
			if visited_edges[point] {
				continue
			}
			path := make([]directionWithKey, 0)
			keys := make([]int, len(edges[point]))
			i := 0
			for k := range edges[point] {
				keys[i] = k
				i++
			}
			edge_dfs(point, keys[0], &edges, &visited_edges, &visited_edges_with_directions, &path)
			for i := range path {
				if path[(i+1)%len(path)].direction != path[i].direction {
					count += 1
				}
			}
		}
		price += count * area
	}
	return strconv.Itoa(price)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
