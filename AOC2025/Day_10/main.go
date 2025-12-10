package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
	"github.com/draffensperger/golp"
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

type NewMachine struct {
	transformations [][]int
	joltages        []int
}

func solveMachine(transformations [][]int, joltages []int) int {
	lp := golp.NewLP(len(joltages), len(transformations))
	objective := make([]float64, 0)
	for i := range len(transformations) {
		lp.SetInt(i, true)
		lp.AddConstraintSparse([]golp.Entry{{Col: i, Val: 1.0}}, golp.GE, 0.0)
		objective = append(objective, 1.0)
	}
	for i, joltage := range joltages {
		constraint := make([]golp.Entry, 0)
		for j, transformation := range transformations {
			for _, inner := range transformation {
				if i == inner {
					constraint = append(constraint, golp.Entry{Col: j, Val: 1.0})
				}
			}
		}
		lp.AddConstraintSparse(constraint, golp.EQ, float64(joltage))
	}
	lp.SetObjFn(objective)
	lp.Solve()
	return int(math.Round(lp.Objective()))
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
		result += solveMachine(machine.transformations, machine.joltages)
	}
	return strconv.Itoa(result)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
