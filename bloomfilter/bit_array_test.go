package bloomfilter

import "testing"

func TestSetBit(t *testing.T) {
	arr := NewBitArray(3)
	arr.SetBit(10)
	if arr.GetBit(10) != true {
		t.Error()
	}
}

func TestGetBitFalse(t *testing.T) {
	arr := NewBitArray(3)
	arr.SetBit(5)
	if arr.GetBit(2) != false {
		t.Error()
	}
}

func TestUnsetBit(t *testing.T) {
	arr := NewBitArray(3)
	arr.SetBit(15)
	arr.UnsetBit(15)
	if arr.GetBit(15) != false {
		t.Error()
	}
}
