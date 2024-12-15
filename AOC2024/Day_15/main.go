package main

import (
	"fmt"
	"sort"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	idx := -1
	for index, line := range lines {
		if line == "" {
			idx = index
		}
	}
	grid := aocutils.CreateCharacterGrid(lines[:idx])
	var start aocutils.Key
	for position, value := range grid.Data {
		if value == '@' {
			start = position
			break
		}
	}
	moves := lines[idx+1:]
	directions := map[rune]aocutils.Key{
		'^': {Row: -1, Col: 0},
		'>': {Row: 0, Col: 1},
		'v': {Row: 1, Col: 0},
		'<': {Row: 0, Col: -1},
	}
	for _, moveline := range moves {
		for _, move := range moveline {
			direction := directions[move]
			var stack aocutils.Stack[aocutils.Key]
			shouldMove := true
			nextpos := start
			for true {
				nextpos = aocutils.Key{Row: nextpos.Row + direction.Row, Col: nextpos.Col + direction.Col}
				if _, exists := grid.Data[nextpos]; !exists {
					break
				}
				if grid.Data[nextpos] == '#' {
					shouldMove = false
					break
				} else if grid.Data[nextpos] == '.' {
					stack.Push(nextpos)
					break
				} else if grid.Data[nextpos] == 'O' {
					stack.Push(nextpos)
				}
			}
			if shouldMove {
				for !stack.IsEmpty() {
					pos, _ := stack.Pop()
					prevpos := aocutils.Key{Row: pos.Row - direction.Row, Col: pos.Col - direction.Col}
					if stack.IsEmpty() {
						start = pos
						grid.Data[prevpos] = '.'
						grid.Data[pos] = '@'
					} else {
						grid.Data[pos] = 'O'
					}
				}
			}
			// grid.PrintChar()
		}
	}
	answer := 0
	for position, value := range grid.Data {
		if value == 'O' {
			answer += 100*position.Row + position.Col
		}
	}
	return strconv.Itoa(answer)
}

func shouldMove(node aocutils.Key, grid *aocutils.Grid, direction aocutils.Key, willMove *bool, visited *map[aocutils.Key]bool, stack *aocutils.Stack[aocutils.Key]) {
	if (*visited)[node] {
		return
	}
	if grid.Data[node] == '#' {
		*willMove = false
		return
	}
	if grid.Data[node] == '.' {
		return
	}
	(*visited)[node] = true
	(*stack).Push(node)
	if grid.Data[node] == '[' {
		shouldMove(aocutils.Key{Row: node.Row, Col: node.Col + 1}, grid, direction, willMove, visited, stack)
	}
	if grid.Data[node] == ']' {
		shouldMove(aocutils.Key{Row: node.Row, Col: node.Col - 1}, grid, direction, willMove, visited, stack)
	}
	nextNode := aocutils.Key{Row: node.Row + direction.Row, Col: node.Col + direction.Col}
	shouldMove(nextNode, grid, direction, willMove, visited, stack)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	idx := -1
	gridlines := make([]string, 0)
	for index, line := range lines {
		gridline := ""
		if line == "" {
			idx = index
			break
		}
		for _, char := range line {
			if char == '#' {
				gridline += "##"
			}
			if char == '.' {
				gridline += ".."
			}
			if char == 'O' {
				gridline += "[]"
			}
			if char == '@' {
				gridline += "@."
			}
		}
		gridlines = append(gridlines, gridline)
	}
	grid := aocutils.CreateCharacterGrid(gridlines)
	var start aocutils.Key
	for position, value := range grid.Data {
		if value == '@' {
			start = position
			break
		}
	}
	moves := lines[idx+1:]
	directions := map[rune]aocutils.Key{
		'^': {Row: -1, Col: 0},
		'>': {Row: 0, Col: 1},
		'v': {Row: 1, Col: 0},
		'<': {Row: 0, Col: -1},
	}
	type movement struct {
		key  aocutils.Key
		char string
	}
	grid_image := aocutils.CreateImage(grid.Rows, grid.Cols, 20, 10, "twostar")
	grid_image.UseFullColors()
	pixelMap := map[rune]int{
		'.': 5,
		'#': 1,
		'@': 10,
		'[': 7,
		']': 9,
	}
	fmt.Println(grid.Rows, grid.Cols)
	for key, value := range grid.Data {
		grid_image.SetZoomedPixel(key.Col, key.Row, pixelMap[rune(value)])
	}
	gif := aocutils.CreateGIF("twostar", 25)
	gif.AddFrame(grid_image)
	count := 0
	for _, moveline := range moves {
		for _, move := range moveline {
			count += 1
			direction := directions[move]
			willMove := true
			visited := make(map[aocutils.Key]bool, 0)
			var stack aocutils.Stack[aocutils.Key]
			shouldMove(start, &grid, direction, &willMove, &visited, &stack)
			sort.Slice(stack, func(a int, b int) bool {
				if move == '^' {
					return stack[a].Row > stack[b].Row
				}
				if move == 'v' {
					return stack[a].Row < stack[b].Row
				}
				if move == '>' {
					return stack[a].Col < stack[b].Col
				}
				return stack[a].Col > stack[b].Col
			})
			if willMove {
				for !stack.IsEmpty() {
					pos, _ := stack.Pop()
					movepos := aocutils.Key{Row: pos.Row + direction.Row, Col: pos.Col + direction.Col}
					grid.Data[movepos], grid.Data[pos] = grid.Data[pos], grid.Data[movepos]
				}
				start = aocutils.Key{Row: start.Row + direction.Row, Col: start.Col + direction.Col}
			}
			fmt.Println(count)
			if gif.Framecount%gif.Frameskip == 0 {
				for key, value := range grid.Data {
					grid_image.SetZoomedPixel(key.Col, key.Row, pixelMap[rune(value)])
				}
			}
			gif.AddFrame(grid_image)
		}
	}
	gif.WriteGIFToFile()
	answer := 0
	for position, value := range grid.Data {
		y := position.Row
		x := position.Col
		if value == '[' {
			answer += 100*y + x
		}
	}
	return strconv.Itoa(answer)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
