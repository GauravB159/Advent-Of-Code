package main

import (
	"fmt"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	grid_image := aocutils.CreateImage(grid.GetNumRows(), grid.GetNumCols(), 10, "onestar")
	grid_image.UsePaletteColors()
	gif := aocutils.CreateGIF("onestar", 40)
	var start aocutils.Key
	direction := aocutils.Key{Row: -1, Col: 0}
	for key := range grid.Data {
		if grid.Data[key] == '^' {
			grid_image.SetZoomedPixel(key.Col, key.Row, 4)
			start = key
		} else if grid.Data[key] == '#' {
			grid_image.SetZoomedPixel(key.Col, key.Row, 5)
		} else {
			grid_image.SetZoomedPixel(key.Col, key.Row, 1)
		}
	}
	visited := make(map[aocutils.Key]bool, 0)
	visited[start] = true
	for true {
		nextPosition := aocutils.Key{Row: start.Row + direction.Row, Col: start.Col + direction.Col}
		gif.AddFrame(grid_image)
		if val, exists := grid.Data[nextPosition]; exists {
			if val == '#' {
				if direction.Row == -1 && direction.Col == 0 {
					direction = aocutils.Key{Row: 0, Col: 1}
				} else if direction.Row == 0 && direction.Col == 1 {
					direction = aocutils.Key{Row: 1, Col: 0}
				} else if direction.Row == 1 && direction.Col == 0 {
					direction = aocutils.Key{Row: 0, Col: -1}
				} else if direction.Row == 0 && direction.Col == -1 {
					direction = aocutils.Key{Row: -1, Col: 0}
				}
			} else {
				grid_image.SetZoomedPixel(nextPosition.Col, nextPosition.Row, 3)
				visited[nextPosition] = true
				start = nextPosition
			}
		} else {
			break
		}
	}
	gif.AddFrame(grid_image)
	gif.WriteGIFToFile()
	return strconv.Itoa(len(visited))
}

type uniqueVisit struct {
	position  aocutils.Key
	direction aocutils.Key
}

func checkLoop(grid aocutils.Grid, start aocutils.Key, gif *aocutils.GIF, grid_image aocutils.Image) bool {
	direction := aocutils.Key{Row: -1, Col: 0}
	visited := make(map[uniqueVisit]bool, 0)
	for true {
		nextPosition := aocutils.Key{Row: start.Row + direction.Row, Col: start.Col + direction.Col}
		if grid.Data[nextPosition] == '#' && visited[uniqueVisit{position: nextPosition, direction: direction}] {
			for range 5 {
				gif.AddFrame(grid_image)
			}
			return true
		}
		if val, exists := grid.Data[nextPosition]; exists {
			if val == '#' {
				visited[uniqueVisit{position: nextPosition, direction: direction}] = true
				if direction.Row == -1 && direction.Col == 0 {
					direction = aocutils.Key{Row: 0, Col: 1}
				} else if direction.Row == 0 && direction.Col == 1 {
					direction = aocutils.Key{Row: 1, Col: 0}
				} else if direction.Row == 1 && direction.Col == 0 {
					direction = aocutils.Key{Row: 0, Col: -1}
				} else if direction.Row == 0 && direction.Col == -1 {
					direction = aocutils.Key{Row: -1, Col: 0}
				}
			} else {
				if (gif.Framecount+5)%gif.Frameskip == 0 {
					grid_image.SetZoomedPixel(nextPosition.Col, nextPosition.Row, 3)
				}
				visited[uniqueVisit{position: nextPosition, direction: direction}] = true
				start = nextPosition
			}
		} else {
			break
		}
	}
	return false
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	grid_image := aocutils.CreateImage(grid.GetNumRows(), grid.GetNumCols(), 10, "twostar")
	grid_image.UsePaletteColors()
	gif := aocutils.CreateGIF("twostar", 25)
	var start aocutils.Key
	for key := range grid.Data {
		if grid.Data[key] == '^' {
			grid_image.SetZoomedPixel(key.Col, key.Row, 4)
			start = key
		} else if grid.Data[key] == '#' {
			grid_image.SetZoomedPixel(key.Col, key.Row, 5)
		} else {
			grid_image.SetZoomedPixel(key.Col, key.Row, 1)
		}
	}
	var prevKey aocutils.Key = aocutils.Key{Row: -1, Col: -1}
	checkKey := aocutils.Key{Row: -1, Col: -1}
	count := 0
	progress := 0
	for key := range grid.Data {
		progress += 1
		fmt.Println(progress, len(grid.Data))
		if grid.Data[key] == '#' {
			continue
		}
		if prevKey != checkKey {
			grid.Data[prevKey] = '.'
		}
		grid.Data[key] = '#'
		prevKey = key
		if checkLoop(grid, start, &gif, grid_image.Clone()) {
			count += 1
		}
	}
	gif.WriteGIFToFile()
	return strconv.Itoa(count)
}

func main() {
	// aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
