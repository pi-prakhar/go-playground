package main

import (
	"fmt"
	"sync"
	"time"
)

func taskHandler(wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	for {
		task, ok := <-ch
		if !ok {
			fmt.Println("task handler quits!")
			return
		}
		time.Sleep(1 * time.Second)
		fmt.Println("Processed task : ", task)
	}

}

func main() {
	fmt.Println("Program Started!")
	tasks := make(chan int, 5)
	workers := 2
	var wg sync.WaitGroup

	for _ = range workers {
		wg.Add(1)
		go taskHandler(&wg, tasks)
	}

	go func() {
		for task := range 10 {
			tasks <- task
		}
		fmt.Println("exited")
		close(tasks)
	}()

	wg.Wait()

	fmt.Println("Program Exited!")
}
