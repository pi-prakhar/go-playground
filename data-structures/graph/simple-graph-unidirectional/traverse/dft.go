package traverse

import (
	"fmt"
	"simple-graph-unidirectional/helpers"
)

func DepthFirstTraverse(graph map[int][]int, source int) {
	if len(graph) == 0 {
		return
	}

	var stack helpers.Stack[int]
	stack.Push(source)
	visited := make(map[int]struct{})

	for !stack.IsEmpty() {
		element, exists := stack.Pop()
		if _, ok := visited[element]; !ok && exists {
			fmt.Println(element)
			visited[element] = struct{}{}
			var neighbours []int = graph[element]
			for _, neighbour := range neighbours {
				stack.Push(neighbour)
			}
		}
	}
}
