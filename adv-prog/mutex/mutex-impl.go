package main

import (
	"fmt"
	"sync"
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
	for {
		if atomic.CompareAndSwapUint32(&m.locked, 0, 1) {
			return
		}
	}
}

func (m *mutex) Unlock() {
	for {
		if atomic.CompareAndSwapUint32(&m.locked, 1, 0) {
			return
		}
	}
}

func increment(m *mutex, wg *sync.WaitGroup, i *uint64) {
	defer wg.Done()
	m.Lock()
	*i += 1
	m.Unlock()
}

func main() {
	m := new(mutex)
	n := uint64(0)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(m, &wg, &n)
	}

	wg.Wait()
	fmt.Printf("i = %d\n", n)
}
