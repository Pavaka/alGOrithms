package main

import (
	"bytes"
	"fmt"
	"math/rand"
)

type Node struct {
	next []*Node
	key  int
}

type SkipList struct {
	head *Node
}

func generateLevel() int {

	level := 0
	for rand.Int()%2 == 1 {
		level++
	}
	return level
}

type overshotHandler func(int, *Node)

func NullOvershotHandler(int, *Node) {}

func (list *SkipList) FindGreatestSmallerThan(key int, onOvershot overshotHandler) *Node {
	if list.head == nil || key < list.head.key {
		return nil
	}

	currNode := list.head
	currTowerPos := len(list.head.next) - 1
	for {
		if currNode.key == key || currTowerPos == -1 {
			return currNode
		}

		if currNode.next[currTowerPos] == nil || currNode.next[currTowerPos].key > key {
			onOvershot(currTowerPos, currNode)
			currTowerPos--
			continue
		}

		currNode = currNode.next[currTowerPos]
	}
}

func (list *SkipList) Insert(key int) {

	if list.head == nil {
		newNode := new(Node)
		newNode.key = key
		newNode.next = make([]*Node, 1)
		list.head = newNode
	}

	newKeyHeight := generateLevel()
	prevNodesByLevel := make([]*Node, newKeyHeight+1)
	onOvershot := func(currLevel int, node *Node) {
		if currLevel > newKeyHeight {
			return
		}

		prevNodesByLevel[currLevel] = node
	}

	maxHeight := 0
	if list.head.next != nil {
		maxHeight = len(list.head.next) - 1
	}

	if newKeyHeight > maxHeight {
		list.head.next = append(list.head.next, make([]*Node, newKeyHeight-maxHeight)...)
	}

	findGreatestSmallerThan := list.FindGreatestSmallerThan(key, onOvershot)
	if findGreatestSmallerThan.key == key {
		return // The key is already present
	}

	newNode := new(Node)
	newNode.key = key
	newNode.next = make([]*Node, newKeyHeight+1)

	for i, prevNode := range prevNodesByLevel {
		nextNode := prevNode.next[i]
		prevNode.next[i] = newNode
		newNode.next[i] = nextNode
	}

}

func (list *SkipList) Find(key int) *Node {
	res := list.FindGreatestSmallerThan(key, NullOvershotHandler)
	if res == nil || res.key != key {
		return nil
	}
	return res
}

func (list *SkipList) Print() {

	expectedPos := make(map[int]int)
	for curr := list.head; curr != nil; {
		expectedPos[curr.key] = len(expectedPos)
		curr = curr.next[0]
	}

	buff := bytes.NewBufferString("")
	for towerPos := len(list.head.next) - 1; towerPos >= 0; towerPos-- {

		currPrintPos := 0
		curr := list.head

		for curr != nil {
			desiredPos := expectedPos[curr.key]
			for ; currPrintPos < desiredPos; currPrintPos++ {
				fmt.Fprint(buff, "      ")
			}

			fmt.Fprint(buff, curr.key, " -> ")
			if curr.key < 10 {
				fmt.Fprint(buff, " ")
			}
			curr = curr.next[towerPos]
			currPrintPos++
		}

		fmt.Println(buff.String())
		buff.Reset()
	}
	fmt.Println()
}

func main() {

}
