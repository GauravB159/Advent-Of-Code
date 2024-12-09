package main

import (
	"fmt"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	disk := make([]string, 0)
	gif := aocutils.CreateGIF("onestar", 10)
	for index, char := range lines[0] {
		val, _ := strconv.Atoi(string(char))
		for range val {
			if index%2 == 1 {

				disk = append(disk, ".")
			} else {

				disk = append(disk, strconv.Itoa(index/2))
			}
		}
	}
	height := 500
	grid_image := aocutils.CreateImage(height, len(disk), 1, "onestar")
	grid_image.UsePaletteColors()
	for index, val := range disk {
		if val == "." {
			for i := range height {
				grid_image.SetZoomedPixel(index, i, 2)
			}
		} else {
			for i := range height {
				grid_image.SetZoomedPixel(index, i, 5)
			}
		}
		if index%3 == 0 {
			gif.AddFrame(grid_image)
		}
	}
	lptr := 0
	rptr := len(disk) - 1
	for lptr < rptr {
		if disk[lptr] != "." {
			lptr += 1
		} else if disk[rptr] == "." {
			rptr -= 1
		} else {
			for i := range height {
				grid_image.SetZoomedPixel(lptr, i, 5)
				grid_image.SetZoomedPixel(rptr, i, 2)
			}
			disk[lptr], disk[rptr] = disk[rptr], disk[lptr]
			gif.AddFrame(grid_image)
		}
	}
	for i := 0; i < 100; i++ {
		gif.AddFrame(grid_image)
	}
	gif.WriteGIFToFile()
	sum := 0
	for index, val := range disk {
		if val == "." {
			break
		}
		intval, _ := strconv.Atoi(val)
		sum += index * intval
	}
	return strconv.Itoa(sum)
}

type location struct {
	index int
	size  int
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	disk := make([]string, 0)
	freespace := make([]location, 0)
	files := make([]location, 0)
	gif := aocutils.CreateGIF("twostar", 10)
	for index, char := range lines[0] {
		val, _ := strconv.Atoi(string(char))
		location_value := location{
			index: len(disk),
			size:  val,
		}
		if index%2 == 1 {
			freespace = append(freespace, location_value)
		} else {
			files = append(files, location_value)
		}
		for range val {
			if index%2 == 1 {
				disk = append(disk, ".")
			} else {
				disk = append(disk, strconv.Itoa(index/2))
			}
		}
	}
	height := 500
	grid_image := aocutils.CreateImage(height, len(disk), 1, "twostar")
	grid_image.UsePaletteColors()
	for index, val := range disk {
		if val == "." {
			for i := range height {
				grid_image.SetZoomedPixel(index, i, 2)
			}
		} else {
			for i := range height {
				grid_image.SetZoomedPixel(index, i, 5)
			}
		}
		if index%3 == 0 {
			gif.AddFrame(grid_image)
		}
	}
	for index := len(files) - 1; index >= 0; index-- {
		for free_index := 0; free_index < len(freespace); free_index++ {
			if freespace[free_index].index < files[index].index && freespace[free_index].size >= files[index].size {
				for i := 0; i < files[index].size; i++ {
					disk[freespace[free_index].index+i] = strconv.Itoa(index)
					disk[files[index].index+i] = "."
					for j := range height {
						grid_image.SetZoomedPixel(freespace[free_index].index+i, j, 5)
						grid_image.SetZoomedPixel(files[index].index+i, j, 2)
					}
					gif.AddFrame(grid_image)
				}
				freespace[free_index].size -= files[index].size
				freespace[free_index].index += files[index].size
				break
			}
		}
	}
	for i := 0; i < 100; i++ {
		gif.AddFrame(grid_image)
	}
	gif.WriteGIFToFile()
	sum := 0
	for index, val := range disk {
		if val == "." {
			continue
		}
		intval, _ := strconv.Atoi(val)
		sum += index * intval
	}
	return strconv.Itoa(sum)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	aocutils.Timer("2 star", twostar, "input.txt")
}
