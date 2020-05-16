package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"time"
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

type shotOrOvershotHandler func(int, *Node)

func NullOvershotHandler(int, *Node) {

}

func MakeSavePrevNodeOnShotOrOvershotHandler(maxObservableHeight int, prevNodesByLevel []*Node) shotOrOvershotHandler {
	onOvershot := func(currLevel int, node *Node) {
		if currLevel > maxObservableHeight {
			return
		}

		prevNodesByLevel[currLevel] = node
	}

	return onOvershot
}

func (list *SkipList) FindGreatestSmallerThanOrEqual(key int, onShotOrOvershot shotOrOvershotHandler, stopAtFirstHit bool) *Node {
	currNode := list.minNode
	currTowerPos := len(list.minNode.next) - 1

	for currTowerPos != -1 {

		if stopAtFirstHit && currNode.next[currTowerPos].key == key {
			return currNode.next[currTowerPos]
		}

		if currNode.next[currTowerPos].key >= key {
			onShotOrOvershot(currTowerPos, currNode)
			currTowerPos--
			continue
		}

		currNode = currNode.next[currTowerPos]

	}

	if currNode.next[0].key == key {
		return currNode.next[0]
	}
	return currNode
}

func (list *SkipList) Insert(key int) {
	newKeyHeight := generateLevel()
	// if the key can be inserted this will hold all nodes
	// before the new insertion node for each level
	// in which it will exist
	prevNodesByLevel := make([]*Node, newKeyHeight+1)
	onOvershot := MakeSavePrevNodeOnShotOrOvershotHandler(newKeyHeight, prevNodesByLevel)

	maxHeight := len(list.minNode.next) - 1

	if newKeyHeight > maxHeight {
		list.minNode.next = append(list.minNode.next, make([]*Node, newKeyHeight-maxHeight)...)
		for i := newKeyHeight; i > maxHeight; i-- {
			list.minNode.next[i] = list.maxNode
		}
	}

	findGreatestSmallerThan := list.FindGreatestSmallerThanOrEqual(key, onOvershot, true)
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

func (list *SkipList) Delete(key int) {
	maxHeight := len(list.minNode.next) - 1
	prevNodesByLevel := make([]*Node, len(list.minNode.next))
	onOvershot := MakeSavePrevNodeOnShotOrOvershotHandler(maxHeight, prevNodesByLevel)

	found := list.FindGreatestSmallerThanOrEqual(key, onOvershot, false)

	if found.key != key {
		return
	}

	foundHeight := len(found.next) - 1
	for i := range prevNodesByLevel {

		if i > foundHeight {
			break
		}
		prevNodesByLevel[i].next[i] = found.next[i]
	}

}

func (list *SkipList) Find(key int) *Node {
	res := list.FindGreatestSmallerThanOrEqual(key, NullOvershotHandler, true)
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
	list := NewSkipList()
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 12; i++ {
		list.Insert(rand.Int() % 80)
	}
	list.Print()
}
