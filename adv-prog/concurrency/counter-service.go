package main

/*
Consider the following interface for an “ID service”:

type counterService interface {
    // Returns values in ascending order; it should be safe to call
    // getNext() concurrently without any additional synchronization.
    getNext() uint64
}

Implement this interface using each of the following four strategies:
- Don’t perform any synchronization
- Atomically increment a counter value using sync/atomic
- Use a sync.Mutex to guard access to a shared counter value
- Launch a separate goroutine with exclusive access to a private counter value; handle getNext() calls by making “requests” and receiving “responses” on two separate channels

Aside from the first (obviously incorrect) strategy, ensure that your implementations are correct by making sure that:
- In the context of a particular goroutine making calls to getNext(), returned values are monotonically increasing
- The maximum value returned by getNext() matches the total number of calls across all goroutines
- Go’s race detector doesn’t detect any race conditions

How do you expect these different strategies to compare in terms of performance? What are the bottlenecks in each case?

*/
