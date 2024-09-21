package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func Base() {
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
		go printSomething(fmt.Sprintf("%d : %s", i, x))
	}
	go printSomething("Go routine printing!!")

	//sleep so that it gets enough time to fire go routine before the full program executs
	time.Sleep(1 * time.Second)
	fmt.Println("Printed!!")
}
