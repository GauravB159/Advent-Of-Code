package main

import (
	"fmt"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := 0
	limit := 100
	for i := 0; i < len(lines); i += 4 {
		avals := strings.Split(lines[i], "Button A: ")[1]
		ax := strings.Split(avals, ", ")[0]
		ay := strings.Split(avals, ", ")[1]
		ax = strings.Split(ax, "+")[1]
		ay = strings.Split(ay, "+")[1]
		ax_int, _ := strconv.Atoi(ax)
		ay_int, _ := strconv.Atoi(ay)

		bvals := strings.Split(lines[i+1], "Button B: ")[1]
		bx := strings.Split(bvals, ", ")[0]
		by := strings.Split(bvals, ", ")[1]
		bx = strings.Split(bx, "+")[1]
		by = strings.Split(by, "+")[1]
		bx_int, _ := strconv.Atoi(bx)
		by_int, _ := strconv.Atoi(by)

		pvals := strings.Split(lines[i+2], "Prize: ")[1]
		px := strings.Split(pvals, ", ")[0]
		py := strings.Split(pvals, ", ")[1]
		px = strings.Split(px, "=")[1]
		py = strings.Split(py, "=")[1]
		px_int, _ := strconv.Atoi(px)
		py_int, _ := strconv.Atoi(py)
		cost := 0
		for a := 0; a < limit; a++ {
			if a*ax_int > px_int {
				break
			}
			if a*ay_int > py_int {
				break
			}
			for b := 0; b < limit; b++ {
				x := a*ax_int + b*bx_int
				if x > px_int {
					break
				}
				y := a*ay_int + b*by_int
				if y > py_int {
					break
				}
				if x == px_int && y == py_int {
					new_cost := a*3 + b
					if new_cost < cost || cost == 0 {
						cost = new_cost
					}
				}
			}
		}
		result += cost
	}
	return strconv.Itoa(result)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := 0
	for i := 0; i < len(lines); i += 4 {
		avals := strings.Split(lines[i], "Button A: ")[1]
		ax := strings.Split(avals, ", ")[0]
		ay := strings.Split(avals, ", ")[1]
		ax = strings.Split(ax, "+")[1]
		ay = strings.Split(ay, "+")[1]
		ax_int, _ := strconv.Atoi(ax)
		ay_int, _ := strconv.Atoi(ay)

		bvals := strings.Split(lines[i+1], "Button B: ")[1]
		bx := strings.Split(bvals, ", ")[0]
		by := strings.Split(bvals, ", ")[1]
		bx = strings.Split(bx, "+")[1]
		by = strings.Split(by, "+")[1]
		bx_int, _ := strconv.Atoi(bx)
		by_int, _ := strconv.Atoi(by)

		pvals := strings.Split(lines[i+2], "Prize: ")[1]
		px := strings.Split(pvals, ", ")[0]
		py := strings.Split(pvals, ", ")[1]
		px = strings.Split(px, "=")[1]
		py = strings.Split(py, "=")[1]
		px_int, _ := strconv.Atoi(px)
		py_int, _ := strconv.Atoi(py)
		px_int += 10000000000000
		py_int += 10000000000000
		temp_a := float64((px_int*by_int - py_int*bx_int)) / float64((ax_int*by_int - bx_int*ay_int))
		if float64(int(temp_a))-temp_a == 0 {
			temp_b := float64(px_int-ax_int*int(temp_a)) / float64(bx_int)
			if float64(int(temp_b))-temp_b == 0 {
				result += (int(temp_a)*3 + int(temp_b))
			}
		}
	}
	return strconv.Itoa(result)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
