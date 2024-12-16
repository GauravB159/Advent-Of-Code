package main

import (
	"container/heap"
	"fmt"
	"math"

	"strconv"

	aocutils "github.com/GauravB159/aoc-go-utils"
)

type Node struct {
	position  aocutils.Key
	direction aocutils.Key
}

type PriorityNode struct {
	node     Node
	distance int
}

type NodeHeap []PriorityNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(PriorityNode))
}

func (h *NodeHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func onestar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	var start Node
	var end aocutils.Key
	grid_image := aocutils.CreateImage(grid.Rows, grid.Cols, 10, 10, "onestar")
	grid_image.UseFullColors()
	gif := aocutils.CreateGIF("onestar", 500)
	visited := make(map[Node]bool)
	distances := make(map[Node]int)
	parent := make(map[Node]Node)
	directions := [4]aocutils.Key{{Row: 0, Col: 1}, {Row: 0, Col: -1}, {Row: 1, Col: 0}, {Row: -1, Col: 0}}
	for position, char := range grid.Data {
		grid_image.SetZoomedPixel(position.Col, position.Row, 1)
		for _, direction := range directions {
			distances[Node{position: position, direction: direction}] = math.MaxInt64
			visited[Node{position: position, direction: direction}] = false
			if char == '#' {
				grid_image.SetZoomedPixel(position.Col, position.Row, 3)
				visited[Node{position: position, direction: direction}] = true
			}
		}
		if char == 'S' {
			grid_image.SetZoomedPixel(position.Col, position.Row, 5)
			start = Node{position: position, direction: aocutils.Key{Row: 0, Col: 1}}
		}
		if char == 'E' {
			grid_image.SetZoomedPixel(position.Col, position.Row, 7)
			end = position
		}
	}
	h := &NodeHeap{}
	heap.Init(h)
	heap.Push(h, PriorityNode{node: start, distance: 0})
	gif.AddFrame(grid_image)
	distances[start] = 0
	parent[start] = Node{position: aocutils.Key{Row: -1, Col: -1}, direction: aocutils.Key{Row: 0, Col: 0}}
	for h.Len() > 0 {
		node := heap.Pop(h).(PriorityNode).node
		(visited)[node] = true
		grid_image.SetZoomedPixel(node.position.Col, node.position.Row, 6)
		gif.AddFrame(grid_image)
		if node.position.Col == end.Col && node.position.Row == end.Row {
			break
		}
		for _, direction := range directions {
			nextNode := Node{position: aocutils.Key{Row: node.position.Row + direction.Row, Col: node.position.Col + direction.Col}, direction: direction}
			var increase int
			if direction.Row == node.direction.Row && direction.Col == node.direction.Col {
				increase = 1
			} else if -1*direction.Row == node.direction.Row && -1*direction.Col == node.direction.Col {
				continue
			} else {
				nextNode = Node{position: aocutils.Key{Row: node.position.Row, Col: node.position.Col}, direction: direction}
				increase = 1000
			}
			if visited[nextNode] {
				continue
			}
			if _, exists := grid.Data[nextNode.position]; !exists {
				visited[nextNode] = true
				continue
			}
			score := distances[node] + increase
			if score < distances[nextNode] {
				heap.Push(h, PriorityNode{node: nextNode, distance: score})
				distances[nextNode] = score
				parent[nextNode] = node
			}
		}
	}
	answer := math.MaxInt64
	var end_node Node
	for _, direction := range directions {
		dist := distances[Node{position: end, direction: direction}]
		if dist < answer {
			end_node = Node{position: end, direction: direction}
			answer = dist
		}
	}
	path := make([]Node, 0)
	pnode := end_node
	for true {
		if pnode.direction.Row == 0 && pnode.direction.Col == 0 {
			break
		}
		path = append(path, pnode)
		pnode = parent[pnode]
	}
	gif.AddFrame(grid_image)
	for i := len(path) - 1; i >= 0; i-- {
		grid_image.SetZoomedPixel(path[i].position.Col, path[i].position.Row, 9)
		for range 5 {
			gif.AddFrame(grid_image)
		}
	}
	for range 500 {
		gif.AddFrame(grid_image)
	}
	gif.WriteGIFToFile()
	return strconv.Itoa(answer)
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
