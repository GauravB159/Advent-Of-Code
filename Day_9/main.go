package main

import (
	"fmt"

	"sort"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateGrid(lines)
	grid_image := aocutils.CreateImage(grid.GetNumRows(), grid.GetNumCols(), 10, "onestar")
	grid_image.UsePaletteWideReds()

	result := 0
	for i, line := range lines {
		for j := range line {
			grid_value, _ := grid.GetValue(i, j)
			grid_image.SetZoomedPixel(i, j, grid_value)
			if value, exists := grid.GetValue(i+1, j); exists && value <= grid_value {
				continue
			} else if value, exists := grid.GetValue(i-1, j); exists && value <= grid_value {
				continue
			} else if value, exists := grid.GetValue(i, j+1); exists && value <= grid_value {
				continue
			} else if value, exists := grid.GetValue(i, j-1); exists && value <= grid_value {
				continue
			}
			result += grid_value
		}
	}
	grid_image.WritePNGToFile()
	return strconv.Itoa(result)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	rows := len(lines)
	cols := len(lines[0])
	grid := aocutils.CreateGrid(lines)
	image_grid := aocutils.CreateGrid(lines)
	visited := aocutils.CreateGrid(lines)

	grid_image := aocutils.CreateImage(rows, cols, 10, "twostar")
	grid_image.UsePaletteReds()
	gif := aocutils.CreateGIF("twostar", 10)
	var stack aocutils.Stack[aocutils.Key]
	for i, line := range lines {
		for j, char := range line {
			value, _ := strconv.Atoi(string(char))
			grid_image.SetZoomedPixel(j, i, 0)
			grid.SetValue(i, j, value)
			image_grid.SetValue(i, j, 0)
			visited.SetValue(i, j, 0)
		}
	}
	basinSize := make([]int, 10)
	for i, line := range lines {
		for j := range line {
			key := aocutils.Key{Row: i, Col: j}
			if value, _ := visited.GetValue(i, j); value == 1 {
				continue
			}
			stack.Push(key)
			count := 0
			for !stack.IsEmpty() {
				location, _ := stack.Pop()
				if visited.Data[location] == 1 {
					continue
				}
				visited.Data[location] = 1
				if grid.Data[location] == 9 {
					continue
				}
				gif.AddFrame(grid_image)
				grid_image.SetZoomedPixel(location.Col, location.Row, grid.Data[location])
				count += 1
				new_location := aocutils.Key{Row: location.Row + 1, Col: location.Col}
				if value, exists := grid.Data[new_location]; exists && value != 9 {
					stack.Push(new_location)
				}
				new_location = aocutils.Key{Row: location.Row - 1, Col: location.Col}
				if value, exists := grid.Data[new_location]; exists && value != 9 {
					stack.Push(new_location)
				}
				new_location = aocutils.Key{Row: location.Row, Col: location.Col + 1}
				if value, exists := grid.Data[new_location]; exists && value != 9 {
					stack.Push(new_location)
				}
				new_location = aocutils.Key{Row: location.Row, Col: location.Col - 1}
				if value, exists := grid.Data[new_location]; exists && value != 9 {
					stack.Push(new_location)
				}
			}
			basinSize = append(basinSize, count)
		}
	}
	gif.WriteGIFToFile()
	sort.Ints(basinSize)
	top_3 := basinSize[len(basinSize)-3:]
	return strconv.Itoa(top_3[0] * top_3[1] * top_3[2])
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
