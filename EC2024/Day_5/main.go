package main

import (
	"fmt"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func part_one(filename string) string {
	lines := aocutils.Readfile(filename)
	col_grid := make([][]int, 0)
	for i := 0; i < len(strings.Fields(lines[0])); i++ {
		col_grid = append(col_grid, make([]int, 0))
		for _, line := range lines {
			val, _ := strconv.Atoi(strings.Fields(line)[i])
			col_grid[i] = append(col_grid[i], val)
		}
	}
	count := 0
	answer := ""
	for i := 0; count < 10; i = (i + 1) % len(col_grid) {
		num := col_grid[i][0]
		col_grid[i] = col_grid[i][1:]
		direction := 1
		prev_direction := 1
		count_two := 0
		j := 0
		clap_col := col_grid[(i+1)%len(col_grid)]
		for j = 0; count_two < num; j += direction {
			count_two++
			if direction == 0 {
				direction = -1 * prev_direction
			} else if direction == 1 && j == len(clap_col)-1 {
				prev_direction = direction
				direction = 0
			} else if direction == -1 && j == 0 {
				prev_direction = direction
				direction = 0
			}
		}
		final := j - direction
		temp := make([]int, 0)
		if direction == 0 {
			direction = prev_direction
		}
		for k, val := range clap_col {
			if k == final && direction != -1 {
				temp = append(temp, num)
			}
			temp = append(temp, val)
			if k == final && direction == -1 {
				temp = append(temp, num)
			}
		}
		col_grid[(i+1)%len(col_grid)] = temp
		answer = ""
		for _, col := range col_grid {
			answer += fmt.Sprintf("%d", col[0])
		}
		count++
	}
	return answer
}

func part_two(filename string) string {
	lines := aocutils.Readfile(filename)
	col_grid := make([][]int, 0)
	for i := 0; i < len(strings.Fields(lines[0])); i++ {
		col_grid = append(col_grid, make([]int, 0))
		for _, line := range lines {
			val, _ := strconv.Atoi(strings.Fields(line)[i])
			col_grid[i] = append(col_grid[i], val)
		}
	}
	count := 0
	result := 0
	answer := ""
	counts := make(map[string]int, 0)
	for i := 0; ; i = (i + 1) % len(col_grid) {
		num := col_grid[i][0]
		col_grid[i] = col_grid[i][1:]
		direction := 1
		prev_direction := 1
		count_two := 0
		j := 0
		clap_col := col_grid[(i+1)%len(col_grid)]
		if len(clap_col) > 1 {
			for j = 0; count_two < num; j += direction {
				count_two++
				if direction == 0 {
					direction = -1 * prev_direction
				} else if direction == 1 && j == len(clap_col)-1 {
					prev_direction = direction
					direction = 0
				} else if direction == -1 && j == 0 {
					prev_direction = direction
					direction = 0
				}
			}
		} else {
			j = 1
		}
		final := j - direction
		temp := make([]int, 0)
		if direction == 0 {
			direction = prev_direction
		}
		for k, val := range clap_col {
			if k == final && direction != -1 {
				temp = append(temp, num)
			}
			temp = append(temp, val)
			if k == final && direction == -1 {
				temp = append(temp, num)
			}
		}
		col_grid[(i+1)%len(col_grid)] = temp
		answer = ""
		for _, col := range col_grid {
			answer += fmt.Sprintf("%d", col[0])
		}
		if _, exists := counts[answer]; !exists {
			counts[answer] = 0
		}
		counts[answer] += 1
		count++
		if counts[answer] == 2024 {
			mul, _ := strconv.Atoi(answer)
			result = mul * count
			break
		}
	}
	return strconv.Itoa(result)
}

func part_three(filename string) string {
	lines := aocutils.Readfile(filename)
	col_grid := make([][]int, 0)
	for i := 0; i < len(strings.Fields(lines[0])); i++ {
		col_grid = append(col_grid, make([]int, 0))
		for _, line := range lines {
			val, _ := strconv.Atoi(strings.Fields(line)[i])
			col_grid[i] = append(col_grid[i], val)
		}
	}
	count := 0
	answer := ""
	counts := make(map[string]int, 0)
	max := 0
	for i := 0; ; i = (i + 1) % len(col_grid) {
		num := col_grid[i][0]
		col_grid[i] = col_grid[i][1:]
		direction := 1
		prev_direction := 1
		count_two := 0
		j := 0
		clap_col := col_grid[(i+1)%len(col_grid)]
		if len(clap_col) > 1 {
			for j = 0; count_two < num; j += direction {
				count_two++
				if direction == 0 {
					direction = -1 * prev_direction
				} else if direction == 1 && j == len(clap_col)-1 {
					prev_direction = direction
					direction = 0
				} else if direction == -1 && j == 0 {
					prev_direction = direction
					direction = 0
				}
			}
		} else {
			j = 1
		}
		final := j - direction
		temp := make([]int, 0)
		if direction == 0 {
			direction = prev_direction
		}
		for k, val := range clap_col {
			if k == final && direction != -1 {
				temp = append(temp, num)
			}
			temp = append(temp, val)
			if k == final && direction == -1 {
				temp = append(temp, num)
			}
		}
		col_grid[(i+1)%len(col_grid)] = temp
		answer = ""
		for _, col := range col_grid {
			answer += fmt.Sprintf("%d", col[0])
		}
		if _, exists := counts[answer]; !exists {
			counts[answer] = 0
		}
		counts[answer] += 1
		count++
		mul, _ := strconv.Atoi(answer)
		if mul > max {
			max = mul
		}
		if counts[answer] == 10000 {
			break
		}
	}
	return strconv.Itoa(max)
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
