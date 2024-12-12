package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		log.Println("server started")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}(&wg)

	wg.Wait()
}
