package main

import (
	"fmt"
	"sync"
	"time"
)

type dbService struct {
	lock       *sync.RWMutex
	connection string
}

func newDbService(connection string) *dbService {
	return &dbService{
		lock:       &sync.RWMutex{},
		connection: connection,
	}
}

func (d *dbService) logState() {
	d.lock.RLock()
	defer d.lock.RUnlock()

	fmt.Printf("connection %q is healthy\n", d.connection)
}

// The bug here is nested locks (reentrant locks?) which
// lead to a deadlock because the lock is never released.
// To fix this bug, we need to explicitly call d.lock.RUnlock()
// after the read is complete and before the call to d.logState()
func (d *dbService) takeSnapshot() {
	d.lock.RLock()
	//defer d.lock.RUnlock()

	fmt.Printf("Taking snapshot over connection %q\n", d.connection)

	d.lock.RUnlock()

	// Simulate slow operation
	time.Sleep(time.Second)

	d.logState()
}

func (d *dbService) updateConnection(connection string) {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.connection = connection
}

func main() {
	d := newDbService("127.0.0.1:3001")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		d.takeSnapshot()
	}()

	// Simulate other DB accesses
	time.Sleep(200 * time.Millisecond)

	wg.Add(1)
	go func() {
		defer wg.Done()

		d.updateConnection("127.0.0.1:8080")
	}()

	wg.Wait()
}
