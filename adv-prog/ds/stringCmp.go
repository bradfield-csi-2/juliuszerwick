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
	//a := ""
	//b := "What in the world is going on?"
	//c := "Hello"
	d := a

	// Should return false.
	//fmt.Println(cmpStrings(a, b))

	// Should return false.
	//fmt.Println(cmpStrings(a, c))

	// Should return false.
	fmt.Println(cmpStrings(a, d))
}

func cmpStrings(a, b string) bool {
	// Converting a Pointer to a uintptr produces the memory
	// address of the value pointed at, as an integer.

	// However, the address of a string really returns the address
	// of the Data pointer in the StringHeader struct. The address
	// of what the Data pointer points to is what you need.
	pa := uintptr(unsafe.Pointer(&a))
	//fmt.Printf("pa: %v\n", pa)
	fmt.Printf("pa: %064b\n", pa)

	pb := uintptr(unsafe.Pointer(&b))
	//fmt.Printf("pb: %v\n", pb)
	fmt.Printf("pb: %064b\n", pb)

	return false
}
