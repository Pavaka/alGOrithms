package main

import (
	"math/rand"
	"testing"
	"time"
)

func makeMockList() *SkipList {
	list := NewSkipList()
	node2 := new(Node)
	node3 := new(Node)
	node5 := new(Node)
	node8 := new(Node)
	node2.key = 2
	node3.key = 3
	node5.key = 5
	node8.key = 8

	list.minNode.next = make([]*Node, 4)

	list.minNode.next[0] = node2
	list.minNode.next[1] = node2
	list.minNode.next[2] = node2
	list.minNode.next[3] = node2

	node2.next = make([]*Node, 4)
	node3.next = make([]*Node, 2)
	node5.next = make([]*Node, 3)
	node8.next = make([]*Node, 1)

	node2.next[0] = node3
	node2.next[1] = node3
	node2.next[2] = node5
	node2.next[3] = list.maxNode

	node3.next[0] = node5
	node3.next[1] = node5

	node5.next[0] = node8
	node5.next[1] = list.maxNode
	node5.next[2] = list.maxNode

	node8.next[0] = list.maxNode

	return list
}

func TestFindFirst(t *testing.T) {
	list := makeMockList()
	node := list.Find(3)
	if node == nil || node.key != 3 {
		t.Errorf("Incorrect node found")
	}
}

func TestFindInMiddle(t *testing.T) {
	list := makeMockList()
	node := list.Find(5)
	if node == nil || node.key != 5 {
		t.Errorf("Incorrect node found")
	}
}

func TestFindLast(t *testing.T) {
	list := makeMockList()
	node := list.Find(8)
	if node == nil || node.key != 8 {
		t.Errorf("Incorrect node found")
	}
}

func TestFindNonPresent(t *testing.T) {
	list := makeMockList()
	node := list.Find(4)
	if node != nil {
		t.Errorf("Incorrect node found")
	}
}

func TestFindEmptyList(t *testing.T) {
	list := NewSkipList()
	node := list.Find(3)
	if node != nil {
		t.Errorf("Incorrect node found")
	}
}

func TestFindGreatestSmallerThanInMiddle(t *testing.T) {
	list := makeMockList()
	node := list.FindGreatestSmallerThanOrEqual(4, NullOvershotHandler, true)
	if node == nil || node.key != 3 {
		t.Errorf("Incorrect node found")
	}
}

func TestFindGreatestSmallerThanOutsideAfter(t *testing.T) {
	list := makeMockList()
	node := list.FindGreatestSmallerThanOrEqual(100, NullOvershotHandler, true)
	if node.key != 8 {
		t.Errorf("Incorrect node found")
	}
}

func TestFindGreatestSmallerThanOutsideBelow(t *testing.T) {
	list := makeMockList()
	node := list.FindGreatestSmallerThanOrEqual(1, NullOvershotHandler, true)
	if node != list.minNode {
		t.Errorf("Incorrect node found")
	}
}

func TestInsertBasic(t *testing.T) {
	list := NewSkipList()
	list.Insert(4)
	list.Insert(1)
	list.Insert(15)
	list.Insert(10)

	if list.Find(1).key != 1 {
		t.Errorf("Incorrect node found")
	}

	if list.Find(4).key != 4 {
		t.Errorf("Incorrect node found")
	}

	if list.Find(10).key != 10 {
		t.Errorf("Incorrect node found")
	}

	if list.Find(15).key != 15 {
		t.Errorf("Incorrect node found")
	}

}

func TestDeleteBasic(t *testing.T) {
	list := NewSkipList()
	rand.Seed(time.Now().UTC().UnixNano())

	list.Insert(2)
	list.Insert(4)

	list.Delete(4)

	if list.Find(4) != nil {
		t.Errorf("Incorrect node found")
	}

	list.Insert(6)

	list.Insert(8)

	list.Delete(8)

	if list.Find(8) != nil {
		t.Errorf("Incorrect node found")
	}
}
func TestDeleteFromEmpty(t *testing.T) {
	list := NewSkipList()
	rand.Seed(time.Now().UTC().UnixNano())
	list.Delete(5)
}

func TestDeleteSingle(t *testing.T) {
	list := NewSkipList()
	rand.Seed(time.Now().UTC().UnixNano())
	list.Insert(5)
	list.Delete(5)
	if list.Find(5) != nil {
		t.Errorf("Incorrect node found")
	}
}
func TestDeleteFirst(t *testing.T) {
	list := NewSkipList()
	rand.Seed(time.Now().UTC().UnixNano())
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)

	list.Delete(1)

	if list.Find(1) != nil {
		t.Errorf("Incorrect node found")
	}
	// Check it is still accessible
	if list.Find(2).key != 2 {
		t.Errorf("Incorrect node found")
	}
}
func TestDeleteLast(t *testing.T) {
	list := NewSkipList()
	rand.Seed(time.Now().UTC().UnixNano())
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)

	list.Delete(3)

	if list.Find(3) != nil {
		t.Errorf("Incorrect node found")
	}
}
