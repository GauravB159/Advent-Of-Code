package main

import (
	"fmt"
	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	disk := make([]string, 0)
	gif := aocutils.CreateGIF("onestar", 250)
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
	width := 310
	grid_image := aocutils.CreateImage(int(len(disk)/width), width, 4, 4, "onestar")
	grid_image.UsePaletteColors()
	for index, val := range disk {
		if val == "." {
			grid_image.SetZoomedPixel(index%width, int(index/width), 2)
		} else {
			grid_image.SetZoomedPixel(index%width, int(index/width), 5)
		}
		if index%10 == 0 {
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
			grid_image.SetZoomedPixel(lptr%width, int(lptr/width), 5)
			grid_image.SetZoomedPixel(rptr%width, int(rptr/width), 2)
			disk[lptr], disk[rptr] = disk[rptr], disk[lptr]
			gif.AddFrame(grid_image)
		}
	}
	for i := 0; i < 1000; i++ {
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
	gif := aocutils.CreateGIF("twostar", 250)
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
	width := 310
	grid_image := aocutils.CreateImage(int(len(disk)/width), width, 4, 4, "twostar")
	grid_image.UsePaletteColors()
	for index, val := range disk {
		if val == "." {
			grid_image.SetZoomedPixel(index%width, int(index/width), 2)
		} else {
			grid_image.SetZoomedPixel(index%width, int(index/width), 5)
		}
		if index%10 == 0 {
			gif.AddFrame(grid_image)
		}
	}
	for index := len(files) - 1; index >= 0; index-- {
		for free_index := 0; free_index < len(freespace); free_index++ {
			if freespace[free_index].index < files[index].index && freespace[free_index].size >= files[index].size {
				for i := 0; i < files[index].size; i++ {
					disk[freespace[free_index].index+i] = strconv.Itoa(index)
					disk[files[index].index+i] = "."
					grid_image.SetZoomedPixel((freespace[free_index].index+i)%width, int((freespace[free_index].index+i)/width), 5)
					grid_image.SetZoomedPixel((files[index].index+i)%width, int((files[index].index+i)/width), 2)
					gif.AddFrame(grid_image)
				}
				freespace[free_index].size -= files[index].size
				freespace[free_index].index += files[index].size
				break
			}
		}
	}
	for i := 0; i < 1000; i++ {
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
