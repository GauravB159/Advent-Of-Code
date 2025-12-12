package main

import (
	"strconv"
	"strings"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

type Tree struct {
	grid       aocutils.Grid
	counts     []int
	totalCount int
}

type Shape struct {
	rotations []aocutils.Grid
}

func checkShapeAtLocation(grid aocutils.Grid, shape aocutils.Grid, location aocutils.Key) bool {
	for i := 0; i < shape.Cols; i++ {
		for j := 0; j < shape.Rows; j++ {
			if !grid.Exists(j+location.Row, i+location.Col) {
				return false
			}
			if grid.Data[aocutils.Key{Row: j + location.Row, Col: i + location.Col}] != '.' {
				if shape.Data[aocutils.Key{Row: j, Col: i}] != '.' {
					return false
				}
			}
		}
	}
	return true
}

func applyShape(grid aocutils.Grid, shape aocutils.Grid, location aocutils.Key) aocutils.Grid {
	for i := 0; i < shape.Cols; i++ {
		for j := 0; j < shape.Rows; j++ {
			if shape.Data[aocutils.Key{Row: j, Col: i}] != '.' {
				grid.SetValue(j+location.Row, i+location.Col, shape.Data[aocutils.Key{Row: j, Col: i}])
			}
		}
	}
	return grid
}

/* Not even required, what a waste of time. Keeping it cause why not.  */
func allShapesFit(tree Tree, shapes []Shape) bool {
	allDone := true
	tree.grid.PrintChar()
	for _, check := range tree.counts {
		if check > 0 {
			allDone = false
		}
	}
	if allDone {
		return true
	}
	result := false
	for key := range tree.grid.Data {
		for shapeIndex, count := range tree.counts {
			if count == 0 {
				continue
			}
			for _, rotation := range shapes[shapeIndex].rotations {
				if checkShapeAtLocation(tree.grid, rotation, key) {
					tree.counts[shapeIndex] -= 1
					new_grid := applyShape(tree.grid.Copy(), rotation, key)
					result = result || allShapesFit(Tree{grid: new_grid, counts: tree.counts}, shapes)
					tree.counts[shapeIndex] += 1
					if result {
						return result
					}
				}
			}

		}
	}
	return false
}

func rotate(grid aocutils.Grid, rotation int) aocutils.Grid {
	if rotation%90 != 0 {
		return grid
	}
	if rotation == 0 {
		return grid
	}
	new_grid := aocutils.CreateGridByDimensions(grid.Cols, grid.Rows, 0)
	for i := 0; i < rotation/90; i++ {
		for key := range grid.Data {
			/* Rotates a co-ordinate by 90 */
			new_grid.SetValue(key.Col, grid.Rows-key.Row-1, grid.Data[key])
		}
		grid = new_grid
		new_grid = aocutils.CreateGridByDimensions(grid.Cols, grid.Rows, 0)
	}
	return grid
}

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	shapes := make([]Shape, 0)
	trees := make([]Tree, 0)
	for i, line := range lines {
		if len(line) > 0 && line[1] == ':' {
			grid := aocutils.CreateCharacterGrid(lines[i+1 : i+4])
			for key := range grid.Data {
				idx, _ := strconv.Atoi(string(line[0]))
				if grid.Data[key] == '#' {
					grid.Data[key] = 'A' + idx
				}
			}
			shape_rotations := make([]aocutils.Grid, 0)
			for rotation := 0; rotation <= 270; rotation += 90 {
				rotation_grid := rotate(grid, rotation)
				shape_rotations = append(shape_rotations, rotation_grid)
			}
			shapes = append(shapes, Shape{rotations: shape_rotations})
		}
		check := strings.Split(line, "x")
		if len(check) > 1 {
			inner := strings.Split(check[1], ":")
			width, _ := strconv.Atoi(check[0])
			height, _ := strconv.Atoi(inner[0])
			counts := make([]int, 0)
			totalCount := 0
			for _, count := range strings.Split(inner[1], " ") {
				if count == "" {
					continue
				}
				count_int, _ := strconv.Atoi(string(count))
				counts = append(counts, count_int)
				totalCount += count_int
			}
			trees = append(trees, Tree{grid: aocutils.CreateGridByDimensions(width, height, '.'), counts: counts, totalCount: totalCount})
		}
	}
	count := 0
	too_large_count := 0
	for _, tree := range trees {
		if tree.totalCount*7 > tree.grid.Rows*tree.grid.Cols {
			too_large_count += 1
		} else {
			count += 1
		}
	}
	return strconv.Itoa(count)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
}
