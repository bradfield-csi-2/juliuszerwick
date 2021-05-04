package main

import (
	"fmt"
	"time"
)

/*
The first bug here is that the loop iteration will likely finish before even the first goroutine has
a chance to finish execution. And because of this, the value of the variable i scoped to the
loop block will be set to 10 when the call to fmt.Printf() finishes.

To fix this bug, we simply need to pass the variable i into each goroutine.

The second "bug" involves the expectation that the goroutines will finish in the order they are called.
This is unlikely to happen given the random nature of how the underlying OS may context switch between
different goroutines.

To fix this so that the goroutines finish "in order" we need to add a time.Sleep() call to ensure that
each newly called goroutine has enough time to finish execution before the next goroutine is called.

Reference: https://dev.to/kkentzo/the-golang-for-loop-gotcha-1n35
*/

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("launched goroutine %d\n", i)
		}(i)
		time.Sleep(time.Millisecond)
	}
	// Wait for goroutines to finish
	time.Sleep(time.Second)
}
