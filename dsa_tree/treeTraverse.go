package dsatree

import "fmt"

type Node struct {
	value       int
	Left, Right *Node
}

func preOrderTraversal(node *Node) {
	if node != nil {
		fmt.Println(node.value)
		preOrderTraversal(node.Left)
		preOrderTraversal(node.Right)
	}
}

func inOrderTraversal(node *Node) {
	if node != nil {
		inOrderTraversal(node.Left)
		fmt.Println(node.value)
		inOrderTraversal(node.Right)
	}
}

func postOrderTraversal(node *Node) {
	if node != nil {
		postOrderTraversal(node.Left)
		postOrderTraversal(node.Right)
		fmt.Println(node.value)
	}
}

func searchNode(node *Node, data int) {
	for node != nil {
		if node.value == data {
			fmt.Printf("Data found: %d\n", node.value)
			return
		} else if data > node.value {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	fmt.Println("Not found")
}

func deleteNode(node *Node, key int) *Node{
	if node == nil {
		return nil
	}
	if key < node.value {
		node.Left = deleteNode(node.Left, key)
	} else if key > node.value {
		node.Right = deleteNode(node.Right, key)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		minNode := findMin(node.Right)
		node.value = minNode.value
		node.Right = deleteNode(node.Right, minNode.value)
	}
	return node
}

func findMin (node *Node) *Node {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func TreeTraverse() {
	root := &Node{
		value: 5,
		Left: &Node{
			value: 3,
			Left:  nil,
			Right: nil,
		},
		Right: &Node{
			value: 8,
			Left: &Node{
				value: 6,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	
	fmt.Println("PreOrder Traverse is:")
	preOrderTraversal(root)
	fmt.Println("InOrder Traverse is:")
	inOrderTraversal(root)
	fmt.Println("PostOrder Traverse is:")
	postOrderTraversal(root)
	searchNode(root, 1)
	searchNode(root, 8)
	deleteNode(root, 8)
	searchNode(root, 8)
}
