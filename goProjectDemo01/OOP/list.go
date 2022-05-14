package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

func Createlist(nums ...int) *Node {
	head := &Node{nums[0], nil}
	prev := head
	for i := 1; i < len(nums); i++ {
		tmp := Node{nums[i], nil}
		prev.next = &tmp
		prev = &tmp
	}
	return head
}

func PrintList(node *Node) {
	for node != nil {
		fmt.Printf("%d ", node.val)
		node = node.next
	}
}

func ReverserList(node *Node) *Node {

}

func main() {
	head := Createlist(7, 8, 7, 1, 4, 8, 9, 2, 8)
	nhead := ReverserList(head)
	PrintList(nhead)
}
