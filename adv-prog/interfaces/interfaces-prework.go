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

	fmt.Println()

	// Exercise 2
	v := Val{num: 5}
	//v.PrintNum()
	b := Blub(v)
	printMethods(b)
}

/*
Exercise 1:

Given an interface{} variable that holds an int value, write a function
that extracts the int value without using a type assertion or type switch.
*/

type iface struct {
	tab  *itab
	data unsafe.Pointer
}

type itab struct {
	inter unsafe.Pointer
	_type unsafe.Pointer
	hash  uint32 // copy of _type.hash. Used for type switches.
	_     [4]byte
	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}

func extractInt(inter interface{}) int {
	i := (*iface)(unsafe.Pointer(&inter))
	d := *(*int)(i.data)

	return d
}

/*
Exercise 2:

Given an arbitrary interface value, write a function that iterates through
the corresponding itable and prints out information about methods that
extracts the int value without using a type assertion or type switch.
*/

type Blub interface {
	PrintNum()
	SayHello()
}

type Val struct {
	num int
}

func (v Val) PrintNum() {
	fmt.Println(v.num)
}

func (v Val) SayHello() {
	fmt.Println("Hello!")
}

func printMethods(inter interface{}) {
	//	t, ok := inter.(Val)
	//	fmt.Printf("t: %v\nok: %v\n", t, ok)

	i := (*iface)(unsafe.Pointer(&inter))
	it := i.tab

	fmt.Printf("i: %v\n", i)
	fmt.Printf("*i: %v\n", *i)
	fmt.Printf("it: %v\n", it)
	fmt.Printf("it.fun[0]: %v\n", it.fun[0])

	fmt.Println()

	f := it.fun[0]
	fp := unsafe.Pointer(f)
	fmt.Printf("fp: %v\n", fp)
	//fmt.Printf("fp string: %v\n", *(*string)(fp))
}

/*
Exercise 3:

Now that you know how interfaces are represented in memory, how do you
think “type assertions” and “type switches” work?
If you were designing Go yourself, how would you approach these features?


Answer:
I think that a type assertion would access the first word in an interface value's
memory representation, follow the pointer to access the itable, and then check the type stored in the itable.

And with a type switch, the return value of the initial type assertion is used to comparewith each of the cases. Each case would be a different type and if the type matches the return value of the type assertion then that case's logic would be executed.
*/
