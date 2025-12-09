package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

type Point struct {
	x int
	y int
}

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	points := make([]Point, 0)
	for _, line := range lines {
		temp := strings.Split(line, ",")
		x, _ := strconv.Atoi(temp[0])
		y, _ := strconv.Atoi(temp[1])
		points = append(points, Point{x: x, y: y})
	}
	max_ := 0
	for i, point_one := range points {
		for j, point_two := range points {
			if j <= i {
				continue
			}
			x_diff := math.Abs(float64(point_two.x-point_one.x)) + 1
			y_diff := math.Abs(float64(point_two.y-point_one.y)) + 1
			area := int(x_diff * y_diff)
			if area > max_ {
				max_ = area
			}
		}
	}
	return strconv.Itoa(max_)
}

type Line struct {
	first  Point
	second Point
}

/* Reference: https://en.wikipedia.org/wiki/Line–line_intersection */
func CheckIntersection(line Line, polygon []Line) bool {
	x1 := line.first.x
	x2 := line.second.x
	y1 := line.first.y
	y2 := line.second.y
	if x1 > x2 || y1 > y2 {
		x2 = line.first.x
		x1 = line.second.x
		y2 = line.first.y
		y1 = line.second.y
	}
	for _, inner_line := range polygon {
		x3 := inner_line.first.x
		x4 := inner_line.second.x
		y3 := inner_line.first.y
		y4 := inner_line.second.y
		if x3 > x4 || y3 > y4 {
			x4 = inner_line.first.x
			x3 = inner_line.second.x
			y4 = inner_line.first.y
			y3 = inner_line.second.y
		}
		t_N := ((x1 - x3) * (y3 - y4)) - (y1-y3)*(x3-x4)
		D := ((x1 - x2) * (y3 - y4)) - ((y1 - y2) * (x3 - x4))
		u_N := (y1-y2)*(x1-x3) - ((x1 - x2) * (y1 - y3))
		if D == 0 {
			continue
		}
		t := float64(t_N) / float64(D)
		u := float64(u_N) / float64(D)
		if t > 0 && t < 1 && u > 0 && u < 1 {
			return true
		}
	}
	return false
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	points := make([]Point, 0)
	max_x := 0
	max_y := 0
	for _, line := range lines {
		temp := strings.Split(line, ",")
		x, _ := strconv.Atoi(temp[0])
		y, _ := strconv.Atoi(temp[1])
		if x > max_x {
			max_x = x
		}
		if y > max_y {
			max_y = y
		}
		points = append(points, Point{x: x, y: y})
	}
	polygon_lines := make([]Line, 0)
	for i, point := range points {
		next_point := points[(i+1)%len(points)]
		polygon_lines = append(polygon_lines, Line{first: point, second: next_point})
	}

	max_ := 0
	for i, point_one := range points {
		for j, point_two := range points {
			if j <= i {
				continue
			}
			point_three := Point{x: point_one.x, y: point_two.y}
			point_four := Point{x: point_two.x, y: point_one.y}
			if CheckIntersection(Line{first: point_one, second: point_three}, polygon_lines) || CheckIntersection(Line{first: point_two, second: point_three}, polygon_lines) || CheckIntersection(Line{first: point_one, second: point_four}, polygon_lines) || CheckIntersection(Line{first: point_two, second: point_four}, polygon_lines) || CheckIntersection(Line{first: point_one, second: point_two}, polygon_lines) || CheckIntersection(Line{first: point_three, second: point_four}, polygon_lines) {
				continue
			}

			x_diff := math.Abs(float64(point_two.x-point_one.x)) + 1
			y_diff := math.Abs(float64(point_two.y-point_one.y)) + 1
			area := int(x_diff * y_diff)
			if area > max_ {
				max_ = area
			}
		}
	}

	return strconv.Itoa(max_)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
