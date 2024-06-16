package main

import (
	"fmt"
	"sync"
)

func printSomethingWithWait(s string, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Println(s)
}

func WaitGroup() {
	var wg sync.WaitGroup

	wg.Add(7)

	words := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
	}

	for i, x := range words {
		go printSomethingWithWait(fmt.Sprintf("%d : %s", i, x), &wg)
	}

	wg.Wait()
	fmt.Println("Printed!!")

}
