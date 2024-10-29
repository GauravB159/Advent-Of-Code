package main

import (
	"fmt"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateGrid(lines)
	grid_image := aocutils.CreateImage(grid.Rows, grid.Cols, 40, "onestar")
	grid_image.UsePaletteWideReds(true)
	total_steps := 100
	gif := aocutils.CreateGIF("onestar", 5)
	for i, line := range lines {
		for j := range line {
			value, _ := grid.GetValue(i, j)
			grid_image.SetZoomedPixel(i, j, value)
		}
	}
	total_flashes := 0
	gif.AddFrame(grid_image)
	for k := 0; k < total_steps; k++ {
		flashed := make(map[aocutils.Key]bool)
		var to_flash aocutils.Stack[aocutils.Key]
		for i, line := range lines {
			for j := range line {
				value, _ := grid.GetValue(i, j)
				value = value + 1
				if value > 9 {
					value = 10
					to_flash.Push(aocutils.Key{Row: i, Col: j})
				}
				grid.SetValue(i, j, value)
				grid_image.SetZoomedPixel(i, j, value)
			}
		}
		gif.AddFrame(grid_image)
		for !to_flash.IsEmpty() {
			key, _ := to_flash.Pop()
			if _, exists := flashed[key]; exists {
				continue
			}
			for m := -1; m <= 1; m++ {
				for l := -1; l <= 1; l++ {
					if m == 0 && l == 0 {
						continue
					}
					newKey := aocutils.Key{Row: key.Row + m, Col: key.Col + l}
					if _, exists := grid.Data[newKey]; exists {
						grid.Data[newKey] += 1
						grid_image.SetZoomedPixel(key.Row+m, key.Col+l, grid.Data[newKey])
						if grid.Data[newKey] > 9 {
							grid.Data[newKey] = 9
							to_flash.Push(newKey)
						}
					}
				}
			}
			gif.AddFrame(grid_image)
			total_flashes += 1
			flashed[key] = true
		}
		for key := range flashed {
			grid.Data[key] = 0
		}

	}
	gif.WriteGIFToFile()
	return strconv.Itoa(total_flashes)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateGrid(lines)
	grid_image := aocutils.CreateImage(grid.Rows, grid.Cols, 40, "twostar")
	grid_image.UsePaletteWideReds(true)
	gif := aocutils.CreateGIF("twostar", 10)
	for i, line := range lines {
		for j := range line {
			value, _ := grid.GetValue(i, j)
			grid_image.SetZoomedPixel(i, j, value)
		}
	}
	gif.AddFrame(grid_image)
	answer := 0
	for k := 0; ; k++ {
		flashed := make(map[aocutils.Key]bool)
		var to_flash aocutils.Stack[aocutils.Key]
		for i, line := range lines {
			for j := range line {
				value, _ := grid.GetValue(i, j)
				value = value + 1
				if value > 9 {
					value = 10
					to_flash.Push(aocutils.Key{Row: i, Col: j})
				}
				grid.SetValue(i, j, value)
				grid_image.SetZoomedPixel(i, j, value)
			}
		}
		gif.AddFrame(grid_image)
		step_flashed := 0

		for !to_flash.IsEmpty() {
			key, _ := to_flash.Pop()
			if _, exists := flashed[key]; exists {
				continue
			}
			for m := -1; m <= 1; m++ {
				for l := -1; l <= 1; l++ {
					if m == 0 && l == 0 {
						continue
					}
					newKey := aocutils.Key{Row: key.Row + m, Col: key.Col + l}
					if _, exists := grid.Data[newKey]; exists {
						grid.Data[newKey] += 1
						grid_image.SetZoomedPixel(key.Row+m, key.Col+l, grid.Data[newKey])
						if grid.Data[newKey] > 9 {
							grid.Data[newKey] = 9
							to_flash.Push(newKey)
						}
					}
				}
			}
			step_flashed += 1

			gif.AddFrame(grid_image)
			flashed[key] = true
		}
		if step_flashed == (grid.Rows * grid.Cols) {
			answer = k + 1
			for a := 0; a < 20; a++ {
				gif.AddFrame(grid_image)
			}
			break
		}
		for key := range flashed {
			grid.Data[key] = 0
		}
	}
	gif.WriteGIFToFile()
	return strconv.Itoa(answer)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
