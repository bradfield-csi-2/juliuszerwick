package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Given an []int slice, return the sum of values in the
// slice without using range or the [] operator.

func main() {
	// The int type is made of 64-bits on a 64-bit machine.
	ints := []int{1, 2, 3, 4, 5}
	fmt.Printf("sum: %d\n", sliceSum(ints))
}

func sliceSum(ints []int) int {
	sum := 0
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&ints))
	sliceData := sliceHeader.Data

	for i := 0; i < len(ints); i++ {
		num := *(*int)(unsafe.Pointer(sliceData + uintptr(8*i)))
		sum += num
	}

	return sum
}
