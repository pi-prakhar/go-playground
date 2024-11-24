package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func unoptimizedWork() int {
	sum := 0
	for i := 0; i < 100000000; i++ {
		// Using rand.Intn() inside the loop which is expensive
		sum += rand.Intn(10)
	}
	return sum
}

func optimizedWork() int {
	sum := 0
	randomNums := make([]int, 100000000)

	// Pre-generate random numbers
	for i := 0; i < 100000000; i++ {
		randomNums[i] = rand.Intn(10)
	}

	// Use the pre-generated random numbers in the loop
	for _, num := range randomNums {
		sum += num
	}
	return sum
}

var data = make([][]byte, 0)

func memoryLeaker() {
	for {
		// Allocate memory and append to the global slice (memory leak)
		block := make([]byte, 1024*1024) // 1 MB allocation
		data = append(data, block)

		// Simulate some work
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// Start the memory leaker in a goroutine
	go memoryLeaker()

	// Expose pprof at http://localhost:6060/debug/pprof/
	fmt.Println("pprof available at http://localhost:6060/debug/pprof/")
	http.ListenAndServe(":6060", nil)
}
