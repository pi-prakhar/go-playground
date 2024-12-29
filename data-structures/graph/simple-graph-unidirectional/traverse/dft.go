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

	for !stack.IsEmpty() {
		element, exists := stack.Pop()
		if exists {
			fmt.Println(element)
			var neighbours []int = graph[element]
			for _, neighbour := range neighbours {
				stack.Push(neighbour)
			}
		}
	}
}
