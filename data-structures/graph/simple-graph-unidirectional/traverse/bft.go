package traverse

import (
	"fmt"
	"simple-graph-unidirectional/helpers"
)

func BreadthFirstTraverse(source int, graph map[int][]int) {
	if len(graph) == 0 {
		return
	}

	var queue helpers.Queue[int]
	queue.Enqueue(source)
	visited := make(map[int]struct{})

	for !queue.IsEmpty() {
		element, exists := queue.Dequeue()
		if _, ok := visited[element]; !ok && exists {
			fmt.Println(element)
			visited[element] = struct{}{}
			neighbours := graph[element]
			for _, neighbour := range neighbours {
				queue.Enqueue(neighbour)
			}
		}
	}
}
