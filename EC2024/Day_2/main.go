package main

import (
	"fmt"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func part_one(filename string) string {
	lines := aocutils.Readfile(filename)
	words := strings.Split(strings.Split(lines[0], ":")[1], ",")
	passage := lines[2]
	count := 0
	for i := 0; i < len(passage); i++ {
		for _, word := range words {
			if i+len(word) >= len(passage) {
				continue
			}
			match := true
			for j := 0; j < len(word); j++ {
				if passage[i+j] != word[j] {
					match = false
					break
				}
			}
			if match {
				count += 1
				break
			}
		}
	}
	return strconv.Itoa(count)
}

func part_two(filename string) string {
	lines := aocutils.Readfile(filename)
	words := strings.Split(strings.Split(lines[0], ":")[1], ",")
	for _, word := range words {
		if len(words) == 1 {
			continue
		}
		words = append(words, aocutils.Reverse(word))
	}
	passage := lines[2:]
	count := 0
	for _, line := range passage {
		runic := make([]int, len(line))
		for i := 0; i < len(line); i++ {
			for _, word := range words {
				if i+len(word) > len(line) {
					continue
				}
				match := true
				for j := 0; j < len(word); j++ {
					if line[i+j] != word[j] {
						match = false
						break
					}
				}
				if match {
					for k := 0; k < len(word); k++ {
						runic[i+k] = 1
					}
				}
			}
		}
		count += aocutils.ArrayIntSum(runic)
	}

	return strconv.Itoa(count)
}

func part_three(filename string) string {
	lines := aocutils.Readfile(filename)
	words := strings.Split(strings.Split(lines[0], ":")[1], ",")
	for _, word := range words {
		if len(words) == 1 {
			continue
		}
		words = append(words, aocutils.Reverse(word))
	}
	passage := aocutils.CreateCharacterGrid(lines[2:])
	marked := aocutils.CreateGrid(lines[2:])
	for key := range passage.Data {
		for _, word := range words {
			matchH := true
			matchV := true
			for i := 0; i < len(word); i++ {
				charV, existsV := passage.Data[aocutils.Key{Row: key.Row + i, Col: key.Col}]
				charH, existsH := passage.Data[aocutils.Key{Row: key.Row, Col: (key.Col + i) % passage.Cols}]
				if !existsH || charH != int(word[i]) {
					matchH = false
				}
				if !existsV || charV != int(word[i]) {
					matchV = false
				}
			}
			if matchH {
				for i := 0; i < len(word); i++ {
					marked.Data[aocutils.Key{Row: key.Row, Col: (key.Col + i) % passage.Cols}] = 1
				}
			}
			if matchV {
				for i := 0; i < len(word); i++ {
					marked.Data[aocutils.Key{Row: key.Row + i, Col: key.Col}] = 1
				}
			}
		}
	}
	sum := 0
	for _, value := range marked.Data {
		sum += value
	}
	return strconv.Itoa(sum)
}

func main() {
	aocutils.Timer("Part 1", part_one, "input_one.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("Part 2", part_two, "input_two.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("Part 3", part_three, "input_three.txt")
}
