package main

import (
	"fmt"
	"sync/atomic"
)

/*
Try to implement a mutex yourself from lower-level concurrency primitives.
You should provide the same interface as sync.Mutex (Lock and Unlock), but obviously do not
just call sync.Mutex under the hood.
Test your implementation by using it to guard concurrent access to a counter variable.
*/

type mutex struct {
	locked uint32
}

func (m *mutex) Lock() {
	// Check if mutex is already locked.
	// If false, update the locked field value and return.
	for {
		if atomic.CompareAndSwapUint32(&m.locked, 0, 1) {
			return
		}
		//fmt.Println("mutex is already locked!")
	}
}

func (m *mutex) Unlock() {
	for {
		if atomic.CompareAndSwapUint32(&m.locked, 1, 0) {
			return
		}
		//fmt.Println("mutex is already unlocked!")
	}
}

// Question: How do we make the Lock() and Unlock()
//					 operations prevent access to the value
//           by other goroutines/function calls?
func increment(m *mutex, i *uint64) {
	m.Lock()
	*i += 1
	m.Unlock()
}

func main() {
	m := new(mutex)
	i := uint64(1)
	fmt.Printf("i = %d\n", i)

	increment(m, &i)
	fmt.Printf("i = %d\n", i)
}
