package main

import (
	"fmt"
	"unsafe"
)

// When the exercise/Elliot says to "reverse engineer" the code/implementation,
// it really means to build it yourself.

const (
	// Maximum number of key/elem pairs a bucket can hold.
	bucketCntBits = 3
	bucketCnt     = 1 << bucketCntBits

	// data offset should be the size of the bmap struct, but needs to be
	// aligned correctly. For amd64p32 this means 64-bit alignment
	// even though pointers are 32 bit.
	dataOffset = unsafe.Offsetof(struct {
		b bmap
		v int64
	}{}.v)
)

// A header for a Go map.
type hmap struct {
	count     int // # live cells == size of map.  Must be first (used by len())
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}

// mapextra holds fields that are not present on all maps.
type mapextra struct {
	// If both key and elem do not contain pointers and are inline, then we mark bucket
	// type as containing no pointers. This avoids scanning such maps.
	// However, bmap.overflow is a pointer. In order to keep overflow buckets
	// alive, we store pointers to all overflow buckets in hmap.extra.overflow and hmap.extra.oldoverflow.
	// overflow and oldoverflow are only used if key and elem do not contain pointers.
	// overflow contains overflow buckets for hmap.buckets.
	// oldoverflow contains overflow buckets for hmap.oldbuckets.
	// The indirection allows to store a pointer to the slice in hiter.
	overflow    *[]*bmap
	oldoverflow *[]*bmap

	// nextOverflow holds a pointer to a free overflow bucket.
	nextOverflow *bmap
}

// A bucket for a Go map.
type bmap struct {
	// tophash generally contains the top byte of the hash value
	// for each key in this bucket. If tophash[0] < minTopHash,
	// tophash[0] is a bucket evacuation state instead.
	tophash [bucketCnt]uint8
	// Followed by bucketCnt keys and then bucketCnt elems.
	// NOTE: packing all the keys together and then all the elems together makes the
	// code a bit more complicated than alternating key/elem/key/elem/... but it allows
	// us to eliminate padding which would be needed for, e.g., map[int64]int8.
	// Followed by an overflow pointer.
}

func main() {
	m := make(map[int]int)

	for i := 0; i < 6; i++ {
		m[i] = i
	}

	//m := map[int]int{1: 1, 2: 2, 3: 3}

	//fmt.Printf("map: %+v\n", m)

	ksum, vsum := sumsOfMap(m)
	fmt.Printf("sum of keys: %d\nsum of values: %d\n", ksum, vsum)
}

func sumsOfMap(m map[int]int) (ksum int, vsum int) {
	/* This type allows us to access the first two fields of the Value struct
	   that is returned when a map is created.

	   type Value struct {
	   	 typ
	     ptr   unsafe.Pointer
	     flag
	   }

		 The ptr field will be to an hmap struct when make()
		 is called to create a map instance.
	*/

	type mapinterface struct {
		maptype unsafe.Pointer // matches Value.typ
		data    unsafe.Pointer // matches Value.ptr
	}

	//fmt.Printf("mapinterface: %+v\n", (*mapinterface)(unsafe.Pointer(&m)))
	//fmt.Printf("mapinterface.data: %+v\n", (*mapinterface)(unsafe.Pointer(&m)).data)

	hmapptr := (*hmap)((*mapinterface)(unsafe.Pointer(&m)).data)
	//fmt.Printf("hmapptr: %+v\n", hmapptr)
	buckets := hmapptr.buckets
	numBuckets := uintptr(1 << hmapptr.B)

	thoffset := unsafe.Sizeof(uint8(0)) // Offset for tophash elements
	kvoffset := unsafe.Sizeof(int(0))   // Offset for key and value elements
	// Full size of a bucket: tophash, keys + values, and overflow pointer
	bucketoffset := dataOffset + (2 * bucketCnt * kvoffset) + unsafe.Sizeof(hmapptr)

	return 0, 0
}
