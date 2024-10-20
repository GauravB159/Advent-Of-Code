package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
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
	grid_image := image.NewRGBA(image.Rect(0, 0, cols*zoom, rows*zoom))
	visited := make(map[Key]bool, len(lines)*len(lines[0]))

	var images []*image.Paletted
	var delays []int

	palette := make([]color.Color, 0, 11)
	palette = append(palette, color.Black)
	for i := 0; i <= 9; i++ {
		palette = append(palette, color.RGBA{R: 60 + uint8(195*(float64(10-i)/10)), G: 0, B: 0, A: 255})
	}
	palette = append(palette, color.Transparent)
	var stack aocutils.Stack[Key]
	for i, line := range lines {
		for j, char := range line {
			value, _ := strconv.Atoi(string(char))
			key := Key{row: i, col: j}
			for k := i * zoom; k < (i+1)*zoom; k++ {
				for l := j * zoom; l < (j+1)*zoom; l++ {
					pixel := (k*cols*zoom + l) * 4
					grid_image.Pix[pixel] = 0
					grid_image.Pix[pixel+3] = 255
				}
			}
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
				for k := location.row * zoom; k < (location.row+1)*zoom; k++ {
					for l := location.col * zoom; l < (location.col+1)*zoom; l++ {
						pixel := (k*cols*zoom + l) * 4
						grid_image.Pix[pixel] = 60 + uint8(195*(float64(10-grid[location])/10))
						grid_image.Pix[pixel+3] = 255
					}
				}
				if pixel_count%frameskip == 0 {
					palettedImage := image.NewPaletted(grid_image.Bounds(), palette)
					draw.FloydSteinberg.Draw(palettedImage, grid_image.Bounds(), grid_image, image.Point{})
					images = append(images, palettedImage)
					delays = append(delays, 1)
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
	lastFrame := images[len(images)-1]
	for i := 0; i < 20; i++ {
		images = append(images, lastFrame)
		delays = append(delays, 100)
	}
	gifFile, err := os.Create("twostar.gif")
	if err != nil {
		panic(err)
	}
	defer gifFile.Close()

	err = gif.EncodeAll(gifFile, &gif.GIF{
		Image: images,
		Delay: delays,
	})
	if err != nil {
		panic(err)
	}
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
