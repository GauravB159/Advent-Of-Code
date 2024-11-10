package main

import (
	"fmt"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func part_one(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	final_count := 0
	grid_image := aocutils.CreateImage(grid.Rows, grid.Cols, 40, "part_one")
	grid_image.UsePaletteWideReds(true)
	gif := aocutils.CreateGIF("part_one", 1)
	for key := range grid.Data {
		if grid.Data[key] == '.' {
			grid.Data[key] = 0
		} else {
			final_count += 1
			grid.Data[key] = 1
		}
	}
	for i := 2; ; i++ {
		count := 0
		changes := make([]aocutils.Key, 0)
		for key := range grid.Data {
			if i-grid.Data[aocutils.Key{Row: key.Row, Col: key.Col + 1}] > 1 {
				continue
			}
			if i-grid.Data[aocutils.Key{Row: key.Row, Col: key.Col - 1}] > 1 {
				continue
			}
			if i-grid.Data[aocutils.Key{Row: key.Row + 1, Col: key.Col}] > 1 {
				continue
			}
			if i-grid.Data[aocutils.Key{Row: key.Row - 1, Col: key.Col}] > 1 {
				continue
			}
			count += 1
			changes = append(changes, key)
		}
		for _, change := range changes {
			grid.Data[change] = i
			grid_image.SetZoomedPixel(change.Col, change.Row, i)
			gif.AddFrame(grid_image)
		}
		final_count += count
		if count == 0 {
			break
		}
	}
	gif.WriteGIFToFile()
	return strconv.Itoa(final_count)
}

func part_two(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	final_count := 0
	grid_image := aocutils.CreateImage(grid.Rows, grid.Cols, 40, "part_two")
	grid_image.UsePaletteWideReds(true)
	gif := aocutils.CreateGIF("part_two", 10)
	for key := range grid.Data {
		if grid.Data[key] == '.' {
			grid.Data[key] = 0
		} else {
			final_count += 1
			grid.Data[key] = 1
		}
	}
	for i := 2; ; i++ {
		count := 0
		changes := make([]aocutils.Key, 0)
		for key := range grid.Data {
			if i-grid.Data[aocutils.Key{Row: key.Row, Col: key.Col + 1}] > 1 {
				continue
			}
			if i-grid.Data[aocutils.Key{Row: key.Row, Col: key.Col - 1}] > 1 {
				continue
			}
			if i-grid.Data[aocutils.Key{Row: key.Row + 1, Col: key.Col}] > 1 {
				continue
			}
			if i-grid.Data[aocutils.Key{Row: key.Row - 1, Col: key.Col}] > 1 {
				continue
			}
			count += 1
			changes = append(changes, key)
		}
		for _, change := range changes {
			grid.Data[change] = i
			grid_image.SetZoomedPixel(change.Col, change.Row, i)
			gif.AddFrame(grid_image)
		}
		final_count += count
		if count == 0 {
			break
		}
	}
	gif.WriteGIFToFile()
	return strconv.Itoa(final_count)
}

func part_three(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	final_count := 0
	grid_image := aocutils.CreateImage(grid.Rows, grid.Cols, 40, "part_three")
	grid_image.UsePaletteReds(true)
	gif := aocutils.CreateGIF("part_three", 200)
	for key := range grid.Data {
		if grid.Data[key] == '.' {
			grid.Data[key] = 0
		} else {
			final_count += 1
			grid.Data[key] = 1
		}
	}
	for i := 2; ; i++ {
		count := 0
		changes := make([]aocutils.Key, 0)
		for key := range grid.Data {
			if grid.Data[key] == 0 {
				continue
			}
			if i-grid.Data[aocutils.Key{Row: key.Row, Col: key.Col + 1}] > 1 {
				continue
			}
			if i-grid.Data[aocutils.Key{Row: key.Row, Col: key.Col - 1}] > 1 {
				continue
			}
			if i-grid.Data[aocutils.Key{Row: key.Row + 1, Col: key.Col}] > 1 {
				continue
			}
			if i-grid.Data[aocutils.Key{Row: key.Row - 1, Col: key.Col}] > 1 {
				continue
			}
			if i-grid.Data[aocutils.Key{Row: key.Row - 1, Col: key.Col - 1}] > 1 {
				continue
			}
			if i-grid.Data[aocutils.Key{Row: key.Row - 1, Col: key.Col + 1}] > 1 {
				continue
			}
			if i-grid.Data[aocutils.Key{Row: key.Row + 1, Col: key.Col - 1}] > 1 {
				continue
			}
			if i-grid.Data[aocutils.Key{Row: key.Row + 1, Col: key.Col + 1}] > 1 {
				continue
			}
			count += 1
			changes = append(changes, key)
		}
		for _, change := range changes {
			grid.Data[change] = i
			grid_image.SetZoomedPixel(change.Col, change.Row, i)
			gif.AddFrame(grid_image)
		}
		final_count += count
		if count == 0 {
			break
		}
	}
	gif.WriteGIFToFile()
	return strconv.Itoa(final_count)
}

func main() {
	aocutils.Timer("Part 1", part_one, "input_one.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("Part 2", part_two, "input_two.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("Part 3", part_three, "input_three.txt")
}
