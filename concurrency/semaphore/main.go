package main

import (
	"fmt"
	"sync"
	"time"
)

func taskHandler(task int, wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()
	ch <- struct{}{}
	time.Sleep(1 * time.Second)
	fmt.Println("Processed task : ", task)
	<-ch
}

func main() {
	fmt.Println("Program Started!")
	tasks := 10
	semaphore := make(chan struct{}, 2)
	var wg sync.WaitGroup

	for task := range tasks {
		wg.Add(1)
		go taskHandler(task, &wg, semaphore)
	}

	wg.Wait()
	close(semaphore)
	fmt.Println("Program Exited!")
}
