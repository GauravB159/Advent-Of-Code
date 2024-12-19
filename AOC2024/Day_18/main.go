package main

import (
	"fmt"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := ""
	grid := [71][71]rune{}
	grid_image := aocutils.CreateImage(71, 71, 10, 10, "onestar")
	grid_image.UseFullColors()
	gif := aocutils.CreateGIF("onestar", 40)
	for i := 0; i < 71; i++ {
		for j := 0; j < 71; j++ {
			grid[i][j] = '.'
			grid_image.SetZoomedPixel(i, j, 2)
		}
	}
	grid_image.SetZoomedPixel(0, 0, 1)
	grid_image.SetZoomedPixel(70, 70, 1)
	gif.AddFrame(grid_image)
	count := 0
	num_bytes := 1024
	for i := range num_bytes {
		line := lines[i]
		x, _ := strconv.Atoi(strings.Split(line, ",")[0])
		y, _ := strconv.Atoi(strings.Split(line, ",")[1])
		grid[y][x] = '#'
		gif.AddFrame(grid_image)
		grid_image.SetZoomedPixel(x, y, 5)
		count += 1
	}
	start := aocutils.Key{Row: 0, Col: 0}
	type Node struct {
		position aocutils.Key
		count    int
		path     []aocutils.Key
	}
	queue := make([]Node, 0)
	queue = append(queue, Node{position: start, count: 0, path: make([]aocutils.Key, 0)})
	directions := [4]aocutils.Key{{Row: 1, Col: 0}, {Row: -1, Col: 0}, {Row: 0, Col: 1}, {Row: 0, Col: -1}}
	visited := make(map[aocutils.Key]bool)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if visited[node.position] {
			continue
		}
		visited[node.position] = true
		if node.position.Row < 0 || node.position.Row >= len(grid) {
			continue
		}
		if node.position.Col < 0 || node.position.Col >= len(grid[0]) {
			continue
		}
		if grid[node.position.Row][node.position.Col] == '#' {
			continue
		}
		if node.position.Row == len(grid)-1 && node.position.Col == len(grid[0])-1 {
			result = strconv.Itoa(node.count)
			grid_image.SetZoomedPixel(node.position.Col, node.position.Row, 6)
			for range 40 {
				gif.AddFrame(grid_image)
			}
			for _, point := range node.path {
				grid_image.SetZoomedPixel(point.Col, point.Row, 8)
				gif.AddFrame(grid_image)
			}
			grid_image.SetZoomedPixel(node.position.Col, node.position.Row, 8)
			for range 40 {
				gif.AddFrame(grid_image)
			}
			break
		}
		grid_image.SetZoomedPixel(node.position.Col, node.position.Row, 6)
		gif.AddFrame(grid_image)
		path := append(node.path, node.position)
		for _, direction := range directions {
			var path_copy []aocutils.Key = make([]aocutils.Key, len(path))
			copy(path_copy, path)
			queue = append(queue, Node{position: aocutils.Key{Row: node.position.Row + direction.Row, Col: node.position.Col + direction.Col}, count: node.count + 1, path: path_copy})
		}
	}
	gif.WriteGIFToFile()
	return result
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := [71][71]rune{}
	grid_image := aocutils.CreateImage(71, 71, 10, 10, "twostar")
	grid_image.UseFullColors()
	gif := aocutils.CreateGIF("twostar", 40)
	for i := 0; i < 71; i++ {
		for j := 0; j < 71; j++ {
			grid[i][j] = '.'
			grid_image.SetZoomedPixel(i, j, 2)
		}
	}
	grid_image.SetZoomedPixel(0, 0, 1)
	grid_image.SetZoomedPixel(70, 70, 1)
	lptr := 0
	rptr := len(lines)
	for lptr < rptr {
		idx := (lptr + rptr) / 2
		for i := 0; i < 71; i++ {
			for j := 0; j < 71; j++ {
				grid[i][j] = '.'
				grid_image.SetZoomedPixel(i, j, 2)
				if i%5 == 0 {
					gif.AddFrame(grid_image)
				}
			}
		}
		for i := range idx + 1 {
			line := lines[i]
			x, _ := strconv.Atoi(strings.Split(line, ",")[0])
			y, _ := strconv.Atoi(strings.Split(line, ",")[1])
			grid[y][x] = '#'
			grid_image.SetZoomedPixel(x, y, 5)
		}
		gif.AddFrame(grid_image)
		start := aocutils.Key{Row: 0, Col: 0}
		type Node struct {
			position aocutils.Key
			count    int
			path     []aocutils.Key
		}
		queue := make([]Node, 0)
		queue = append(queue, Node{position: start, count: 0, path: make([]aocutils.Key, 0)})
		directions := [4]aocutils.Key{{Row: 1, Col: 0}, {Row: -1, Col: 0}, {Row: 0, Col: 1}, {Row: 0, Col: -1}}
		visited := make(map[aocutils.Key]bool)
		reached := false
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			if visited[node.position] {
				continue
			}
			visited[node.position] = true
			if node.position.Row < 0 || node.position.Row >= len(grid) {
				continue
			}
			if node.position.Col < 0 || node.position.Col >= len(grid[0]) {
				continue
			}
			if grid[node.position.Row][node.position.Col] == '#' {
				continue
			}
			if node.position.Row == len(grid)-1 && node.position.Col == len(grid[0])-1 {
				reached = true
				grid_image.SetZoomedPixel(node.position.Col, node.position.Row, 6)
				for range 40 {
					gif.AddFrame(grid_image)
				}
				for _, point := range node.path {
					grid_image.SetZoomedPixel(point.Col, point.Row, 8)
					gif.AddFrame(grid_image)
				}
				grid_image.SetZoomedPixel(node.position.Col, node.position.Row, 8)
				break
			}
			grid_image.SetZoomedPixel(node.position.Col, node.position.Row, 6)
			gif.AddFrame(grid_image)
			path := append(node.path, node.position)
			for _, direction := range directions {
				var path_copy []aocutils.Key = make([]aocutils.Key, len(path))
				copy(path_copy, path)
				queue = append(queue, Node{position: aocutils.Key{Row: node.position.Row + direction.Row, Col: node.position.Col + direction.Col}, count: node.count + 1, path: path_copy})
			}
		}
		if !reached {
			rptr = idx - 1
			for range 2000 {
				gif.AddFrame(grid_image)
			}
		} else {
			lptr = idx + 1
			for range 2000 {
				gif.AddFrame(grid_image)
			}
		}
	}
	gif.WriteGIFToFile()
	return lines[lptr+1]
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
