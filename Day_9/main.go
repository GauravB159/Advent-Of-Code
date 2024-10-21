package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"

	"os"
	"sort"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

type Key struct {
	row int
	col int
}

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	rows := len(lines)
	cols := len(lines[0])
	grid := make(map[Key]int, rows*cols)

	zoom := 10
	grid_image := image.NewRGBA(image.Rect(0, 0, cols*zoom, rows*zoom))
	outputFile, _ := os.Create("onestar.png")

	for row, line := range lines {
		for col, char := range line {
			value, _ := strconv.Atoi(string(char))
			grid[Key{row: row, col: col}] = value + 1
		}
	}
	result := 0
	for i, line := range lines {
		for j := range line {
			grid_value := grid[Key{row: i, col: j}]
			for k := i * zoom; k < (i+1)*zoom; k++ {
				for l := j * zoom; l < (j+1)*zoom; l++ {
					location := (k*cols*zoom + l) * 4
					grid_image.Pix[location] = 60 + uint8(175*(float64(10-grid[Key{row: i, col: j}])/10))
					grid_image.Pix[location+3] = 255
				}
			}
			if value, exists := grid[Key{row: i + 1, col: j}]; exists && value <= grid_value {
				continue
			} else if value, exists := grid[Key{row: i - 1, col: j}]; exists && value <= grid_value {
				continue
			} else if value, exists := grid[Key{row: i, col: j + 1}]; exists && value <= grid_value {
				continue
			} else if value, exists := grid[Key{row: i, col: j - 1}]; exists && value <= grid_value {
				continue
			}
			result += grid[Key{row: i, col: j}]
		}
	}
	png.Encode(outputFile, grid_image)
	outputFile.Close()
	return strconv.Itoa(result)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	rows := len(lines)
	cols := len(lines[0])
	grid := make(map[Key]int, rows*cols)
	image_grid := make(map[Key]int, rows*cols)

	zoom := 10
	visited := make(map[Key]bool, len(lines)*len(lines[0]))

	var images []*image.Paletted

	palette := make([]color.Color, 0, 11)
	palette = append(palette, color.Black)
	for i := 0; i <= 9; i++ {
		palette = append(palette, color.RGBA{R: 60 + uint8(195*(float64(10-i)/10)), G: 0, B: 0, A: 255})
	}
	palette = append(palette, color.Transparent)
	grid_image := image.NewPaletted(image.Rect(0, 0, cols*zoom, rows*zoom), palette)

	var stack aocutils.Stack[Key]
	for i, line := range lines {
		for j, char := range line {
			value, _ := strconv.Atoi(string(char))
			key := Key{row: i, col: j}
			aocutils.SetZoomedPixel(j, i, zoom, grid_image, 0)
			grid[key] = value
			image_grid[key] = 0
			visited[key] = false
		}
	}
	basinSize := make([]int, 10)
	pixel_count := 0
	frameskip := 50
	for i, line := range lines {
		for j := range line {
			key := Key{row: i, col: j}
			if visited[key] {
				continue
			}
			stack.Push(key)
			count := 0
			for !stack.IsEmpty() {
				location, _ := stack.Pop()
				if visited[location] {
					continue
				}
				visited[location] = true
				if grid[location] == 9 {
					continue
				}
				pixel_count += 1
				aocutils.SetZoomedPixel(location.col, location.row, zoom, grid_image, grid[location])
				if pixel_count%frameskip == 0 {
					copied_image := image.NewPaletted(grid_image.Rect, grid_image.Palette)
					copy(copied_image.Pix, grid_image.Pix)
					images = append(images, copied_image)
				}
				count += 1
				new_location := Key{row: location.row + 1, col: location.col}
				if value, exists := grid[new_location]; exists && value != 9 {
					stack.Push(new_location)
				}
				new_location = Key{row: location.row - 1, col: location.col}
				if value, exists := grid[new_location]; exists && value != 9 {
					stack.Push(new_location)
				}
				new_location = Key{row: location.row, col: location.col + 1}
				if value, exists := grid[new_location]; exists && value != 9 {
					stack.Push(new_location)
				}
				new_location = Key{row: location.row, col: location.col - 1}
				if value, exists := grid[new_location]; exists && value != 9 {
					stack.Push(new_location)
				}
			}
			basinSize = append(basinSize, count)
		}
	}
	aocutils.CreateGIF(images, palette, "twostar")
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
