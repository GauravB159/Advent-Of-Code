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
	count := 0
	for _, line := range lines {
		check := strings.Split(line, "mul(")
		for val := range check {
			inner := strings.Split(check[val], ",")
			if leftVal, leftErr := strconv.Atoi(inner[0]); leftErr == nil {
				right := strings.Split(inner[1], ")")
				if rightVal, rightErr := strconv.Atoi(right[0]); rightErr == nil && len(right) > 1 {
					result += leftVal * rightVal
					count += 1
				}
			}
		}
	}
	return strconv.Itoa(result)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := 0
	activate := true
	for _, line := range lines {
		check := strings.Split(line, "mul(")
		for val := range check {
			currActivate := activate
			if strings.Contains(check[val], "do()") {
				currActivate = true
			} else if strings.Contains(check[val], "don't()") {
				currActivate = false
			}
			if !activate {
				activate = currActivate
				continue
			}
			activate = currActivate
			inner := strings.Split(check[val], ",")
			if leftVal, leftErr := strconv.Atoi(inner[0]); leftErr == nil {
				right := strings.Split(inner[1], ")")
				if rightVal, rightErr := strconv.Atoi(right[0]); rightErr == nil && len(right) > 1 {
					result += leftVal * rightVal
				}
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
