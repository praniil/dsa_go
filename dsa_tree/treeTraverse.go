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
}
