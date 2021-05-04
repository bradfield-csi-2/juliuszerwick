package main

import (
	"fmt"
)

const numTasks = 3

func main() {
	/*
		The channel below is a nil channel.
		All send and receive operations on a nil channel will block,
		thus the code cannot continue past the send operation in
		each goroutine.

		To fix this, we need to use the make() built-in to create
		an unbuffered channel that is ready to use.
	*/

	//var done chan struct{}
	done := make(chan struct{})
	for i := 0; i < numTasks; i++ {
		go func() {
			fmt.Println("running task...")

			// Signal that task is done
			done <- struct{}{}
		}()
	}

	// Wait for tasks to complete
	for i := 0; i < numTasks; i++ {
		<-done
	}
	fmt.Printf("all %d tasks done!\n", numTasks)
}
