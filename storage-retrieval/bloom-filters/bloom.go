package main

import (
	"encoding/binary"
)

type bloomFilter interface {
	add(item string)

	// `false` means the item is definitely not in the set
	// `true` means the item might be in the set
	maybeContains(item string) bool

	// Number of bytes used in any underlying storage
	memoryUsage() int
}

/*
	Below is my implementation of a Bloom Filter.
*/

type myBloomFilter struct {
	data []uint8
}

func newBloomFilter() *myBloomFilter {
	return &myBloomFilter{
		data: make([]uint8, 1000),
	}
}

// Add item to BF.
func (b *myBloomFilter) add(item string) {
	// Take item argument and extract a number for hashing.
	r := []rune(item)

	sum := 0
	for _, i := range r {
		sum += int(i)
	}

	l := len(b.data)

	// Hash it once.
	h1 := hashOne(sum, l)

	// Hash it twice.
	h2 := hashTwo(sum, l)

	// Hash it thrice.
	h3 := hashThree(h1, h2, l)

	// Take each resultant value from hashing and set that index
	// in Bloom Filter array to 1.
	b.data[h1] = 1
	b.data[h2] = 1
	b.data[h3] = 1
}

func hashOne(num, m int) int {
	return num % m
}

func hashTwo(num, m int) int {
	return (num * 17) % m
}

func hashThree(a, b, m int) int {
	return (a + (b * 71)) % m
}

func (b *myBloomFilter) maybeContains(item string) bool {
	// Take item argument and extract a number for hashing.
	r := []rune(item)

	sum := 0
	for _, i := range r {
		sum += int(i)
	}

	l := len(b.data)

	// Hash it once.
	h1 := hashOne(sum, l)

	// Hash it twice.
	h2 := hashTwo(sum, l)

	// Hash it thrice.
	h3 := hashThree(h1, h2, l)

	// If item's hashed positions are 1, might exist and return true.
	bit1 := b.data[h1]
	bit2 := b.data[h2]
	bit3 := b.data[h3]

	if bit1 == 1 && bit2 == 1 && bit3 == 1 {
		return true
	}

	// Else, if even one of the hased positions is 0, does not exist and return false.
	return false
}

func (b *myBloomFilter) memoryUsage() int {
	return binary.Size(b.data)
}

/*
	Below is the trivial implementation of a Bloom Filter included
	in the prework.
*/

type trivialBloomFilter struct {
	data []uint64
}

func newTrivialBloomFilter() *trivialBloomFilter {
	return &trivialBloomFilter{
		data: make([]uint64, 1000),
	}
}

func (b *trivialBloomFilter) add(item string) {
	// Do nothing
}

func (b *trivialBloomFilter) maybeContains(item string) bool {
	// Technically, any item "might" be in the set
	return true
}

func (b *trivialBloomFilter) memoryUsage() int {
	return binary.Size(b.data)
}
