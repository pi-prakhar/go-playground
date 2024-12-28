package main

import (
	"fmt"
	"graph/helpers"
)

func main() {
	var stack helpers.Stack[int]
	stack.Push(10)
	stack.Push(20)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())


	var queue helpers.Queue[int]
	fmt.Println(queue.IsEmpty())
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

}
