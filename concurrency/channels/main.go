package main

import "fmt"

func main() {
	var queue Queue
	queue.InitQueue(10)

	for i := 0; i < 11; i++ {
		if ok := queue.Enqueue(i); !ok {
			fmt.Println("queue is full")
		}
	}

	for i := 0; i < 14; i++ {
		if val, ok := queue.Dequeue(); ok {
			fmt.Println(val)
		} else {
			fmt.Println("queue is empty")
		}
	}
}
