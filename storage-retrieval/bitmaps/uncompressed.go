package bitmap

const wordSize = 64

type uncompressedBitmap struct {
	data []uint64
}

func newUncompressedBitmap() *uncompressedBitmap {
	return &uncompressedBitmap{}
}

func (b *uncompressedBitmap) Get(x uint32) bool {
	v := b.data[x]

	if v == 0 {
		return false
	}

	return true
}

func (b *uncompressedBitmap) Set(x uint32) {
	if x >= uint32(len(b.data)) {
		b.growBitmapData(x)
	}

	b.data[x] = 1
}

func (b *uncompressedBitmap) Union(other *uncompressedBitmap) *uncompressedBitmap {
	var data []uint64
	return &uncompressedBitmap{
		data: data,
	}
}

func (b *uncompressedBitmap) Intersect(other *uncompressedBitmap) *uncompressedBitmap {
	var data []uint64
	return &uncompressedBitmap{
		data: data,
	}
}

func (b *uncompressedBitmap) growBitmapData(x uint32) {
	padding := make([]uint64, x+1)
	b.data = append(b.data, padding...)
}
