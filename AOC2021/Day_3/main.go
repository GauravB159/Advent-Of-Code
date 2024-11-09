package main

import (
	"fmt"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	counts := make([]int, len(lines[0]))
	for _, line := range lines {
		for index, char := range line {
			if char == '1' {
				counts[index] += 1
			}
		}
	}
	gamma_rate := ""
	epsilon_rate := ""
	for _, value := range counts {
		if value > len(lines)/2 {
			gamma_rate += "1"
			epsilon_rate += "0"
		} else {
			gamma_rate += "0"
			epsilon_rate += "1"
		}
	}
	decimal_gamma, _ := strconv.ParseInt(gamma_rate, 2, 0)
	decimal_epsilon, _ := strconv.ParseInt(epsilon_rate, 2, 0)
	return strconv.Itoa(int(decimal_gamma * decimal_epsilon))
}

func Intersection(a []string, b []string) []string {
	aMap := make(map[string]bool)
	for _, value := range a {
		aMap[value] = true
	}
	answer := make([]string, 0)
	for _, value := range b {
		if aMap[value] {
			answer = append(answer, value)
		}
	}
	return answer
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	bitmap := make(map[string][]string)
	for _, line := range lines {
		for index, char := range line {
			key := string(char) + strconv.Itoa(index)
			if _, exists := bitmap[key]; !exists {
				bitmap[key] = make([]string, 0)
			}
			bitmap[key] = append(bitmap[key], line)
		}
	}
	oxygen := lines
	co2 := lines
	for i := 0; i < len(lines[0]); i++ {
		zero_key := "0" + strconv.Itoa(i)
		one_key := "1" + strconv.Itoa(i)
		if len(oxygen) > 1 {
			oxygen_one_intersection := Intersection(oxygen, bitmap[one_key])
			oxygen_zero_intersection := Intersection(oxygen, bitmap[zero_key])
			if len(oxygen_one_intersection) >= len(oxygen_zero_intersection) {
				oxygen = oxygen_one_intersection
			} else {
				oxygen = oxygen_zero_intersection
			}
		}
		if len(co2) > 1 {
			co2_one_intersection := Intersection(co2, bitmap[zero_key])
			co2_zero_intersection := Intersection(co2, bitmap[one_key])
			if len(co2_one_intersection) <= len(co2_zero_intersection) {
				co2 = co2_one_intersection
			} else {
				co2 = co2_zero_intersection
			}
		}
	}
	decimal_oxygen, _ := strconv.ParseInt(oxygen[0], 2, 0)
	decimal_co2, _ := strconv.ParseInt(co2[0], 2, 0)
	return strconv.Itoa(int(decimal_oxygen * decimal_co2))
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
