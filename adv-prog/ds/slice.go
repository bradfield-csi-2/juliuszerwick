package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := []int16{1, 2, 3}

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("SliceHeader: %+v\n", sh)

	shData := sh.Data
	fmt.Printf("Data: %+v\n", shData)

	// Correctly accesses first element in slice.
	el0 := *(*int16)(unsafe.Pointer(shData))
	fmt.Printf("el0: %+v\n", el0)

	// Access second element by adding offset of 2 bytes (16 bits).
	el1 := *(*int16)(unsafe.Pointer(shData + uintptr(2)))
	fmt.Printf("el1: %+v\n", el1)

	// Access third element by adding offset of 4 bytes (32 bits).
	el2 := *(*int16)(unsafe.Pointer(shData + uintptr(4)))
	fmt.Printf("el2: %+v\n", el2)
}
