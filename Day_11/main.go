package main

import (
	"fmt"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := ""
	grid := aocutils.CreateGrid(lines)
	grid_image := aocutils.CreateImage(grid.Rows, grid.Cols, 40, "onestar")
	grid_image.UsePaletteWideReds()
	total_steps := 100
	gif := aocutils.CreateGIF("onestar", 1)
	for i, line := range lines {
		for j := range line {
			value, _ := grid.GetValue(i, j)
			grid_image.SetZoomedPixel(i, j, value)
		}
	}
	gif.AddFrame(grid_image)
	for k := 0; k < total_steps; k++ {
		for i, line := range lines {
			for j := range line {
				value, _ := grid.GetValue(i, j)
				value = (value + 1) % 10
				grid.SetValue(i, j, value)
				grid_image.SetZoomedPixel(i, j, value)
			}
		}
		gif.AddFrame(grid_image)
	}
	grid_image.WritePNGToFile()
	gif.WriteGIFToFile()
	return result
}

// func twostar(filename string) string {
// 	lines := aocutils.Readfile(filename)
// 	result := ""
// 	for _, line := range lines {
// 		fmt.Println(line)
// 	}
// 	return result
// }

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	// aocutils.Timer("2 star", twostar, "input.txt")
}
