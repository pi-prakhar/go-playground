package traverse

import (
	"simple-graph-unidirectional/helpers"
)


func HasPathDFT(graph map[int][]int, src int, dst int) bool {
	if (src == dst){
		return true
	}
	
	visited := make(map[int]struct{})
	var stack helpers.Stack[int]
	
	stack.Push(src)
	for !stack.IsEmpty() {
		element, _ := stack.Pop()
		if _, ok := visited[element]; !ok {
			if element == dst {
				return true
			}
			visited[element] = struct{}{}
			neighbours := graph[element]
			for _, neighbour := range neighbours {
				stack.Push(neighbour)
			}
		}
	}

	return false
}
		


