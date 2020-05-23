package main

import (
	"fmt"

	bf "github.com/Pavaka/alGOrithms/bloomfilter"
)

func main() {
	// Sample test list
	// list := skiplist.NewSkipList()
	// rand.Seed(time.Now().UTC().UnixNano())

	// for i := 0; i < 20; i++ {
	// 	list.Insert(rand.Int() % 80)
	// }
	// list.Print()

	// Sample bit array
	// arr := bf.NewBitArray(3)
	// arr.SetBit(0)
	// arr.SetBit(3)
	// arr.SetBit(17)
	// arr.SetBit(15)
	// arr.Print()
	// arr.UnsetBit(17)
	// arr.Print()

	// Sample bloom filter
	filter := bf.NewBloomFilter()
	filter.Add("pavaka")
	filter.Add("dido")
	filter.Print()
	fmt.Println(filter.Contains("pavaka"))
	fmt.Println(filter.Contains("dido"))
	fmt.Println(filter.Contains("mro"))
}
