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
	count := 0
	for _, line := range lines {
		values := strings.Fields(line)
		differences := make([]int, 0)
		for value := range values[:len(values)-1] {
			next_value, _ := strconv.Atoi(values[value+1])
			curr_value, _ := strconv.Atoi(values[value])
			differences = append(differences, next_value-curr_value)
		}
		sign := 1
		if differences[0] < 0 {
			sign = -1
		}
		safe := true
		for diff := range differences {
			diff_sign := 1
			if differences[diff] < 0 {
				diff_sign = -1
			}
			if !(math.Abs(float64(differences[diff])) >= 1 && math.Abs(float64(differences[diff])) <= 3) {
				safe = false
				break
			}
			if sign != diff_sign {
				safe = false
				break
			}
		}
		if !safe {
			continue
		}
		count += 1
	}
	return strconv.Itoa(count)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	count := 0
	for _, line := range lines {
		values := strings.Fields(line)
		safe := false
		for i := -1; i < len(values); i++ {
			temp_values := make([]string, 0)
			for idx := range values {
				if idx == i {
					continue
				}
				temp_values = append(temp_values, values[idx])
			}
			differences := make([]int, 0)
			for value := range temp_values[:len(temp_values)-1] {
				next_value, _ := strconv.Atoi(temp_values[value+1])
				curr_value, _ := strconv.Atoi(temp_values[value])
				differences = append(differences, next_value-curr_value)
			}
			sign := 1
			if differences[0] < 0 {
				sign = -1
			}
			inner_safe := true
			for diff := range differences {
				diff_sign := 1
				if differences[diff] < 0 {
					diff_sign = -1
				}
				if !(math.Abs(float64(differences[diff])) >= 1 && math.Abs(float64(differences[diff])) <= 3) {
					inner_safe = false
					break
				}
				if sign != diff_sign {
					inner_safe = false
					break
				}
			}
			if inner_safe {
				safe = true
				break
			}
		}
		if safe {
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
