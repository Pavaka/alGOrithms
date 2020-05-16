package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Node of the skip list
type Node struct {
	next []*Node
	key  int
}

// SkipList is list alike data structure
// that offers faster traversal for the cost
// of more memory
type SkipList struct {
	minNode *Node
	maxNode *Node
}

// NewSkipList constructs a SkipList
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

// Insert a key in the list
// Does nothing if the key is present
func (list *SkipList) Insert(key int) {
	newKeyLevel := generateLevel()
	maxLevel := len(list.minNode.next) - 1

	if newKeyLevel > maxLevel {
		list.minNode.next = append(list.minNode.next, make([]*Node, newKeyLevel-maxLevel)...)
		for i := newKeyLevel; i > maxLevel; i-- {
			list.minNode.next[i] = list.maxNode
		}
	}

	// if the key can be inserted this will hold all nodes
	// before the new insertion node for each level
	// in which it will exist
	prevNodesByLevel := make([]*Node, newKeyLevel+1)
	onLevelSearchEnd := makeSavePrevNodeHandler(newKeyLevel, prevNodesByLevel)

	foundNode := list.findGreatestSmallerThanOrEqual(key, onLevelSearchEnd, true)
	if foundNode.key == key {
		return // The key is already present
	}

	newNode := new(Node)
	newNode.key = key
	newNode.next = make([]*Node, newKeyLevel+1)

	for i, prevNode := range prevNodesByLevel {
		nextNode := prevNode.next[i]
		prevNode.next[i] = newNode
		newNode.next[i] = nextNode
	}

}

// Delete a key from the list if present
func (list *SkipList) Delete(key int) {
	maxHeight := len(list.minNode.next) - 1
	prevNodesByLevel := make([]*Node, len(list.minNode.next))
	onOvershot := makeSavePrevNodeHandler(maxHeight, prevNodesByLevel)

	found := list.findGreatestSmallerThanOrEqual(key, onOvershot, false)

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

// Find if a key is present in the list
func (list *SkipList) Find(key int) *Node {
	res := list.findGreatestSmallerThanOrEqual(key, nullHandler, true)
	if res == list.minNode || res.key != key {
		return nil
	}
	return res
}

// Print the content of the skip list
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

func generateLevel() int {

	level := 0
	for rand.Int()%2 == 1 {
		level++
	}
	return level
}

// we finish a search on a level when either we find
// the node, or we overshoot it
type onLevelSearchEndHandler func(int, *Node)

func nullHandler(int, *Node) {}

// when the search for a level ends we
// save the prev node on this level
func makeSavePrevNodeHandler(maxObservableLevel int, prevNodesByLevel []*Node) onLevelSearchEndHandler {
	handler := func(currLevel int, node *Node) {
		if currLevel > maxObservableLevel {
			return
		}

		prevNodesByLevel[currLevel] = node
	}
	return handler
}

func (list *SkipList) findGreatestSmallerThanOrEqual(key int, onLevelSearchEnd onLevelSearchEndHandler, stopAtFirstHit bool) *Node {
	currNode := list.minNode
	currLevel := len(list.minNode.next) - 1

	for currLevel != -1 {

		if stopAtFirstHit && currNode.next[currLevel].key == key {
			return currNode.next[currLevel]
		}

		if currNode.next[currLevel].key >= key {
			onLevelSearchEnd(currLevel, currNode)
			currLevel--
			continue
		}

		currNode = currNode.next[currLevel]
	}

	if currNode.next[0].key == key {
		return currNode.next[0]
	}
	return currNode
}

func main() {
	// Sample test list
	list := NewSkipList()
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 20; i++ {
		list.Insert(rand.Int() % 80)
	}
	list.Print()
}
