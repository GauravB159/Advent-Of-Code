package main

import (
	"fmt"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	var directions [8]aocutils.Key = [8]aocutils.Key{{Row: 1, Col: 0}, {Row: -1, Col: 0}, {Row: 0, Col: 1}, {Row: 0, Col: -1}, {Row: 1, Col: -1}, {Row: 1, Col: 1}, {Row: -1, Col: -1}, {Row: -1, Col: 1}}
	count := 0
	for position := range grid.Data {
		for _, direction := range directions {
			check := ""
			for i := 0; i < 4; i++ {
				key := aocutils.Key{Row: position.Row + i*direction.Row, Col: position.Col + i*direction.Col}
				check += string(grid.Data[key])
				if check == "XMAS" {
					count += 1
				}
			}
		}
	}
	return strconv.Itoa(count)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	count := 0
	for position := range grid.Data {
		leftDiagonal := string(grid.Data[aocutils.Key{Row: position.Row - 1, Col: position.Col - 1}]) + string(grid.Data[position]) + string(grid.Data[aocutils.Key{Row: position.Row + 1, Col: position.Col + 1}])
		rightDiagonal := string(grid.Data[aocutils.Key{Row: position.Row - 1, Col: position.Col + 1}]) + string(grid.Data[position]) + string(grid.Data[aocutils.Key{Row: position.Row + 1, Col: position.Col - 1}])
		if (leftDiagonal == "MAS" || leftDiagonal == "SAM") && (rightDiagonal == "MAS" || rightDiagonal == "SAM") {
			count += 1
		}
	}
	return strconv.Itoa(count)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
