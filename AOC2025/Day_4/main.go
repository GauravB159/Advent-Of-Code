package main

import (
	"fmt"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	grid_image := aocutils.CreateImage(grid.Rows, grid.Cols, 5, 5, "onestar")
	grid_image.UseFullColors()
	sum := 0
	valmap := map[string]int{
		".": 1,
		"@": 7,
	}
	gif := aocutils.CreateGIF("onestar", 100)

	for i, line := range lines {
		for j := range line {
			value, _ := grid.GetValue(i, j)
			grid_image.SetZoomedPixel(j, i, valmap[string(rune(value))])
		}
	}

	for i, line := range lines {
		for j := range line {
			value, _ := grid.GetValue(i, j)
			count_ := 0
			if rune(value) != '@' {
				continue
			}
			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if x == 0 && y == 0 {
						continue
					}
					check, _ := grid.GetValue(i+x, j+y)
					if rune(check) == '@' {
						count_ += 1
					}
				}
			}
			if count_ < 4 {
				// fmt.Println(i, j)
				grid_image.SetZoomedPixel(j, i, 5)
				sum += 1
			}
			gif.AddFrame(grid_image)
		}
	}
	gif.WriteGIFToFile()
	return strconv.Itoa(sum)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	grid_image := aocutils.CreateImage(grid.Rows, grid.Cols, 5, 5, "twostar")
	grid_image.UseFullColors()
	sum := 0
	valmap := map[string]int{
		".": 1,
		"@": 7,
	}
	gif := aocutils.CreateGIF("twostar", 100)

	for i, line := range lines {
		for j := range line {
			value, _ := grid.GetValue(i, j)
			grid_image.SetZoomedPixel(j, i, valmap[string(rune(value))])
		}
	}
	for {
		changes := make([]aocutils.Key, 0)
		for i, line := range lines {
			for j := range line {
				value, _ := grid.GetValue(i, j)
				count_ := 0
				if rune(value) != '@' {
					continue
				}
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if x == 0 && y == 0 {
							continue
						}
						check, _ := grid.GetValue(i+x, j+y)
						if rune(check) == '@' {
							count_ += 1
						}
					}
				}
				if count_ < 4 {
					grid_image.SetZoomedPixel(j, i, 5)
					changes = append(changes, aocutils.Key{Row: i, Col: j})
					gif.AddFrame(grid_image)
				}
			}
		}
		sum += len(changes)
		for _, change := range changes {
			grid.SetValue(change.Row, change.Col, '.')
			grid_image.SetZoomedPixel(change.Col, change.Row, 1)
			gif.AddFrame(grid_image)
		}
		if len(changes) == 0 {
			break
		}
	}
	gif.WriteGIFToFile()
	return strconv.Itoa(sum)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
