package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var inter interface{}
	inter = 5

	// Exercise 1
	i := extractInt(inter)
	fmt.Printf("inter data: %v\n", i)
}

/*
Exercise 1:

Given an interface{} variable that holds an int value, write a function
that extracts the int value without using a type assertion or type switch.
*/

type iface struct {
	tab  unsafe.Pointer
	data unsafe.Pointer
}

func extractInt(inter interface{}) int {
	i := (*iface)(unsafe.Pointer(&inter))
	d := *(*int)(i.data)

	return d
}
