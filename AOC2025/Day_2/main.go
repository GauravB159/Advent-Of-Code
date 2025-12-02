package main

import (
	"fmt"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	sum_ := 0
	for _, line := range strings.Split(lines[0], ",") {
		ranges := strings.Split(line, "-")
		before, _ := strconv.Atoi(ranges[0])
		after, _ := strconv.Atoi(ranges[1])
		for i := before; i <= after; i++ {
			check := strconv.Itoa(i)
			if check[:len(check)/2] == check[len(check)/2:] {
				value, _ := strconv.Atoi(check)
				sum_ += value
			}
		}
	}
	return strconv.Itoa(sum_)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	sum_ := 0
	for _, line := range strings.Split(lines[0], ",") {
		ranges := strings.Split(line, "-")
		before, _ := strconv.Atoi(ranges[0])
		after, _ := strconv.Atoi(ranges[1])
		for i := before; i <= after; i++ {
			check := strconv.Itoa(i)
			invalid_id := false
			check_len := len(check)
			for j := 1; j < check_len; j++ {
				valid := true
				if check_len%j != 0 {
					continue
				}
				parts := check_len / j
				for k := 0; k < parts-1; k++ {
					left := check[k*j : (k+1)*j]
					right := check[(k+1)*j : (k+2)*j]
					if left != right {
						valid = false
						break
					}
				}
				if valid {
					invalid_id = true
					break
				}
			}
			if invalid_id {
				value, _ := strconv.Atoi(check)
				sum_ += value
			}
		}
	}
	return strconv.Itoa(sum_)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
