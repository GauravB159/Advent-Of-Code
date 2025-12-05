package main

import (
	"fmt"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

type Range struct {
	min int
	max int
}

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	parsingRange := true
	ranges := make([]Range, 0)
	candidates := make([]int, 0)
	for _, line := range lines {
		if line == "" {
			parsingRange = false
			continue
		}
		if parsingRange {
			temp := strings.Split(line, "-")
			min, _ := strconv.Atoi(temp[0])
			max, _ := strconv.Atoi(temp[1])
			ranges = append(ranges, Range{min: min, max: max})
		} else {
			check, _ := strconv.Atoi(line)
			candidates = append(candidates, check)
		}
	}
	count := 0
	for _, candidate := range candidates {
		found := false
		for _, range_ := range ranges {
			if candidate >= range_.min && candidate <= range_.max {
				found = true
				break
			}
		}
		if found {
			count += 1
		}
	}
	return strconv.Itoa(count)
}

func get_updated_ranges(ranges map[int]int) (map[int]int, int) {
	new_ranges := make(map[int]int, 0)
	count := 0
	for range_outer_min := range ranges {
		range_outer_max := ranges[range_outer_min]
		new_ranges[range_outer_min] = range_outer_max
		for range_inner_min := range ranges {
			if range_inner_min == range_outer_min {
				continue
			}
			range_inner_max := ranges[range_inner_min]
			if range_inner_min > range_outer_min && range_inner_max <= range_outer_max {
				count += 1
				delete(new_ranges, range_inner_min)
				continue
			} else if range_inner_min >= range_outer_min && range_inner_min <= range_outer_max {
				count += 1
				new_ranges[range_outer_min] = range_inner_max
			} else if range_inner_max >= range_outer_min && range_inner_max <= range_outer_max {
				count += 1
				new_ranges[range_inner_min] = range_outer_max
			}
		}
	}
	return new_ranges, count
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	parsingRange := true
	ranges := make(map[int]int, 0)
	for _, line := range lines {
		if line == "" {
			parsingRange = false
			continue
		}
		if parsingRange {
			temp := strings.Split(line, "-")
			min, _ := strconv.Atoi(temp[0])
			max, _ := strconv.Atoi(temp[1])
			if _, exists := ranges[min]; exists {
				if ranges[min] < max {
					ranges[min] = max
				}
			} else {
				ranges[min] = max
			}
		}
	}
	count := -1
	for count != 0 {
		new_ranges, new_count := get_updated_ranges(ranges)
		ranges = new_ranges
		count = new_count
	}
	result := 0
	for min := range ranges {
		result += ranges[min] - min + 1
	}
	return strconv.Itoa(result)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
