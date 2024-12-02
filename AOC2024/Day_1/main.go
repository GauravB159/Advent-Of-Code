package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := 0
	left := make([]int, 0)
	right := make([]int, 0)
	for _, line := range lines {
		values := strings.Fields(line)
		leftInt, _ := strconv.Atoi(values[0])
		rightInt, _ := strconv.Atoi(values[1])
		left = append(left, leftInt)
		right = append(right, rightInt)
	}
	sortedLeft := sort.IntSlice(left)
	sortedRight := sort.IntSlice(right)
	sortedLeft.Sort()
	sortedRight.Sort()
	for i := range sortedLeft {
		result += int(math.Abs(float64(sortedLeft[i]) - float64(sortedRight[i])))
	}
	return strconv.Itoa(result)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := 0
	left := make([]int, 0)
	right := make([]int, 0)
	for _, line := range lines {
		values := strings.Fields(line)
		leftInt, _ := strconv.Atoi(values[0])
		rightInt, _ := strconv.Atoi(values[1])
		left = append(left, leftInt)
		right = append(right, rightInt)
	}
	sortedLeft := sort.IntSlice(left)
	sortedRight := sort.IntSlice(right)
	sortedLeft.Sort()
	sortedRight.Sort()
	for i := range sortedLeft {
		number := sortedLeft[i]
		for j := range sortedRight {
			if sortedRight[j] == number {
				result += sortedRight[j]
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
