package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	a, _ := strconv.Atoi(strings.Split(lines[0], ": ")[1])
	b, _ := strconv.Atoi(strings.Split(lines[1], ": ")[1])
	c, _ := strconv.Atoi(strings.Split(lines[2], ": ")[1])
	program := strings.Split(strings.Split(lines[4], ": ")[1], ",")
	iptr := 0
	combo_operand := [7]int{0, 1, 2, 3, a, b, c}
	dvmap := map[string]int{
		"0": 4,
		"6": 5,
		"7": 6,
	}
	file := make([]string, 0)
	for iptr < len(program) {
		opcode := program[iptr]
		operand, _ := strconv.Atoi(program[iptr+1])
		fmt.Println(opcode, operand, combo_operand)
		if opcode == "0" || opcode == "6" || opcode == "7" {
			numerator := combo_operand[4]
			denominator := math.Pow(2, float64(combo_operand[operand]))
			div := int(numerator / int(denominator))
			combo_operand[dvmap[opcode]] = div
		} else if program[iptr] == "1" {
			combo_operand[5] ^= operand
		} else if program[iptr] == "2" {
			combo_operand[5] = combo_operand[operand] % 8
		} else if program[iptr] == "3" {
			if combo_operand[4] == 0 {
				iptr += 2
				continue
			}
			iptr = operand
			continue
		} else if program[iptr] == "4" {
			combo_operand[5] ^= combo_operand[6]
		} else if program[iptr] == "5" {
			file = append(file, strconv.Itoa(combo_operand[operand]%8))
		}
		iptr += 2
	}
	return strings.Join(file, ",")
}

// func twostar(filename string) string {
// 	lines := aocutils.Readfile(filename)
// 	result := ""
// 	for _, line := range lines {
// 		fmt.Println(line)
// 	}
// 	return result
// }

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	// aocutils.Timer("2 star", twostar, "input.txt")
}
