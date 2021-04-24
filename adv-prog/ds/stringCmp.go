package main

import (
	"fmt"
	"unsafe"
)

// Given two strings, return a boolean that indicates whether
// the underlying string data is stored at the same memory
// location.
func main() {
	a := "Hello"
	b := "What in the world is going on?"
	c := "Hello"
	d := a
	e := "hello"

	// Should return false.
	fmt.Println(cmpStrings(a, b))

	// Should return false.
	// But returns true??
	// Does Go recognize different instances of the same
	// string value and place the data to the same address?
	fmt.Println(cmpStrings(a, c))

	// Should return true.
	fmt.Println(cmpStrings(a, d))

	// Should return false.
	fmt.Println(cmpStrings(a, e))
}

func cmpStrings(a, b string) bool {
	// Converting a Pointer to a uintptr produces the memory
	// address of the value pointed at, as an integer.

	// However, the address of a string really returns the address
	// of the Data pointer in the StringHeader struct. The address
	// of what the Data pointer points to is the actual address
	// of the string data.
	// The runtime pkg contains info on the StringHeader struct.
	pa := unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&a)))
	fmt.Printf("pa: %v\n", pa)
	//fmt.Printf("pa: %064b\n", pa)

	pb := unsafe.Pointer(*(*uintptr)(unsafe.Pointer(&b)))
	fmt.Printf("pb: %v\n", pb)
	//fmt.Printf("pb: %064b\n", pb)

	return pa == pb
}
