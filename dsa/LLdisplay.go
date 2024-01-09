package dsa

import (
	"fmt"
)

type LinkedList struct {
	head *Node
}

type Node struct {
	data     int
	nextNode *Node
}

func (list *LinkedList) insertItems(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head //we dont use the pointer head during the operation but we use a temp pointer i.e current because head should only point the first node
		for newNode.nextNode != nil {
			current = current.nextNode //p = p-> next
		}
		current.nextNode = newNode //last item ko paxadi newnode is made
	}
}

func (list *LinkedList) display() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.nextNode //current vanne pointer i.e the head increased each time
	}
}

func (list *LinkedList) traversing(action func(node *Node)) {
	current := list.head
	for current != nil {
		action(current)
		current = current.nextNode
	}
}

func (list *LinkedList) search(data int) (bool, int) {
	current := list.head
	for current.nextNode != nil {
		if current.data == data {
			return true, current.data
		}
		current = current.nextNode
	}
	return false, 0
}

func LLdisplay() {
	var linkLis LinkedList
	linkLis.insertItems(25)
	linkLis.insertItems(35)
	linkLis.display()
	linkLis.traversing(func(node *Node) {
		fmt.Printf("%d -> ", node.data)
	})
	result, value := linkLis.search(20)
	if result == true {
		fmt.Printf("%v. The value is: %d", result, value)
	} else {
		fmt.Println("\nNOT FOUND")
	}
}
