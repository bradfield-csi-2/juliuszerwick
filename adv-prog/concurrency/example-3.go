package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var urls = []string{
		"https://bradfieldcs.com/courses/architecture/",
		"https://bradfieldcs.com/courses/networking/",
		"https://bradfieldcs.com/courses/databases/",
	}

	/*
		Same as with example-1, the loop iteration will complete before each of the
		goroutines have a chance to increment the WaitGroup. And by doing so, the main
		goroutine will reach wg.Wait() and observe that the counter is still zero.
		This leads the main goroutine to continue to the final fmt.Println() call
		without any of the url strings being printed.

		To fix this bug, we need to move the wg.Add(1) call out of the goroutine body
		and above the goroutine call within the loop. This will ensure that the loop
		increments the counter and have wg.Wait() actually wait until each goroutine finishes
		and decrements the counter.
	*/
	var wg sync.WaitGroup
	for i := range urls {
		wg.Add(1)
		go func(i int) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()

			_, err := http.Get(urls[i])
			if err != nil {
				panic(err)
			}

			fmt.Println("Successfully fetched", urls[i])
		}(i)
	}

	// Wait for all url fetches
	wg.Wait()
	fmt.Println("all url fetches done!")
}
