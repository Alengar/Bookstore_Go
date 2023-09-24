package config

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	baseURL     = "http://localhost:8080"
	numRequests = 100
	concurrency = 10
)

func makeRequest(path string, wg *sync.WaitGroup, results chan time.Duration) {
	defer wg.Done()

	startTime := time.Now()

	resp, err := http.Get(baseURL + path)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	elapsedTime := time.Since(startTime)
	results <- elapsedTime
}

func main() {
	var wg sync.WaitGroup
	results := make(chan time.Duration, numRequests)

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go makeRequest("/books", &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var totalResponseTime time.Duration
	var numCompletedRequests int

	for elapsed := range results {
		totalResponseTime += elapsed
		numCompletedRequests++
	}

	avgResponseTime := totalResponseTime / time.Duration(numCompletedRequests)
	fmt.Printf("Total Requests: %d\n", numRequests)
	fmt.Printf("Completed Requests: %d\n", numCompletedRequests)
	fmt.Printf("Average Response Time: %v\n", avgResponseTime)
}
