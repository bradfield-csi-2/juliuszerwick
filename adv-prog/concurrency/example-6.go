package main

import (
	"fmt"
	"sync"
)

type coordinator struct {
	lock   sync.RWMutex
	leader string
}

func newCoordinator(leader string) *coordinator {
	return &coordinator{
		lock:   sync.RWMutex{},
		leader: leader,
	}
}

func (c *coordinator) logState() {
	c.lock.RLock()
	defer c.lock.RUnlock()

	// Prints a single quoted character.
	fmt.Printf("leader = %q\n", c.leader)
}

func (c *coordinator) setLeader(leader string, shouldLog bool) {
	// The bug is due to the deferred c.lock.Unlock() which will execute at the end
	// of this function body.
	// This leads to situations in which if shouldLog == true there will be a call
	// to c.lock.RLock() in logState() without the necessary c.lock.Unlock() before
	// it.
	// The fix is to change the deferred c.lock.Unlock() into an explicit, non-deferred call
	// directly after the assignment operation is complete.
	c.lock.Lock()

	c.leader = leader

	c.lock.Unlock()

	if shouldLog {
		c.logState()
	}
}

func main() {
	c := newCoordinator("us-east")
	c.logState()
	c.setLeader("us-west", true)
}
