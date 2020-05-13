package main

import (
	"bytes"
	"fmt"
)

type Node struct {
	next  []*Node
	value int
}

type SkipList struct {
	head  *Node
	sizes []int
}

func (list *SkipList) Find(value int) *Node {
	if list.head == nil {
		return nil
	}

	currNode := list.head
	currTowerPos := len(list.sizes) - 1
	for {
		if currNode.value == value {
			return currNode
		}

		if currTowerPos == -1 {
			return nil
		}

		if currNode.next[currTowerPos] == nil || currNode.next[currTowerPos].value > value {
			currTowerPos--
			continue
		}
		currNode = currNode.next[currTowerPos]
	}
}

func (list *SkipList) Print() {

	expectedPos := make(map[int]int)
	for curr := list.head; curr != nil; {
		expectedPos[curr.value] = len(expectedPos)
		curr = curr.next[0]
	}

	buff := bytes.NewBufferString("")
	for towerPos := len(list.sizes) - 1; towerPos >= 0; towerPos-- {

		currPrintPos := 0
		curr := list.head

		for curr != nil {
			desiredPos := expectedPos[curr.value]
			for ; currPrintPos < desiredPos; currPrintPos++ {
				fmt.Fprint(buff, "     ")
			}

			fmt.Fprint(buff, curr.value, " -> ")
			curr = curr.next[towerPos]
			currPrintPos++
		}

		fmt.Println(buff.String())
		buff.Reset()
	}
}

func main() {
}
