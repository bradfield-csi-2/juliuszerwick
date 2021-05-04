package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	I'm not sure I see what the bug is here. Could it be
	that we should only print one of the responses?
*/

var responses = []string{
	"200 OK",
	"402 Payment Required",
	"418 I'm a teapot",
}

func randomDelay(maxMillis int) time.Duration {
	return time.Duration(rand.Intn(maxMillis)) * time.Millisecond
}

func query(endpoint string) string {
	// Simulate querying the given endpoint
	delay := randomDelay(100)
	time.Sleep(delay)

	i := rand.Intn(len(responses))
	return responses[i]
}

// Query each of the mirrors in parallel and return the first
// response (this approach increases the amount of traffic but
// significantly improves "tail latency")
func parallelQuery(endpoints []string) string {
	results := make(chan string)
	for i := range endpoints {
		go func(i int) {
			results <- query(endpoints[i])
			//fmt.Println("In parallelQuery")
		}(i)
	}
	//fmt.Println("before <-results")
	return <-results
}

func main() {
	var endpoints = []string{
		"https://fakeurl.com/endpoint",
		"https://mirror1.com/endpoint",
		"https://mirror2.com/endpoint",
	}

	// Simulate long-running server process that makes continuous queries
	for {
		fmt.Println(parallelQuery(endpoints))
		//	fmt.Println("Before delay in infinite for loop!")
		delay := randomDelay(100)
		time.Sleep(delay)
	}
}
