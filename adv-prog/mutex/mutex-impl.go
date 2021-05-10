package main

import "fmt"

/*
Try to implement a mutex yourself from lower-level concurrency primitives.
You should provide the same interface as sync.Mutex (Lock and Unlock), but obviously do not
just call sync.Mutex under the hood.
Test your implementation by using it to guard concurrent access to a counter variable.
*/

type mutex struct {
	locked bool
}

func (m *mutex) Lock() {
	// Check if mutex is locked.
	for {
		if !m.locked {
			m.locked = true
			return
		}
		fmt.Println("mutex is already locked!")
	}
	// If false, update mutex state to locked.
	// Else, block.
}

func (m *mutex) Unlock() {
	// Check if mutex is locked.
	for {
		if m.locked {
			m.locked = false
			return
		}
		fmt.Println("mutex is not locked!")
	}
	// If true, update mutex state to unlocked.
	// Else, block.
}

func increment(i *int, m *mutex) {
	m.Lock()
	*i += 1
	m.Unlock()
}

func main() {
	m := new(mutex)
	x := 1
	fmt.Printf("x = %d\n", x)

	increment(&x, m)
	fmt.Printf("x = %d\n", x)
}
