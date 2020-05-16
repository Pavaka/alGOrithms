package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
)

type Node struct {
	next []*Node
	key  int
}

type SkipList struct {
	minNode *Node
	maxNode *Node
}

func NewSkipList() *SkipList {
	list := new(SkipList)

	list.maxNode = new(Node)
	list.maxNode.key = math.MaxInt32

	list.minNode = new(Node)
	list.minNode.key = math.MinInt32
	list.minNode.next = make([]*Node, 1)
	list.minNode.next[0] = list.maxNode

	return list
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
	currNode := list.minNode
	currTowerPos := len(list.minNode.next) - 1
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
	newKeyHeight := generateLevel()
	// if the key can be inserted this will hold all nodes
	// before the new insertion node for each level
	// in which it will exist
	prevNodesByLevel := make([]*Node, newKeyHeight+1)
	onOvershot := func(currLevel int, node *Node) {
		if currLevel > newKeyHeight {
			return
		}

		prevNodesByLevel[currLevel] = node
	}

	maxHeight := len(list.minNode.next) - 1

	if newKeyHeight > maxHeight {
		list.minNode.next = append(list.minNode.next, make([]*Node, newKeyHeight-maxHeight)...)
		for i := newKeyHeight; i > maxHeight; i-- {
			list.minNode.next[i] = list.maxNode
		}
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
	if res == list.minNode || res.key != key {
		return nil
	}
	return res
}

func (list *SkipList) Print() {
	expectedPos := make(map[int]int)
	for curr := list.minNode; curr != list.maxNode; curr = curr.next[0] {
		expectedPos[curr.key] = len(expectedPos)
	}

	buff := bytes.NewBufferString("")
	for towerPos := len(list.minNode.next) - 1; towerPos >= 0; towerPos-- {

		currPrintPos := 0
		curr := list.minNode

		for curr != list.maxNode {

			desiredPos := expectedPos[curr.key]
			for ; currPrintPos < desiredPos; currPrintPos++ {
				fmt.Fprint(buff, "      ")
			}

			if curr != list.minNode {
				fmt.Fprint(buff, curr.key, " -> ")

				if curr.key < 10 {
					fmt.Fprint(buff, " ")
				}
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
