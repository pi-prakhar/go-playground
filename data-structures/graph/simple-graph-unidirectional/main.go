package main

import (
	"simple-graph-unidirectional/traverse"
)

var Graph map[int][]int = map[int][]int{
	1: []int{2, 3},
	2: []int{4},
	3: []int{5},
	4: []int{3},
	5: []int{},
	6: []int{5},
}

func main() {
	traverse.DepthFirstTraverse(Graph, 1)
}
