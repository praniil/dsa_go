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
	prevNode *Node
}

func (list *LinkedList) insertItems(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head //we dont use the pointer head during the operation but we use a temp pointer i.e current because head should only point the first node
		for current.nextNode != nil {
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
	count := 0
	current := list.head
	for current != nil {
		count++
		action(current)
		current = current.nextNode
	}
	fmt.Println("no of nodes: ", count)
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

// func (list *LinkedList) insertTop (data int) {
// 	newNode := &Node{data: data}

// 	if list.head == nil {
// 		list.head = newNode
// 	} else {
// 		currentNode := list.head

// 	}
// }

func (list *LinkedList) insertBeginning(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
	} else {
		newNode.nextNode = list.head
		list.head = newNode
	}

}

func (list *LinkedList) insertAtSpecificPosition(data int, position int) {
	newNode := &Node{data: data}
	current := list.head
	count := 0
	if position == 0 {
		newNode.nextNode = current
		list.head = newNode
		return
	}
	for current.nextNode != nil && count < position-1 {
		current = current.nextNode
		count++
	}
	tempNode := current.nextNode
	current.nextNode = newNode
	newNode.nextNode = tempNode
}

func (list *LinkedList) deleteNode(data int) {
	current := list.head
	for current.nextNode.data != data {
		current = current.nextNode
		if current.nextNode == nil {
			fmt.Printf("%d not found!", data)
		}
	}
	current.nextNode = current.nextNode.nextNode
	current = current.nextNode
}

func LLdisplay() {
	var linkLis LinkedList
	linkLis.insertItems(15)
	linkLis.insertItems(25)
	linkLis.insertItems(35)
	linkLis.insertItems(45)
	linkLis.insertItems(55)
	linkLis.insertItems(65)
	// linkLis.insertBeginning(15)
	linkLis.insertAtSpecificPosition(200, 10)

	// linkLis.deleteNode(2)
	linkLis.display()
	// linkLis.traversing(func(node *Node) {
	// 	fmt.Printf("%d -> ", node.data)
	// })
	result, value := linkLis.search(20)
	if result == true {
		fmt.Printf("\n%v. The value is: %d", result, value)
	} else {
		fmt.Println("\nNOT FOUND")
	}
}
