package main

import (
	"fmt"
	"time"
)

// func SimulateRequest(requestDuration time.Duration, totalDuration time.Duration, tokenConsumption int, tb *TokenBucket) {
// 	ticker := time.NewTicker(requestDuration)
// 	stopper := time.After(totalDuration)
// 	defer ticker.Stop()
// 	for {
// 		select {
// 		case <-stopper:
// 			return
// 		case <-ticker.C:
// 			tokensLeft, status := tb.UseToken(tokenConsumption)
// 			if status {
// 				fmt.Println("Successfull Request!!, tokens left : ", tokensLeft)
// 			} else {
// 				fmt.Println("Failed Request!!, Rate Limited")
// 			}
// 		}
// 	}
// }

func SimulateRequest(requestDuration time.Duration, totalDuration time.Duration, tokenConsumption int, sw *SlidingWindow) {
	ticker := time.NewTicker(requestDuration)
	stopper := time.After(totalDuration)
	defer ticker.Stop()
	for {
		select {
		case <-stopper:
			return
		case <-ticker.C:
			tokensLeft, status := sw.UseToken()
			if status {
				fmt.Println("Successfull Request!!, tokens left : ", tokensLeft)
			} else {
				fmt.Println("Failed Request!!, Rate Limited")
			}
		}
	}
}

func main() {
	fmt.Println("Testing")
	//tb := NewTokenBucket(5, 1*time.Second)
	//SimulateRequest(500*time.Millisecond, 10*time.Second, 1, newTB)

	sw := NewSlidingWindow(5*time.Second, 5)
	SimulateRequest(500*time.Millisecond, 10*time.Second, 1, sw)
}
