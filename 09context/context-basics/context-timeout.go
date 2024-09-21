package contextbasics

import (
	"context"
	"fmt"
	"time"
)

func TimeoutMain() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Ensure cancel is called to release resources

	go timeoutWorker(ctx)

	// Simulate doing some work in the main function
	time.Sleep(2 * time.Second)
	fmt.Println("Main function completed")
}

func timeoutWorker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker: context timeout or canceled:", ctx.Err())
			return
		default:
			fmt.Println("Worker: working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
