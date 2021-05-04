package main

import (
	"fmt"
	"sync"
)

const (
	numGoroutines = 100
	numIncrements = 100
)

/*
	The bug here was that the mutex variable was not a true global mutex.
	Because it was initialized in the main function and passed into the
	safeIncrement function, each call to safeIncrement was handling a local
	copy of the mutex.

	Thus, all of the goroutines were still reading and updating the count
	value without any mutual exclusion prevent data races between them.

	To fix this, we need to declare the mutex in the global scope and remove
	it from veing an argument in the safeIncrement function signature.
*/
var globalLock sync.Mutex

type counter struct {
	count int
}

func safeIncrement(c *counter) {
	globalLock.Lock()
	defer globalLock.Unlock()

	c.count += 1
}

func main() {
	c := &counter{
		count: 0,
	}

	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for j := 0; j < numIncrements; j++ {
				safeIncrement(c)
			}
		}()
	}

	wg.Wait()
	fmt.Println(c.count)
}
