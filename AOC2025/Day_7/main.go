package main

import (
	"fmt"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := 0
	grid := aocutils.CreateCharacterGrid(lines)
	grid_image := aocutils.CreateImage(grid.Rows, grid.Cols, 5, 5, "onestar")
	grid_image.UsePaletteRedToYellow()
	gif := aocutils.CreateGIF("onestar", 50)
	var start_pos aocutils.Key
	direction := aocutils.Key{Row: 1, Col: 0}
	grid_color_map := map[string]int{
		"S": 1,
		".": 1,
		"^": 5,
		"|": 9,
	}
	for key, value := range grid.Data {
		if value == 'S' {
			start_pos = key
			continue
		}
		grid_image.SetZoomedPixel(key.Col, key.Row, grid_color_map[string(rune(value))])
	}
	var stack aocutils.Stack[aocutils.Key]
	stack = append([]aocutils.Key{start_pos}, stack...)
	visited := make(map[aocutils.Key]bool)
	for len(stack) != 0 {
		node, _ := stack.Pop()
		if _, exists := grid.Data[node]; !exists {
			continue
		}
		if _, exists := visited[node]; exists {
			continue
		}
		visited[node] = true
		if grid.Data[node] == '^' {
			result += 1
			stack = append([]aocutils.Key{{Row: node.Row, Col: node.Col + 1}}, stack...)
			stack = append([]aocutils.Key{{Row: node.Row, Col: node.Col - 1}}, stack...)
		} else {
			grid_image.SetZoomedPixel(node.Col, node.Row, grid_color_map["|"])
			stack = append([]aocutils.Key{{Row: node.Row + direction.Row, Col: node.Col + direction.Col}}, stack...)
		}
		gif.AddFrame(grid_image)
	}
	gif.WriteGIFToFile()
	return strconv.Itoa(result)
}
func recurse(node aocutils.Key, total int, cache *map[aocutils.Key]int, grid *aocutils.Grid, grid_image *aocutils.Image, gif *aocutils.GIF) int {
	if _, exists := (*grid).Data[node]; !exists {
		return 1
	}
	if _, exists := (*cache)[node]; exists {
		return (*cache)[node]
	}
	var new_total int

	// grid_color_map := map[string]int{
	// 	".": 1,
	// 	"^": 5,
	// 	"|": 9,
	// }
	// temp := grid.Data[node]
	if grid.Data[node] == '^' {
		new_total = total
		gif.AddFrame(*grid_image)
		new_total += recurse(aocutils.Key{Row: node.Row, Col: node.Col + 1}, 0, cache, grid, grid_image, gif)
		gif.AddFrame(*grid_image)
		new_total += recurse(aocutils.Key{Row: node.Row, Col: node.Col - 1}, 0, cache, grid, grid_image, gif)
		gif.AddFrame(*grid_image)
	} else {
		grid_image.SetZoomedPixel(node.Col, node.Row, 9)
		gif.AddFrame(*grid_image)
		new_total += recurse(aocutils.Key{Row: node.Row + 1, Col: node.Col}, 0, cache, grid, grid_image, gif)
		gif.AddFrame(*grid_image)
	}
	grid_image.SetZoomedPixel(node.Col, node.Row, 3)
	(*cache)[node] = new_total
	return new_total
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	grid_image := aocutils.CreateImage(grid.Rows, grid.Cols, 5, 5, "twostar")
	grid_image.UsePaletteRedToYellow()
	gif := aocutils.CreateGIF("twostar", 100)
	grid_color_map := map[string]int{
		".": 1,
		"^": 5,
		"|": 9,
	}
	var start_pos aocutils.Key
	for key, value := range grid.Data {
		if value == 'S' {
			start_pos = key
			continue
		}
		grid_image.SetZoomedPixel(key.Col, key.Row, grid_color_map[string(rune(value))])
	}
	cache := make(map[aocutils.Key]int)
	result := recurse(start_pos, 0, &cache, &grid, &grid_image, &gif)
	gif.WriteGIFToFile()
	return strconv.Itoa(result)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
