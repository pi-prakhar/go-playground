package main

import (
	"fmt"
	"time"
)

func SimulateRequest(requestDuration time.Duration, totalDuration time.Duration, tokenConsumption int, tb *TokenBucket) {
	ticker := time.NewTicker(requestDuration)
	stopper := time.After(totalDuration)
	defer ticker.Stop()
	for {
		select {
		case <-stopper:
			return
		case <-ticker.C:
			tokensLeft, status := tb.Usetokens(tokenConsumption)
			if status {
				fmt.Println("Successfull Request!! tokens left : ", tokensLeft)
			} else {
				fmt.Println("Failed Request !!, Rate Limited")
			}
		}
	}
}

func main() {
	fmt.Println("Testing")
	newTB := NewTokenBucket(5, 1*time.Second)
	SimulateRequest(500*time.Millisecond, 10*time.Second, 1, newTB)

}
