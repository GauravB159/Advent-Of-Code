package main

import (
	"fmt"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := ""
	numbers := strings.Split(lines[0], ",")
	board_rows_cols := make(map[string]int, 0)
	reverse_map := make(map[string][]string)
	boards := make(map[int]map[string]string, 0)
	count := 0
	for i := 2; i < len(lines); i += 6 {
		boards[count] = make(map[string]string)
		for j, line := range lines[i : i+5] {
			for k, value := range strings.Fields(line) {
				boards[count][fmt.Sprintf("%d-%d", j, k)] = value
				if _, exists := reverse_map[value]; !exists {
					reverse_map[value] = make([]string, 0)
				}
				reverse_map[value] = append(reverse_map[value], fmt.Sprintf("%d-%d-%d", count, j, k))
				board_rows_cols[fmt.Sprintf("%d-x-%d", count, j)] = 0
				board_rows_cols[fmt.Sprintf("%d-%d-x", count, k)] = 0
			}
		}
		count += 1
	}
	for _, number := range numbers {
		for _, position := range reverse_map[number] {
			position_split := strings.Split(position, "-")
			board, _ := strconv.Atoi(position_split[0])
			row := position_split[1]
			col := position_split[2]
			board_rows_cols[fmt.Sprintf("%d-x-%s", board, row)] += 1
			board_rows_cols[fmt.Sprintf("%d-%s-x", board, col)] += 1
			delete(boards[board], fmt.Sprintf("%s-%s", row, col))
			if board_rows_cols[fmt.Sprintf("%d-x-%s", board, row)] == 5 || board_rows_cols[fmt.Sprintf("%d-%s-x", board, col)] == 5 {
				sum := 0
				for _, value := range boards[board] {
					value_int, _ := strconv.Atoi(value)
					sum += value_int
				}
				number_int, _ := strconv.Atoi(number)
				return strconv.Itoa(sum * number_int)
			}
		}
	}
	return result
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := ""
	numbers := strings.Split(lines[0], ",")
	board_rows_cols := make(map[string]int, 0)
	reverse_map := make(map[string][]string)
	boards := make(map[int]map[string]string, 0)
	count := 0
	for i := 2; i < len(lines); i += 6 {
		boards[count] = make(map[string]string)
		for j, line := range lines[i : i+5] {
			for k, value := range strings.Fields(line) {
				boards[count][fmt.Sprintf("%d-%d", j, k)] = value
				if _, exists := reverse_map[value]; !exists {
					reverse_map[value] = make([]string, 0)
				}
				reverse_map[value] = append(reverse_map[value], fmt.Sprintf("%d-%d-%d", count, j, k))
				board_rows_cols[fmt.Sprintf("%d-x-%d", count, j)] = 0
				board_rows_cols[fmt.Sprintf("%d-%d-x", count, k)] = 0
			}
		}
		count += 1
	}
	for _, number := range numbers {
		for _, position := range reverse_map[number] {
			position_split := strings.Split(position, "-")
			board, _ := strconv.Atoi(position_split[0])
			row := position_split[1]
			col := position_split[2]
			board_rows_cols[fmt.Sprintf("%d-x-%s", board, row)] += 1
			board_rows_cols[fmt.Sprintf("%d-%s-x", board, col)] += 1
			if _, exists := boards[board]; exists {
				delete(boards[board], fmt.Sprintf("%s-%s", row, col))
				if board_rows_cols[fmt.Sprintf("%d-x-%s", board, row)] == 5 || board_rows_cols[fmt.Sprintf("%d-%s-x", board, col)] == 5 {
					if len(boards) == 1 {
						sum := 0
						for _, value := range boards[board] {
							value_int, _ := strconv.Atoi(value)
							sum += value_int
						}
						number_int, _ := strconv.Atoi(number)
						return strconv.Itoa(sum * number_int)
					} else {
						delete(boards, board)
					}
				}
			}
		}
	}
	return result
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
