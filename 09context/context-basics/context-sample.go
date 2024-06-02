package contextbasics

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// URL to fetch
var urls = []string{
	"https://www.google.com",
	"https://www.golang.org",
	"https://www.github.com",
}

// Result struct to hold the result of the fetch
type Result struct {
	URL  string
	Body string
	Err  error
}

// fetchContent fetches content from a given URL
func fetchContent(ctx context.Context, url string, ch chan<- Result) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		ch <- Result{URL: url, Err: err}
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- Result{URL: url, Err: err}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- Result{URL: url, Err: err}
		return
	}

	ch <- Result{URL: url, Body: string(body)}
}

func ContextSample() {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure cancel is called to release resources

	resultCh := make(chan Result)
	doneCh := make(chan struct{})

	// Start a worker goroutine for each URL
	for _, url := range urls {
		go fetchContent(ctx, url, resultCh)
	}

	// Collect results
	go func() {
		for range urls {
			select {
			case res := <-resultCh:
				if res.Err != nil {
					fmt.Printf("Error fetching %s: %v\n", res.URL, res.Err)
				} else {
					fmt.Printf("Fetched content from %s: %s\n", res.URL, res.Body[:100]) // Print first 100 characters
				}
			case <-ctx.Done():
				fmt.Println("Timeout: stopping fetch operations")
				doneCh <- struct{}{}
				return
			}
		}
		doneCh <- struct{}{}
	}()

	// Wait for results or timeout
	<-doneCh
	fmt.Println("Program finished")
}
