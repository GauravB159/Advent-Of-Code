package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	overlaps := make(map[string]int, len(lines)*len(lines))
	for _, line := range lines {
		points := strings.Split(line, " -> ")
		start := strings.Split(points[0], ",")
		end := strings.Split(points[1], ",")
		start_x, _ := strconv.Atoi(start[0])
		start_y, _ := strconv.Atoi(start[1])
		end_x, _ := strconv.Atoi(end[0])
		end_y, _ := strconv.Atoi(end[1])
		var x_direction int = 0
		var y_direction int = 0
		if end_x != start_x {
			x_direction = int(float64(end_x-start_x) / math.Abs(float64(end_x)-float64(start_x)))
		}
		if end_y != start_y {
			y_direction = int(float64(end_y-start_y) / math.Abs(float64(end_y)-float64(start_y)))
		}
		if x_direction != 0 && y_direction != 0 {
			continue
		}
		for x, y := start_x, start_y; x != end_x+x_direction || y != end_y+y_direction; x, y = x+x_direction, y+y_direction {
			if _, exist := overlaps[fmt.Sprintf("%d-%d", x, y)]; !exist {
				overlaps[fmt.Sprintf("%d-%d", x, y)] = 0
			}
			overlaps[fmt.Sprintf("%d-%d", x, y)] += 1
		}
	}
	count := 0
	for _, value := range overlaps {
		if value > 1 {
			count += 1
		}
	}
	return strconv.Itoa(count)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	overlaps := make(map[string]int, len(lines)*len(lines))
	for _, line := range lines {
		points := strings.Split(line, " -> ")
		start := strings.Split(points[0], ",")
		end := strings.Split(points[1], ",")
		start_x, _ := strconv.Atoi(start[0])
		start_y, _ := strconv.Atoi(start[1])
		end_x, _ := strconv.Atoi(end[0])
		end_y, _ := strconv.Atoi(end[1])
		var x_direction int = 0
		var y_direction int = 0
		if end_x != start_x {
			x_direction = int(float64(end_x-start_x) / math.Abs(float64(end_x)-float64(start_x)))
		}
		if end_y != start_y {
			y_direction = int(float64(end_y-start_y) / math.Abs(float64(end_y)-float64(start_y)))
		}
		for x, y := start_x, start_y; x != end_x+x_direction || y != end_y+y_direction; x, y = x+x_direction, y+y_direction {
			overlaps[fmt.Sprintf("%d-%d", x, y)] += 1
		}
	}
	count := 0
	for _, value := range overlaps {
		if value > 1 {
			count += 1
		}
	}
	return strconv.Itoa(count)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
