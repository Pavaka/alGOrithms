package bloomfilter

import "testing"

func TestEmptyContainsBasic(t *testing.T) {
	filter := NewBloomFilter()
	if filter.Contains("pavaka") != false {
		t.Error()
	}
}

func TestContainsBasic(t *testing.T) {
	filter := NewBloomFilter()
	filter.Add("pavaka")
	filter.Add("dido")
	filter.Add("lybaka")

	testMatrix := []struct {
		value  string
		result bool
	}{
		{"pavaka", true},
		{"dido", true},
		{"lybaka", true},
		{"lebron", false},
		{"lillard", false},
	}

	for _, test := range testMatrix {
		if filter.Contains(test.value) != test.result {
			t.Error()
		}
	}
}
