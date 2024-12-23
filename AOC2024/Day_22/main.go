package main

import (
	"fmt"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	turns := 2000
	prune := 16777216
	sum := 0
	for _, line := range lines {
		secret, _ := strconv.Atoi(line)
		for range turns {
			secret = ((secret * 64) ^ secret) % prune
			secret = ((secret / 32) ^ secret) % prune
			secret = ((secret * 2048) ^ secret) % prune
		}
		sum += secret
	}
	return strconv.Itoa(sum)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	turns := 2000
	prune := 16777216
	vals := make(map[string]int)
	for _, line := range lines {
		secret, _ := strconv.Atoi(line)
		secrets := make([]int, 0)
		for range turns {
			secret = ((secret * 64) ^ secret) % prune
			secret = ((secret / 32) ^ secret) % prune
			secret = ((secret * 2048) ^ secret) % prune
			secrets = append(secrets, secret%10)
		}
		diffs := make([]int, 0)
		for idx := range secrets[1:] {
			diffs = append(diffs, secrets[idx+1]-secrets[idx])
		}
		seen := make(map[string]bool)
		for idx := range diffs[:len(diffs)-3] {
			key := strconv.Itoa(diffs[idx]) + strconv.Itoa(diffs[idx+1]) + strconv.Itoa(diffs[idx+2]) + strconv.Itoa(diffs[idx+3])
			if seen[key] {
				continue
			}
			seen[key] = true
			vals[key] += secrets[idx+4]
		}
	}
	max := 0
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return strconv.Itoa(max)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
