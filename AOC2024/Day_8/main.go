package main

import (
	"fmt"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	char_map := make(map[string][]aocutils.Key, 0)
	for key, value := range grid.Data {
		if value != '.' {
			if _, exists := char_map[string(value)]; !exists {
				char_map[string(value)] = make([]aocutils.Key, 0)
			}
			char_map[string(value)] = append(char_map[string(value)], key)
		}
	}
	antinodes := make(map[aocutils.Key]bool, 0)
	for char := range char_map {
		if len(char_map[char]) < 2 {
			continue
		}
		for key_one := range char_map[char] {
			for key_two := range char_map[char] {
				if key_one >= key_two {
					continue
				}
				x_diff := char_map[char][key_one].Row - char_map[char][key_two].Row
				y_diff := char_map[char][key_one].Col - char_map[char][key_two].Col
				new_one := aocutils.Key{Row: char_map[char][key_two].Row - x_diff, Col: char_map[char][key_two].Col - y_diff}
				new_two := aocutils.Key{Row: char_map[char][key_one].Row + x_diff, Col: char_map[char][key_one].Col + y_diff}
				if _, exists := grid.Data[new_one]; exists {
					antinodes[new_one] = true
				}
				if _, exists := grid.Data[new_two]; exists {
					antinodes[new_two] = true
				}
			}
		}
	}
	return strconv.Itoa(len(antinodes))
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	char_map := make(map[string][]aocutils.Key, 0)
	for key, value := range grid.Data {
		if value != '.' {
			if _, exists := char_map[string(value)]; !exists {
				char_map[string(value)] = make([]aocutils.Key, 0)
			}
			char_map[string(value)] = append(char_map[string(value)], key)
		}
	}
	antinodes := make(map[aocutils.Key]bool, 0)
	for char := range char_map {
		if len(char_map[char]) < 2 {
			continue
		}
		for key_one := range char_map[char] {
			for key_two := range char_map[char] {
				if key_one >= key_two {
					continue
				}
				x_diff := char_map[char][key_one].Row - char_map[char][key_two].Row
				y_diff := char_map[char][key_one].Col - char_map[char][key_two].Col
				check := true
				for x := 0; check; x += 1 {
					check = false
					new_one := aocutils.Key{Row: char_map[char][key_two].Row - x_diff*x, Col: char_map[char][key_two].Col - y_diff*x}
					new_two := aocutils.Key{Row: char_map[char][key_one].Row + x_diff*x, Col: char_map[char][key_one].Col + y_diff*x}
					if _, exists := grid.Data[new_one]; exists {
						check = true
						antinodes[new_one] = true
					}
					if _, exists := grid.Data[new_two]; exists {
						check = true
						antinodes[new_two] = true
					}
				}

			}
		}
	}
	return strconv.Itoa(len(antinodes))
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
