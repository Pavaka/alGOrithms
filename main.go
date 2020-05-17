package main

import (
	"math/rand"
	"time"

	"github.com/Pavaka/alGOrithms/skiplist"
)

func main() {
	// Sample test list
	list := skiplist.NewSkipList()
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 20; i++ {
		list.Insert(rand.Int() % 80)
	}
	list.Print()
}
