package main

/*
int square(int num) {
	return num * num;
}
*/
import "C"

// We can define C functions within comments like those above.
// The pseudo-package C must be on a separate import line and
// placed directly after the commented out C code.
import "fmt"

func main() {
	n := 2

	// We need to cast our Go int value into a C int value with C.int()
	// To properly store and handle the int value returned by the C
	// function, we need to cast it to a Go string with int()
	sq := int(C.square(C.int(n)))
	fmt.Printf("Square of %d is %d\n", n, sq)
}
