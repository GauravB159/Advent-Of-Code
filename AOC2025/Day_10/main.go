package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

type Machine struct {
	target          int
	transformations []int
}

func find_answer_one(target int, transformations []int, current int, result int, level int) int {
	if current == target {
		return result
	}
	if level >= len(transformations) {
		return 100000000
	}
	first := find_answer_one(target, transformations, current^transformations[level], result+1, level+1)
	second := find_answer_one(target, transformations, current, result, level+1)
	return min(first, second)
}

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := 0
	machines := make([]Machine, 0)
	for _, line := range lines {
		line_split := strings.Split(line, " ")
		target := line_split[0]
		target = target[1 : len(target)-1]
		target_number := 0.0
		for i, value := range target {
			if value == '#' {
				target_number += math.Pow(2, float64(i))
			}
		}
		transformations := make([]int, 0)
		for _, transformation := range line_split[1 : len(line_split)-1] {
			temp_transformation := strings.Split(transformation[1:len(transformation)-1], ",")
			transformation_number := 0.0
			for _, num := range temp_transformation {
				num_, _ := strconv.Atoi(num)
				transformation_number += math.Pow(2, float64(num_))
			}
			transformations = append(transformations, int(transformation_number))
		}

		machines = append(machines, Machine{target: int(target_number), transformations: transformations})
	}
	for _, machine := range machines {
		answer := find_answer_one(machine.target, machine.transformations, 0, 0, 0)
		result += answer
	}
	return strconv.Itoa(result)
}

func check_arrays_equality(first []int, second []int) int {
	if len(first) != len(second) {
		return 1
	}
	equal := 0
	for i := range first {
		if first[i] < second[i] {
			equal = -1
		} else if first[i] > second[i] {
			equal = 1
			break
		}
	}
	return equal
}

func get_key(slice []int) string {
	result := ""
	for _, val := range slice {
		result += strconv.Itoa(val) + ","
	}
	return result
}

func find_answer_two(joltages []int, transformations [][]int, current []int, level int, cache map[string]int) int {
	key := get_key(current)
	if value, exists := cache[key]; exists {
		return value
	}
	equality := check_arrays_equality(current, joltages)
	switch equality {
	case 1:
		cache[key] = 1000000
		return 1000000
	case 0:
		cache[key] = 0
		return 0
	}
	min_ := 1000000000
	for _, transformation := range transformations {
		temp_current := make([]int, len(current))
		copy(temp_current, current)
		for _, num := range transformation {
			temp_current[num] += 1
		}
		result := 1 + find_answer_two(joltages, transformations, temp_current, level+1, cache)
		if result < min_ {
			min_ = result
		}
	}
	cache[key] = min_
	return min_
}

type NewMachine struct {
	transformations [][]int
	joltages        []int
}

/* Only works for the example, not for the actual input. Too slow, need a solver. */
func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := 0
	machines := make([]NewMachine, 0)
	for _, line := range lines {
		line_split := strings.Split(line, " ")
		transformations := make([][]int, 0)
		for _, transformation := range line_split[1 : len(line_split)-1] {
			temp_transformation := make([]int, 0)
			for _, num := range strings.Split(transformation[1:len(transformation)-1], ",") {
				num_int, _ := strconv.Atoi(num)
				temp_transformation = append(temp_transformation, num_int)
			}
			transformations = append(transformations, temp_transformation)
		}
		temp_joltages := line_split[len(line_split)-1]
		joltages := make([]int, 0)
		for _, val := range strings.Split(temp_joltages[1:len(temp_joltages)-1], ",") {
			temp, _ := strconv.Atoi(val)
			joltages = append(joltages, temp)
		}
		machines = append(machines, NewMachine{transformations: transformations, joltages: joltages})
	}
	for _, machine := range machines {
		current := make([]int, len(machine.joltages))
		cache := make(map[string]int)
		answer := find_answer_two(machine.joltages, machine.transformations, current, 0, cache)
		fmt.Println(machine, answer)
		result += answer
	}
	return strconv.Itoa(result)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	// fmt.Println()
	// aocutils.Timer("2 star", twostar, "input.txt")
}
