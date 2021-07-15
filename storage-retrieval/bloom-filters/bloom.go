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

func (b *myBloomFilter) add(item string) {
	// Add item to BF.
}

func (b *myBloomFilter) maybeContains(item string) bool {
	// If item's hashed positions are 1, might exist and return true.

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
