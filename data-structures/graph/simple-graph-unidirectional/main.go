package main

import (
	"fmt"
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
	fmt.Println("Depth First Traversal")
	traverse.DepthFirstTraverse(Graph, 1)
	fmt.Println("Breadth First Traversal")
	traverse.BreadthFirstTraverse(1, Graph)
	fmt.Println("Check has Path 1 -> 6",traverse.HasPathDFT(Graph, 1, 6))
	fmt.Println("Check has Path 6 -> 5",traverse.HasPathDFT(Graph, 6, 5))
	fmt.Println("Check has Path 4 -> 5",traverse.HasPathDFT(Graph, 4, 5))
	fmt.Println("Check has Path 1 -> 5",traverse.HasPathDFT(Graph, 1, 5)) 
}
