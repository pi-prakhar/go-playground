package contextbasics

import (
	"context"
	"fmt"
	"time"
)

func DeadlineMain() {
	deadline := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel() // Ensure cancel is called to release resources

	go deadlineWorker(ctx)

	// Simulate doing some work in the main function
	time.Sleep(3 * time.Second)
	fmt.Println("Main function completed")
}

func deadlineWorker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker: context deadline exceeded or canceled:", ctx.Err())
			return
		default:
			fmt.Println("Worker: working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
