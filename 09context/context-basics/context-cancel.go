package contextbasics

import (
	"context"
	"fmt"
	"time"
)

func CancelMain() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure cancel is called to release resources

	go cancelWorker(ctx)

	time.Sleep(2 * time.Second)
	cancel() // Cancel the context to stop the worker

	// Give some time for the goroutine to finish
	time.Sleep(1 * time.Second)
	fmt.Println("Main function exiting")
}

func cancelWorker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker: context canceled")
			return
		default:
			fmt.Println("Worker: working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
