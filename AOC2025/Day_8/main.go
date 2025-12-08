package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

type Coord struct {
	x int
	y int
	z int
}

type CoordPair struct {
	first    Coord
	second   Coord
	distance float64
}

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	n := 1000
	k := 3
	result := 1
	coords := make([]Coord, 0)
	for _, line := range lines {
		temp := strings.Split(line, ",")
		x, _ := strconv.Atoi(temp[0])
		y, _ := strconv.Atoi(temp[1])
		z, _ := strconv.Atoi(temp[2])
		coords = append(coords, Coord{x: x, y: y, z: z})
	}
	groups := make(map[Coord]int)
	reverse_groups := make(map[int][]Coord)
	distances := make([]CoordPair, 0)
	for i, first := range coords {
		groups[first] = i
		reverse_groups[i] = []Coord{first}
		for j, second := range coords {
			if i <= j {
				continue
			}
			distance := math.Sqrt(math.Pow(float64(first.x)-float64(second.x), 2) + math.Pow(float64(first.y)-float64(second.y), 2) + math.Pow(float64(first.z)-float64(second.z), 2))
			pair := CoordPair{first: first, second: second, distance: distance}
			distances = append(distances, pair)
		}
	}
	sort.Slice(distances, func(x int, y int) bool {
		return distances[x].distance < distances[y].distance
	})

	for _, coord := range distances[:n] {
		if groups[coord.first] == groups[coord.second] {
			continue
		}
		if len(reverse_groups[groups[coord.second]]) > 0 {
			temp := groups[coord.second]
			for _, group_member := range reverse_groups[groups[coord.second]] {
				groups[group_member] = groups[coord.first]
				reverse_groups[groups[coord.first]] = append(reverse_groups[groups[coord.first]], group_member)
			}
			delete(reverse_groups, temp)
		} else {
			temp := groups[coord.first]
			for _, group_member := range reverse_groups[groups[coord.first]] {
				groups[group_member] = groups[coord.second]
				reverse_groups[groups[coord.second]] = append(reverse_groups[groups[coord.second]], group_member)
			}
			delete(reverse_groups, temp)
		}
	}
	sizes := make([]int, 0)
	for _, group := range reverse_groups {
		sizes = append(sizes, -len(group))
	}
	sort.IntSlice(sizes).Sort()
	for _, x := range sizes[:k] {
		result *= (-x)
	}
	return strconv.Itoa(result)
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	result := 1
	coords := make([]Coord, 0)
	for _, line := range lines {
		temp := strings.Split(line, ",")
		x, _ := strconv.Atoi(temp[0])
		y, _ := strconv.Atoi(temp[1])
		z, _ := strconv.Atoi(temp[2])
		coords = append(coords, Coord{x: x, y: y, z: z})
	}
	groups := make(map[Coord]int)
	reverse_groups := make(map[int][]Coord)
	distances := make([]CoordPair, 0)
	for i, first := range coords {
		groups[first] = i
		reverse_groups[i] = []Coord{first}
		for j, second := range coords {
			if i <= j {
				continue
			}
			distance := math.Sqrt(math.Pow(float64(first.x)-float64(second.x), 2) + math.Pow(float64(first.y)-float64(second.y), 2) + math.Pow(float64(first.z)-float64(second.z), 2))
			pair := CoordPair{first: first, second: second, distance: distance}
			distances = append(distances, pair)
		}
	}
	sort.Slice(distances, func(x int, y int) bool {
		return distances[x].distance < distances[y].distance
	})
	var result_coords CoordPair
	for _, coord := range distances {
		if groups[coord.first] == groups[coord.second] {
			continue
		}
		if len(reverse_groups[groups[coord.second]]) > 0 {
			temp := groups[coord.second]
			for _, group_member := range reverse_groups[groups[coord.second]] {
				groups[group_member] = groups[coord.first]
				reverse_groups[groups[coord.first]] = append(reverse_groups[groups[coord.first]], group_member)
			}
			delete(reverse_groups, temp)
		} else {
			temp := groups[coord.first]
			for _, group_member := range reverse_groups[groups[coord.first]] {
				groups[group_member] = groups[coord.second]
				reverse_groups[groups[coord.second]] = append(reverse_groups[groups[coord.second]], group_member)
			}
			delete(reverse_groups, temp)
		}
		if len(reverse_groups) == 1 {
			result_coords = coord
			break
		}
	}
	result = result_coords.first.x * result_coords.second.x
	return strconv.Itoa(result)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
