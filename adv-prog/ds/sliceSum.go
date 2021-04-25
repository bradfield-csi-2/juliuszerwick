package main

import (
	"fmt"
	"unsafe"
)

// Given an []int slice, return the sum of values in the
// slice without using range or the [] operator.

func main() {
	var m int = 1
	mOffset := unsafe.Sizeof(m)

	ints := []int{1, 2, 3}
	//sum := 0
	num := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&ints))))
	//num := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&ints))))
	fmt.Printf("num: %v\n", num)
	fmt.Println()

	for i := 0; i < len(ints); i++ {
		fmt.Printf("i: %d\n", i)
		// Offset will be 4 bytes times the current index.
		offset := uintptr(i) * mOffset
		intsPtr := uintptr(unsafe.Pointer(&ints))
		fmt.Printf("intsPtr: %v\n", intsPtr)

		v := intsPtr + offset
		fmt.Printf("v: %v\n", v)

		num := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&ints)) + offset))
		fmt.Printf("num: %v\n", num)
		fmt.Println()
	}

	//fmt.Printf("sum: %d\n", sum)
}
