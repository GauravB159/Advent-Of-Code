package main

import (
	"fmt"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func next_day(population *[]int) int {
	var temp int = (*population)[8]
	sum := 0
	for i := 7; i >= 0; i-- {
		temp, (*population)[i] = (*population)[i], temp
		sum += (*population)[i]
	}
	(*population)[6] += temp
	(*population)[8] = temp
	return sum + 2*temp
}

func exponential_growth(line string, days int) string {
	nums := strings.Split(line, ",")
	counts := make([]int, 9)
	for _, num := range nums {
		num, _ := strconv.Atoi(num)
		counts[num] += 1
	}
	var sum int
	for day := 0; day < days; day++ {
		sum = next_day(&counts)
	}
	return strconv.Itoa(sum)
}

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	return exponential_growth(lines[0], 80)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	return exponential_growth(lines[0], 256)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
