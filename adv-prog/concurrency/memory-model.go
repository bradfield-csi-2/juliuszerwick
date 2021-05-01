package main

/*
The main point of the article titled The Go Memory Model
is that there are no guarantees of the exact ordering of
reads and writes to the same variable by multiple goroutines
without the use of synchronization primitives.

If you use such primitives, then you can enforce some guarantees!
*/

var a, b int

func f() {
	a = 1
	b = 2
}

func g() {
	print(b)
	print(a)
}

func main() {
	go f()
	g()
}
