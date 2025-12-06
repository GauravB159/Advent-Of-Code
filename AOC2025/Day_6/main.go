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
	operands := make([][]int, 0)
	for _, line := range lines[:len(lines)-1] {
		temp := strings.Split(line, " ")
		values := make([]int, 0)
		for _, value := range temp {
			if value == "" {
				continue
			}
			int_value, _ := strconv.Atoi(value)
			values = append(values, int_value)
		}
		operands = append(operands, values)
	}
	operators := strings.ReplaceAll(lines[len(lines)-1], " ", "")
	for i := 0; i < len(operands[0]); i++ {
		answer := 0
		if operators[i] == '*' {
			answer = 1
		}
		for j := 0; j < len(operands); j++ {
			if operators[i] == '*' {
				answer *= operands[j][i]
			} else {
				answer += operands[j][i]
			}
		}
		result += answer
	}
	return strconv.Itoa(result)
}

type Range struct {
	min      int
	max      int
	operator string
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := 0
	ranges := make([]Range, 0)
	curr_range := Range{
		min: 0,
		max: -1,
	}
	for i, operator := range lines[len(lines)-1] {
		if i == 0 {
			curr_range.operator = string(operator)
			continue
		}
		if string(operator) != " " {
			curr_range.max = i - 1
			ranges = append(ranges, curr_range)
			curr_range = Range{
				min:      i,
				max:      -1,
				operator: string(operator),
			}
		}
	}
	max_line_length := 0
	for _, line := range lines {
		if len(line) > max_line_length {
			max_line_length = len(line)
		}
	}
	ranges = append(ranges, curr_range)
	for _, range_ := range ranges {
		max_ := range_.max
		if max_ == -1 {
			max_ = max_line_length
		}
		answer := 0
		if range_.operator == "*" {
			answer = 1
		}
		for j := range_.min; j < max_; j++ {
			num := ""
			for _, line := range lines[:len(lines)-1] {
				if line[j] == ' ' {
					continue
				}
				num += string(line[j])
			}
			actual_num, _ := strconv.Atoi(num)
			if range_.operator == "*" {
				answer *= actual_num
			} else {
				answer += actual_num
			}
		}
		result += answer
	}
	return strconv.Itoa(result)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
