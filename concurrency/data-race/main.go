package main

import (
	"fmt"
	"sync"
)

func dataRaceFixedUnBufferedChannel() {
	//trying to acheive using wait group
	//TODO : if in future found to implement without wait group fix this problem
	var counter int = 0
	var wg sync.WaitGroup
	counterChan := make(chan int)
	doneChan := make(chan bool)
	fmt.Println("Program started")
	go func() {
		fmt.Println("incrementing started")
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				// time.Sleep(1 * time.Second)
				counterChan <- 1
				fmt.Println("incrementing 1")
			}()
		}
		wg.Wait()
		fmt.Println("Incrementing stopped")
		close(counterChan)
	}()

	go func() {
		fmt.Println("Counting started")
		for {
			count, isOpen := <-counterChan
			if isOpen {
				counter = counter + count
				// fmt.Println("counter : ", counter)
				// time.Sleep(1 * time.Second)
			} else {
				fmt.Println("counting stopped")
				doneChan <- true
				return
			}
		}
	}()

	<-doneChan
	fmt.Println("Final counter value ", counter)
	fmt.Println("Program finished")

}

func dataRaceFixedBufferedChannel() {
	var counter int = 0
	counterChan := make(chan int, 100)
	doneChan := make(chan bool)
	fmt.Println("Program started")
	var wg sync.WaitGroup
	go func() {
		fmt.Println("incrementing started")
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				counterChan <- 1
			}()
		}
		wg.Wait()
		fmt.Println("Incrementing stopped")
		close(counterChan)
	}()

	go func() {
		for {
			count, isOpen := <-counterChan
			if isOpen {
				counter = counter + count
				fmt.Println("incrementing 1")
				fmt.Println("counter : ", counter)
			} else {
				fmt.Println("counting stopped")
				doneChan <- true
			}
		}

	}()
	<-doneChan
	fmt.Println("Final counter value ", counter)
	fmt.Println("Program finished")

}
func dataRaceProgramFixed() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup
	const numGoroutines = 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(c *int, mu *sync.Mutex) {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			counter++
		}(&counter, &mu)
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}

func dataRaceProgram() {
	var counter int
	var wg sync.WaitGroup

	const numGoroutines = 1000

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(c *int) {
			defer wg.Done()
			counter++
		}(&counter)
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}

func demonstrateMapRace() {
	// m := make(map[int]int)
	var wg sync.WaitGroup
	var m sync.Map

	// Multiple goroutines writing to map
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			m.Store(val, val) // Data race here!
		}(i)
	}
	wg.Wait()
	fmt.Println(m)
}

func main() {
	// dataRaceProgram()
	// demonstrateMapRace()
	dataRaceFixedBufferedChannel()
	// dataRaceFixedUnBufferedChannel()
	// dataRaceProgramFixed()
}
