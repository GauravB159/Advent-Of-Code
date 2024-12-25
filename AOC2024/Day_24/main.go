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
	switched := false
	values := make(map[string]bool)
	connections := make([]string, 0)
	for _, line := range lines {
		if line == "" {
			switched = true
			continue
		}
		if switched {
			connections = append(connections, line)
		} else {
			val := strings.Split(line, ": ")[0]
			num, _ := strconv.Atoi(strings.Split(line, ": ")[1])
			values[val] = (num == 1)
		}
	}
	for true {
		connected := false
		for _, connection := range connections {
			result := strings.Split(connection, " -> ")[1]
			if _, exists := values[result]; exists {
				continue
			}
			operation := strings.Fields(strings.Split(connection, " -> ")[0])
			if _, exists := values[operation[0]]; !exists {
				continue
			}
			if _, exists := values[operation[2]]; !exists {
				continue
			}
			if operation[1] == "OR" {
				values[result] = values[operation[0]] || values[operation[2]]
			}
			if operation[1] == "AND" {
				values[result] = values[operation[0]] && values[operation[2]]
			}
			if operation[1] == "XOR" {
				values[result] = values[operation[0]] != values[operation[2]]
			}
			connected = true
		}
		if !connected {
			break
		}
	}
	var bits [60]bool
	for key := range values {
		if key[0] == 'z' {
			bit, _ := strconv.Atoi(key[1:])
			bits[bit] = values[key]
		}
	}
	sum := 0.0
	for i, val := range bits {
		if val {
			sum += math.Pow(2, float64(i))
		}
	}
	return strconv.Itoa(int(sum))
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	switched := false
	values := make(map[string]bool)
	connections := make([]string, 0)
	for _, line := range lines {
		if line == "" {
			switched = true
			continue
		}
		if switched {
			connections = append(connections, line)
		} else {
			val := strings.Split(line, ": ")[0]
			num, _ := strconv.Atoi(strings.Split(line, ": ")[1])
			values[val] = (num == 1)
		}
	}
	for true {
		connected := false
		for _, connection := range connections {
			result := strings.Split(connection, " -> ")[1]
			if _, exists := values[result]; exists {
				continue
			}
			operation := strings.Fields(strings.Split(connection, " -> ")[0])
			if _, exists := values[operation[0]]; !exists {
				continue
			}
			if _, exists := values[operation[2]]; !exists {
				continue
			}
			if operation[1] == "OR" {
				values[result] = values[operation[0]] || values[operation[2]]
			}
			if operation[1] == "AND" {
				values[result] = values[operation[0]] && values[operation[2]]
			}
			if operation[1] == "XOR" {
				values[result] = values[operation[0]] != values[operation[2]]
			}
			connected = true
		}
		if !connected {
			break
		}
	}
	var xbits [60]bool
	var ybits [60]bool
	var zbits [60]bool
	for key := range values {
		if key[0] == 'z' {
			bit, _ := strconv.Atoi(key[1:])
			zbits[bit] = values[key]
		}
		if key[0] == 'x' {
			bit, _ := strconv.Atoi(key[1:])
			xbits[bit] = values[key]
		}
		if key[0] == 'y' {
			bit, _ := strconv.Atoi(key[1:])
			ybits[bit] = values[key]
		}
	}
	sum := 0.0
	fmt.Println(xbits, ybits, zbits)
	return strconv.Itoa(int(sum))
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
