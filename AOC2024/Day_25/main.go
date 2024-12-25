package main

import (
	"fmt"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	objects := make([][]string, 0)
	objects = append(objects, make([]string, 0))
	idx := 0
	for _, line := range lines {
		if line == "" {
			idx += 1
			objects = append(objects, make([]string, 0))
			continue
		}
		objects[idx] = append(objects[idx], line)
	}
	keys := make([]aocutils.Grid, 0)
	locks := make([]aocutils.Grid, 0)
	for _, object := range objects {
		if object[0] == "....." {
			keys = append(keys, aocutils.CreateCharacterGrid(object))
		} else {
			locks = append(locks, aocutils.CreateCharacterGrid(object))
		}
	}
	count := 0
	for _, lock := range locks {
		for _, key := range keys {
			all := make(map[aocutils.Key]int)
			valid := true
			for val, char := range lock.Data {
				if char == '#' {
					all[val] += 1
				} else {
					all[val] += 0
				}
			}
			for val, char := range key.Data {
				if char == '#' {
					all[val] += 1
					if all[val] > 1 {
						valid = false
						break
					}
				} else {
					all[val] += 0
				}
			}
			if valid {
				count += 1
			}
		}
	}
	return strconv.Itoa(count)
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
