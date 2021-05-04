package main

import (
	"fmt"
)

/*
	The bug here is that the expected order of operations is not guaranteed.
	With unbuffered channels, a send operation will always occur before a
	receive operation on the same channel.

	To fix this bug, we need to swap the locations of the send and receive
	operations. The goroutine should send the token (struct{}{}) over the channel
	and the main goroutine should receive the token. This ensure that the main
	goroutine is blocked until the send operations is complete.
*/

func main() {
	done := make(chan struct{}, 1)
	go func() {
		fmt.Println("performing initialization...")
		done <- struct{}{}
	}()

	<-done
	fmt.Println("initialization done, continuing with rest of program")
}
