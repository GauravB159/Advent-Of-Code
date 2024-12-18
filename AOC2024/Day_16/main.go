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
	graph := make(map[Node]map[Node]int)
	distances := make(map[Node]int)
	parent := make(map[Node]Node)
	directions := [4]aocutils.Key{{Row: 0, Col: 1}, {Row: 0, Col: -1}, {Row: 1, Col: 0}, {Row: -1, Col: 0}}
	for position, char := range grid.Data {
		grid_image.SetZoomedPixel(position.Col, position.Row, 1)
		for _, direction := range directions {
			node := Node{position: position, direction: direction}
			for _, nextDirection := range directions {
				var increase int
				nextNode := Node{position: aocutils.Key{Row: node.position.Row + direction.Row, Col: node.position.Col + direction.Col}, direction: nextDirection}
				if grid.Data[node.position] == '#' {
					continue
				}
				if _, exists := graph[node]; !exists {
					graph[node] = make(map[Node]int, 0)
				}
				if nextDirection.Row == node.direction.Row && nextDirection.Col == node.direction.Col {
					if grid.Data[nextNode.position] == '#' {
						continue
					}
					increase = 1
				} else if -1*nextDirection.Row == node.direction.Row && -1*nextDirection.Col == node.direction.Col {
					continue
				} else {
					nextNode = Node{position: aocutils.Key{Row: node.position.Row, Col: node.position.Col}, direction: nextDirection}
					increase = 1000
				}
				if _, exists := grid.Data[nextNode.position]; !exists {
					continue
				}
				graph[node][nextNode] = increase
			}
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
			// start = Node{position: aocutils.Key{Row: 3, Col: 4}, direction: aocutils.Key{Row: 0, Col: -1}}
		}
		if char == 'E' {
			grid_image.SetZoomedPixel(position.Col, position.Row, 7)
			end = position
		}
	}
	fmt.Println(len(graph))

	// for node, connections := range graph {
	// 	for connection, distance := range connections {
	// 		if connection.direction.Row != node.direction.Row || connection.direction.Col != node.direction.Col {
	// 			continue
	// 		}
	// 		new_connection := Node{
	// 			position:  node.position,
	// 			direction: node.direction,
	// 		}
	// 		count := distance
	// 		var next_connection Node
	// 		for true {
	// 			next_connection = Node{position: aocutils.Key{Row: new_connection.position.Row + start.direction.Row, Col: new_connection.position.Col + start.direction.Col}, direction: start.direction}
	// 			// fmt.Println(graph[new_connection])
	// 			if increase, exists := graph[new_connection][next_connection]; exists {
	// 				delete(graph[new_connection], next_connection)
	// 				new_connection = next_connection
	// 				count += increase
	// 			} else {
	// 				break
	// 			}
	// 		}
	// 		graph[connection][new_connection] = count
	// 		// fmt.Println(connection, count, new_connection, distance)
	// 	}
	// }
	count := 0
	for _, value := range graph {
		if len(value) == 0 {
			count += 1
		}
	}
	fmt.Println(count)
	fmt.Println(len(graph))

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
		for nextNode, increase := range graph[node] {
			score := distances[node] + increase
			// fmt.Println(score, node, nextNode)
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
func dfs(node aocutils.Key, grid *aocutils.Grid, end aocutils.Key, prevDirection aocutils.Key, visited *map[Node]bool, runningScore int) int {
	key := Node{position: node, direction: prevDirection}
	if (*grid).Data[node] == '#' {
		return 100000000
	}
	if node.Row == end.Row && node.Col == end.Col {
		// fmt.Println("HERE2")
		if runningScore == 11048 {
			fmt.Println("HERE3")
			fmt.Println()
		}
		return 0
	}

	(*visited)[key] = true
	directions := [4]aocutils.Key{{Row: 0, Col: 1}, {Row: 0, Col: -1}, {Row: 1, Col: 0}, {Row: -1, Col: 0}}
	minscore := math.MaxInt64
	for _, direction := range directions {
		nextNode := aocutils.Key{Row: node.Row + direction.Row, Col: node.Col + direction.Col}
		if (*visited)[Node{position: nextNode, direction: direction}] {
			continue
		}
		increase := 1
		if direction.Row*-1 == prevDirection.Row && direction.Col*-1 == prevDirection.Col {
			continue
		}
		if direction.Row != prevDirection.Row || direction.Col != prevDirection.Col {
			increase = 1001
		}
		if runningScore+increase > 101492 {
			continue
		}
		checkscore := increase + dfs(nextNode, grid, end, direction, visited, runningScore+increase)
		if checkscore < minscore {
			minscore = checkscore
		}
	}
	(*visited)[key] = false
	return minscore
}

func twostar(filename string) string {
	lines := aocutils.Readfile(filename)
	grid := aocutils.CreateCharacterGrid(lines)
	var start, end aocutils.Key
	for position, char := range grid.Data {
		if char == 'S' {
			start = position
		}
		if char == 'E' {
			end = position
		}
	}
	visited := make(map[Node]bool)
	score := dfs(start, &grid, end, aocutils.Key{Row: 0, Col: 1}, &visited, 0)
	return strconv.Itoa(score)
}

func main() {
	aocutils.Timer("1 star", onestar, "input.txt")
	fmt.Println()
	fmt.Println()
	// aocutils.Timer("2 star", twostar, "input.txt")
}
