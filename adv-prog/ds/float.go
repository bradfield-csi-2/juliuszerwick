package main

import (
	"fmt"
	"unsafe"
)

// Given a float64, return a uint64 with the same binary representation.
// Once you have a uint64, you can confirm the result is correct via fmt.Printf("%064b\n", x).

func main() {

	var f float64 = 1.0
	fmt.Printf("unsafe.Pointer:     %064b\n", unsafe.Pointer(&f))
	fb := floatBits(f)
	// Print bit pattern of float64
	fmt.Printf("uint64 bit pattern: %064b\n", fb)
	fmt.Printf("uint64 value:       %d\n", fb)
}

// floatBits will return the bit pattern for a float64 value.
func floatBits(f float64) uint64 {
	// Take f parameter and determine bit pattern.
	// Translate  bit pattern into a value of type uint64
	i := *(*uint64)(unsafe.Pointer(&f))
	//fmt.Printf("strconv.FormatUint: %v\n", strconv.FormatUint(i, 2))
	// Return new value
	return i
}
