package main

import "fmt"

func main() {
	m := make(map[int]int)

	for i := 0; i < 6; i++ {
		m[i] = i
	}

	//fmt.Printf("%+v\n", m)

	ksum, vsum := sumsOfMap(m)
	fmt.Printf("sum of keys: %d\nsum of values: %d\n", ksum, vsum)
}

func sumsOfMap(m map[int]int) (int, int) {

	return 0, 0
}
