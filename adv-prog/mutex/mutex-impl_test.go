package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
)

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

func mutexTest(b int) {
	m := new(mutex)
	n := uint64(0)
	var wg sync.WaitGroup

	for i := 0; i < b; i++ {
		wg.Add(1)
		go increment(m, &wg, &n)
	}

	wg.Wait()
	//fmt.Printf("i = %d\n", n)
}

func incrementRealMutex(m *sync.Mutex, wg *sync.WaitGroup, i *uint64) {
	defer wg.Done()
	m.Lock()
	*i += 1
	m.Unlock()
}

func realMutexTest(b int) {
	var m = &sync.Mutex{}
	n := uint64(0)
	var wg sync.WaitGroup

	for i := 0; i < b; i++ {
		wg.Add(1)
		go incrementRealMutex(m, &wg, &n)
	}

	wg.Wait()
	//fmt.Printf("i = %d\n", n)
}

func BenchmarkMutex10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mutexTest(10)
	}
}

func BenchmarkMutex100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mutexTest(100)
	}
}

func BenchmarkMutex1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		mutexTest(1000)
	}
}

func BenchmarkRealMutex10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		realMutexTest(10)
	}
}

func BenchmarkRealMutex100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		realMutexTest(100)
	}
}

func BenchmarkRealMutex1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		realMutexTest(1000)
	}
}
