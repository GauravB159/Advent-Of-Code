package main

import (
	"fmt"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	sum_ := 0
	for _, line := range lines {
		max_ := 0
		for i := range line {
			for j := range line {
				if j <= i {
					continue
				}
				check, _ := strconv.Atoi(string(line[i]) + string(line[j]))
				if check > max_ {
					max_ = check
				}
			}
		}
		sum_ += max_
	}
	return strconv.Itoa(sum_)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	sum_ := 0
	N := 12
	for _, line := range lines {
		max_ := ""
		for i := range line {
			if len(max_) == 0 {
				max_ += string(line[i])
				continue
			}
			remaining := len(line) - i
			if len(max_)+remaining == N {
				max_ += string(line[i:])
				break
			}
			candidate, _ := strconv.Atoi(string(line[i]))
			found := false
			for j := range max_ {
				check, _ := strconv.Atoi(string(max_[j]))
				if candidate > check && remaining+j >= N {
					found = true
					max_ = max_[:j] + string(line[i])
					break
				}
			}
			if !found && len(max_) < N {
				max_ += string(line[i])
			}
		}
		max_num, _ := strconv.Atoi(max_)
		sum_ += max_num
	}
	return strconv.Itoa(sum_)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
