package bloom

import (
	"encoding/binary"
	_ "fmt"

	m3 "github.com/spaolacci/murmur3"
	bv "juliuszerwick/storage-retrieval/bloom-filters/bitvector"
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

type MyBloomFilter struct {
	//data []uint8
	Data *bv.BitVector
}

func NewMyBloomFilter() *MyBloomFilter {
	data := make([]byte, 100000)
	//data := make([]byte, 20)

	return &MyBloomFilter{
		Data: bv.NewBitVector(data, 800000),
		//	Data: NewBitVector(data, 160),
	}
}

// Add item to BF.
func (b *MyBloomFilter) Add(item string) {
	//	fmt.Printf("Added item: %s\n\n", item)
	// Take item argument and extract a number for hashing.
	//r := []rune(item)

	//sum := 0
	//for _, i := range r {
	//	sum += int(i)
	//}

	l := b.Data.Length

	// Hash it once.
	h1 := hashOne(item, l)

	// Hash it twice.
	h2 := hashTwo(item, l)

	// Hash it thrice.
	h3 := hashThree(h1, h2, l)

	//fmt.Printf("h1 = %v, h2 = %v, h3 = %v\n\n", h1, h2, h3)

	// Take each resultant value from hashing and set that index
	// in Bloom Filter array to 1.
	//b.data[h1] = 1
	//b.data[h2] = 1
	//b.data[h3] = 1
	b.Data.Set(1, h1)
	b.Data.Set(1, h2)
	b.Data.Set(1, h3)

	//fmt.Printf("bitVector: %v\n\n", b.data)
}

//func hashOne(num, m int) int {
func hashOne(w string, m int) int {
	//return num % m
	v := int(m3.Sum64([]byte(w))) % m
	if v < 0 {
		v = -1 * v
	}

	return v
}

func hashTwo(w string, m int) int {
	//return (num * 17) % m
	v := (int(m3.Sum64([]byte(w))) * 17) % m
	if v < 0 {
		v = -1 * v
	}

	return v
}

func hashThree(a, b, m int) int {
	return (a + (b * 71)) % m
}

func (b *MyBloomFilter) MaybeContains(item string) bool {
	// Take item argument and extract a number for hashing.
	//r := []rune(item)

	//sum := 0
	//for _, i := range r {
	//	sum += int(i)
	//}

	l := b.Data.Length

	// Hash it once.
	h1 := hashOne(item, l)

	// Hash it twice.
	h2 := hashTwo(item, l)

	// Hash it thrice.
	h3 := hashThree(h1, h2, l)

	// If item's hashed positions are 1, might exist and return true.
	//bit1 := b.data[h1]
	//bit2 := b.data[h2]
	//bit3 := b.data[h3]
	// Check each retrieved bit if 0 before checking next bit? Small optimization?
	bit1 := b.Data.Element(h1)
	bit2 := b.Data.Element(h2)
	bit3 := b.Data.Element(h3)

	if bit1 == 1 && bit2 == 1 && bit3 == 1 {
		return true
	}

	// Else, if even one of the hased positions is 0, does not exist and return false.
	return false
}

func (b *MyBloomFilter) MemoryUsage() int {
	return binary.Size(b.Data.Data)
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
