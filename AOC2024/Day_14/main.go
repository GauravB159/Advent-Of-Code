package main

import (
	"fmt"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

type Robot struct {
	px int
	py int
	vx int
	vy int
}

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := make([]Robot, 0)
	for _, line := range lines {
		pv := strings.Fields(line)
		pvals := strings.Split(strings.Split(pv[0], "=")[1], ",")
		vvals := strings.Split(strings.Split(pv[1], "=")[1], ",")
		px, _ := strconv.Atoi(pvals[0])
		py, _ := strconv.Atoi(pvals[1])
		vx, _ := strconv.Atoi(vvals[0])
		vy, _ := strconv.Atoi(vvals[1])
		robot := Robot{
			px: px,
			py: py,
			vx: vx,
			vy: vy,
		}
		grid = append(grid, robot)
	}
	seconds := 100
	width := 101
	height := 103
	final_grid := make(map[Robot]int, 0)
	for _, robot := range grid {
		next_px := (robot.px + robot.vx*seconds) % width
		if next_px < 0 {
			next_px = width + next_px
		}
		next_py := (robot.py + robot.vy*seconds) % height
		if next_py < 0 {
			next_py = height + next_py
		}
		final_robot := Robot{
			px: next_px,
			py: next_py,
		}
		final_grid[final_robot] += 1
	}
	counts := [4]int{0, 0, 0, 0}
	for i := range height {
		for j := range width {
			if j == (width-1)/2 {
				fmt.Print(" ")
				continue
			}
			if i == (height-1)/2 {
				fmt.Print(" ")
				continue
			}
			value := final_grid[Robot{
				px: j,
				py: i,
			}]
			if value > 0 {
				if float32(j) < float32(width)/2 && float32(i) < float32(height)/2 {
					counts[0] += value
				}
				if float32(j) < float32(width)/2 && float32(i) > float32(height)/2 {
					counts[1] += value
				}
				if float32(j) > float32(width)/2 && float32(i) < float32(height)/2 {
					counts[2] += value
				}
				if float32(j) > float32(width)/2 && float32(i) > float32(height)/2 {
					counts[3] += value
				}
			}
			fmt.Print(value)
		}
		fmt.Println()
	}
	fmt.Println(counts)
	mult := 1
	for _, val := range counts {
		mult *= val
	}
	return strconv.Itoa(mult)
}

func biggest_connected_components(node Robot, grid map[Robot]int, visited *map[Robot]bool, count *int) {
	if (*visited)[node] || grid[node] == 0 {
		return
	}

	(*visited)[node] = true
	(*count) += 1
	biggest_connected_components(Robot{px: node.px + 1, py: node.py}, grid, visited, count)
	biggest_connected_components(Robot{px: node.px - 1, py: node.py}, grid, visited, count)
	biggest_connected_components(Robot{px: node.px, py: node.py + 1}, grid, visited, count)
	biggest_connected_components(Robot{px: node.px, py: node.py - 1}, grid, visited, count)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := make([]Robot, 0)
	for _, line := range lines {
		pv := strings.Fields(line)
		pvals := strings.Split(strings.Split(pv[0], "=")[1], ",")
		vvals := strings.Split(strings.Split(pv[1], "=")[1], ",")
		px, _ := strconv.Atoi(pvals[0])
		py, _ := strconv.Atoi(pvals[1])
		vx, _ := strconv.Atoi(vvals[0])
		vy, _ := strconv.Atoi(vvals[1])
		robot := Robot{
			px: px,
			py: py,
			vx: vx,
			vy: vy,
		}
		grid = append(grid, robot)
	}
	seconds := 10000
	width := 101
	height := 103
	gif := aocutils.CreateGIF("twostar", 10)
	for second := range seconds {
		final_grid := make(map[Robot]int, 0)
		visited := make(map[Robot]bool)
		for _, robot := range grid {
			next_px := (robot.px + robot.vx*second) % width
			if next_px < 0 {
				next_px = width + next_px
			}
			next_py := (robot.py + robot.vy*second) % height
			if next_py < 0 {
				next_py = height + next_py
			}
			final_robot := Robot{
				px: next_px,
				py: next_py,
			}
			final_grid[final_robot] += 1
			visited[final_robot] = false
		}
		for node := range visited {
			if visited[node] {
				continue
			}
			count := 0
			biggest_connected_components(node, final_grid, &visited, &count)
			if count > 10 {
				grid_image := aocutils.CreateImage(height, width, 10, "twostar_frame_"+strconv.Itoa(second))
				grid_image.UsePaletteColors()
				for node := range final_grid {
					grid_image.SetZoomedPixel(node.py, node.px, final_grid[node]+1)
				}
				for range count {
					gif.AddFrame(grid_image)
				}
				break
			}
		}
	}
	gif.WriteGIFToFile()
	return ""
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
