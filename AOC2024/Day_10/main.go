package main

import (
	"fmt"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func dfs_one(grid aocutils.Grid, node aocutils.Key, level int, finals map[aocutils.Key]bool, visited map[aocutils.Key]bool, path []aocutils.Key, grid_image *aocutils.Image, gif *aocutils.GIF) {
	if _, exists := grid.Data[node]; !exists {
		return
	}
	if grid.Data[node] != level {
		return
	}
	if _, exists := finals[node]; !exists && level == 9 {
		for index, innernode := range path {
			grid_image.SetZoomedPixel(innernode.Row, innernode.Col, index+1)
			gif.AddFrame(*grid_image)
		}
		grid_image.SetZoomedPixel(node.Row, node.Col, 10)
		gif.AddFrame(*grid_image)
		finals[node] = true
		return
	}
	visited[node] = true
	path = append(path, node)
	dfs_one(grid, aocutils.Key{Row: node.Row + 1, Col: node.Col}, level+1, finals, visited, path, grid_image, gif)
	dfs_one(grid, aocutils.Key{Row: node.Row - 1, Col: node.Col}, level+1, finals, visited, path, grid_image, gif)
	dfs_one(grid, aocutils.Key{Row: node.Row, Col: node.Col + 1}, level+1, finals, visited, path, grid_image, gif)
	dfs_one(grid, aocutils.Key{Row: node.Row, Col: node.Col - 1}, level+1, finals, visited, path, grid_image, gif)
	if len(path) > 2 {
		path = path[:len(path)-2]
	}
	visited[node] = false
}

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateGrid(lines)
	grid_image := aocutils.CreateImage(grid.GetNumRows(), grid.GetNumCols(), 15, "onestar")
	grid_image.UsePaletteRedToYellow()
	gif := aocutils.CreateGIF("onestar", 10)
	sum := 0
	for position, value := range grid.Data {
		if value == 0 {
			grid_image.SetZoomedPixel(position.Row, position.Col, 1)
			gif.AddFrame(grid_image)
		}
	}
	for position, value := range grid.Data {
		if value == 9 {
			grid_image.SetZoomedPixel(position.Row, position.Col, 10)
			gif.AddFrame(grid_image)
		}
	}
	for position, value := range grid.Data {
		visited := make(map[aocutils.Key]bool, 0)
		finals := make(map[aocutils.Key]bool, 0)
		path := make([]aocutils.Key, 0)
		if value == 0 {
			dfs_one(grid, position, 0, finals, visited, path, &grid_image, &gif)
			sum += len(finals)
		}
	}
	gif.WriteGIFToFile()
	return strconv.Itoa(sum)
}

func dfs_two(grid aocutils.Grid, node aocutils.Key, level int, count *int, visited map[aocutils.Key]bool) {
	if _, exists := grid.Data[node]; !exists {
		return
	}
	if grid.Data[node] != level {
		return
	}
	if level == 9 {
		*count += 1
		return
	}
	visited[node] = true
	dfs_two(grid, aocutils.Key{Row: node.Row + 1, Col: node.Col}, level+1, count, visited)
	dfs_two(grid, aocutils.Key{Row: node.Row - 1, Col: node.Col}, level+1, count, visited)
	dfs_two(grid, aocutils.Key{Row: node.Row, Col: node.Col + 1}, level+1, count, visited)
	dfs_two(grid, aocutils.Key{Row: node.Row, Col: node.Col - 1}, level+1, count, visited)
	visited[node] = false
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateGrid(lines)
	sum := 0
	for position, value := range grid.Data {
		count := 0
		visited := make(map[aocutils.Key]bool, 0)

		if value == 0 {
			dfs_two(grid, position, 0, &count, visited)
			sum += count
		}
	}
	return strconv.Itoa(sum)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
