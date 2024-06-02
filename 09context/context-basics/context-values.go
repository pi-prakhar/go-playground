package contextbasics

import (
	"context"
	"fmt"
	"time"
)

type key string

const requestIDKey key = "requestID"

func ValueMain() {
	ctx := context.WithValue(context.Background(), requestIDKey, "12345")

	go valueWorker(ctx)

	// Simulate doing some work in the main function
	time.Sleep(2 * time.Second)
	fmt.Println("Main function completed")
}

func valueWorker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker: context canceled")
			return
		default:
			if requestID, ok := ctx.Value(requestIDKey).(string); ok {
				fmt.Println("Worker: working with Request ID:", requestID)
			} else {
				fmt.Println("Worker: no Request ID found")
			}
			time.Sleep(500 * time.Millisecond)
		}
	}
}
