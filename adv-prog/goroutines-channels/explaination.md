# Prompt

Investigate the steps involved when a goroutine attempts to send on an unbuffered channel and then blocks.
Try to describe the process in as much detail as possible.

You may find it helpful to consult chan.go and proc.go in the runtime package of the Go source code.

## Channel creation

When an unbuffered channel is created using the make built-in, the value returned is a pointer to an `hchan` struct which represents
the channel's structure. There are several fields in an `hchan` struct, and the ones we want to focus on are the following:

- `buf`: a buffer that will keep copies of the data sent over the channel
- `recevq`: a struct that stores the goroutines waiting to receive from the channel
- `sendq`: a struct that stores the goroutines waiting to send over the channel
- `lock`: a mutex that locks the buffer when a send operation copies a value into the buffer or when a receive operation reads + removes a value from the buffer

## Send on an unbuffered channel

When a goroutine attempts to send a value on an unbuffered channel, the channel will perform the following steps:

- use the `hchan.lock` to acquire a lock on the buffer
- copy the value to the buffer (enqueue the value)
- release the lock on the buffer

## Blocks


