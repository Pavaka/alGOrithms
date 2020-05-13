package main

import "testing"

func makeMockList() *SkipList {
	list := new(SkipList)
	node2 := new(Node)
	node3 := new(Node)
	node5 := new(Node)
	node8 := new(Node)
	node2.value = 2
	node3.value = 3
	node5.value = 5
	node8.value = 8

	list.head = node2
	node2.next = make([]*Node, 4)
	node3.next = make([]*Node, 2)
	node5.next = make([]*Node, 3)
	node8.next = make([]*Node, 1)

	node2.next[0] = node3
	node2.next[1] = node3
	node2.next[2] = node5

	node3.next[0] = node5
	node3.next[1] = node5

	node5.next[0] = node8

	list.sizes = make([]int, 4)
	list.sizes[0] = 4
	list.sizes[0] = 3
	list.sizes[0] = 2
	list.sizes[0] = 1

	return list
}

func TestFind(t *testing.T) {

	list := makeMockList()
	node := list.Find(3)
	if node == nil || node.value != 3 {
		t.Errorf("Incorrect node found")
	}

	node = list.Find(8)
	if node == nil || node.value != 8 {
		t.Errorf("Incorrect node found")
	}

	node = list.Find(4)
	if node != nil {
		t.Errorf("Incorrect node found")
	}
}

func TestRunPrint(t *testing.T) {
	list := makeMockList()
	list.Print()
}
