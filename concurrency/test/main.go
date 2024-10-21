package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

func sender(ch chan<- int, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Sent : ", i)
	ch <- i
}

func reciever(ch <-chan int, done chan<- struct{}) {
	go func() {
		for {
			select {
			case val, ok := <-ch:
				if !ok {
					fmt.Println("channel closed")
					done <- struct{}{}
					return
				}
				fmt.Println("Recieved : ", val)
			case <-time.After(2 * time.Second):
				fmt.Println("Timeout")
				done <- struct{}{}
				return
			}
		}
	}()

	fmt.Println("Reciever ended")
}

func main() {
	fmt.Println("Program started")

	f, err := os.Create("cpu_profile.prof")
	if err != nil {
		fmt.Println("Could not create CPU profile:", err)
		return
	}
	pprof.WriteHeapProfile(f)
	defer f.Close()

	ch := make(chan int)
	done := make(chan struct{})
	var wg sync.WaitGroup
	senders := 5

	reciever(ch, done)
	for i := range senders {
		wg.Add(1)
		go sender(ch, i, &wg)
		time.Sleep(1 * time.Second)

	}

	wg.Wait()
	close(ch)
	<-done
	fmt.Println("Program exited")
	time.Sleep(5 * time.Second)
}
