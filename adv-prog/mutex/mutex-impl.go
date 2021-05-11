package main

import "fmt"

/*
Try to implement a mutex yourself from lower-level concurrency primitives.
You should provide the same interface as sync.Mutex (Lock and Unlock), but obviously do not
just call sync.Mutex under the hood.
Test your implementation by using it to guard concurrent access to a counter variable.
*/

type mutex struct {
}

func (m *mutex) Lock() {
}

func (m *mutex) Unlock() {
}

// Question: How do we make the Lock() and Unlock()
//					 operations prevent access to the value
//           by other goroutines/function calls?
func increment(m *mutex) {
}

func main() {
	m := new(mutex)
	m.value = uint64(1)
	fmt.Printf("m.value = %d\n", m.value)

	increment(m)
	fmt.Printf("m.value = %d\n", m.value)
}
