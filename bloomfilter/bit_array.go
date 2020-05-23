package bloomfilter

import (
	"bytes"
	"fmt"
)

// BitArray is a memory buffer
// with big granularity interface
type BitArray struct {
	mem []byte
}

// NewBitArray create new bit array
func NewBitArray(numBytes int) *BitArray {

	res := new(BitArray)
	res.mem = make([]byte, numBytes)
	return res
}

// SetBit inside bit array
func (arr *BitArray) SetBit(pos int) {
	bytePos := pos / 8
	inBytePos := pos % 8
	newByte := byte((1 << 7) >> inBytePos)
	arr.mem[bytePos] |= newByte
}

// UnsetBit inside bit array
func (arr *BitArray) UnsetBit(pos int) {
	bytePos := pos / 8
	inBytePos := pos % 8
	newByte := 0xFF &^ (byte((1 << 7) >> inBytePos))
	arr.mem[bytePos] &= newByte
}

func (arr *BitArray) GetBit(pos int) bool {
	bytePos := pos / 8
	inBytePos := pos % 8
	newByte := byte((1 << 7) >> inBytePos)
	res := arr.mem[bytePos] & newByte
	return res != 0
}

// Print the content of a bit array
func (arr *BitArray) Print() {
	buff := bytes.NewBufferString("")
	for i := range arr.mem {
		fmt.Fprintf(buff, "%08b ", arr.mem[i])
	}
	fmt.Println(buff)
}
