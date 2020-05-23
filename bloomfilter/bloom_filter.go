package bloomfilter

import (
	"crypto/md5"
)

// BloomFilter ...
type BloomFilter struct {
	array BitArray
}

const bitArraySize = 256
const hashFunctionsCount = 4

// NewBloomFilter creates a new bloom filter
func NewBloomFilter() *BloomFilter {
	filter := new(BloomFilter)
	filter.array = *NewBitArray(bitArraySize / 8)
	return filter
}

// Add a value to bloom filter
func (filter *BloomFilter) Add(value string) {
	hashPos := genHashPositions(value)
	for _, pos := range hashPos {
		filter.array.SetBit(pos)
	}
}

// Contains returns true if the value *might* be in the filter
// or false if the value is *defenitely not* in the filter
func (filter *BloomFilter) Contains(value string) bool {
	hashPos := genHashPositions(value)
	for _, pos := range hashPos {
		if filter.array.GetBit(pos) == false {
			return false
		}
	}

	return true
}

// Print the underlying storage of the filter
func (filter *BloomFilter) Print() {
	filter.array.Print()
}

func genHashPositions(value string) []int {
	res := make([]int, hashFunctionsCount)
	// TODO: Use faster hashing like HashMix or MurmurHash
	hashVal := md5.Sum([]byte(value))

	for i := 0; i < hashFunctionsCount; i++ {
		var val int = (int(hashVal[i*4+0]) << 24) | int((hashVal[i*4+1]))<<16 | int((hashVal[i*4+2]))<<8 | int(hashVal[i*4+3])
		res[i] = val % bitArraySize
	}

	return res
}
