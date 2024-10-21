package main

import (
	"fmt"
	"sort"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	brackets := map[rune]rune{'(': ')', '[': ']', '<': '>', '{': '}'}
	scores := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	score := 0
	for _, line := range lines {
		var stack aocutils.Stack[rune]
		for _, char := range line {
			if value, exists := brackets[char]; exists {
				stack.Push(value)
			} else if !stack.IsEmpty() {
				match, _ := stack.Pop()
				if match != char {
					score += scores[char]
					break
				}
			}
		}
	}
	return strconv.Itoa(score)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	brackets := map[rune]rune{'(': ')', '[': ']', '<': '>', '{': '}'}
	scores := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}
	autocomplete_scores := make([]int, 0)
	for _, line := range lines {
		corrupted := false
		var stack aocutils.Stack[rune]
		for _, char := range line {
			if value, exists := brackets[char]; exists {
				stack.Push(value)
			} else if !stack.IsEmpty() {
				match, _ := stack.Pop()
				if match != char {
					corrupted = true
				}
			}
		}
		if !corrupted {
			score := 0
			for !stack.IsEmpty() {
				value, _ := stack.Pop()
				score *= 5
				score += scores[value]
			}
			autocomplete_scores = append(autocomplete_scores, score)
		}
	}
	sort.Ints(autocomplete_scores)
	return strconv.Itoa(autocomplete_scores[int(len(autocomplete_scores)/2)])
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
