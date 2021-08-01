package bitmap

const wordSize = 64

type uncompressedBitmap struct {
	data []uint64
}

func newUncompressedBitmap() *uncompressedBitmap {
	return &uncompressedBitmap{}
}

func (b *uncompressedBitmap) Get(x uint32) bool {
	if x >= uint32(len(b.data)) {
		b.growBitmapData(x)
	}

	v := b.data[x]
	if v == 0 {
		return false
	}

	return true
}

func (b *uncompressedBitmap) Set(x uint32) {
	// Check if argument is larger than current size of bitmap data slice.
	// If true, grow the bitmap data slice.
	if x >= uint32(len(b.data)) {
		b.growBitmapData(x)
	}

	b.data[x] = 1
}

func (b *uncompressedBitmap) Union(other *uncompressedBitmap) *uncompressedBitmap {
	// Find length of both bitmaps A and B.
	// Start loop with end condition set to length of smaller bitmap.
	// Fill in new data slice with value depending on value in both bitmaps at the same
	// index.
	//		- If one of the two bitmaps contains 1 at that index, set data[i] = 1.
	//		- Else, set data[i] = 0.
	// After the loop finishes, if bitmaps were not equal in length then append remaining values from larger bitmap to new data slice.
	//		- data = append(data, [subslice of remaining elements]...)

	var data []uint64
	return &uncompressedBitmap{
		data: data,
	}
}

func (b *uncompressedBitmap) Intersect(other *uncompressedBitmap) *uncompressedBitmap {
	// Similar operation as Union logic, but only set data[i] = 1 if both bitmaps have
	// value of 1 at the same index.
	// And after the loop, don't append the remaining elements as they only exist in
	// one of the two bitmaps. Thus, they are not in the intersect.
	var data []uint64
	return &uncompressedBitmap{
		data: data,
	}
}

// Grow the bitmap data slice by appending an empty slice with a length equal
// to the argument plus one.
// By growing the slice in this way, we will usually add many 0s as well that
// will speed up future Set() and Get() operations.
func (b *uncompressedBitmap) growBitmapData(x uint32) {
	padding := make([]uint64, x+1)
	b.data = append(b.data, padding...)
}
