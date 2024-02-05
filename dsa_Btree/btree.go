package dsa_Btree

import(
	"fmt"
)

const maxKeys = 3

type BTreeNode struct{
	keys []int
	children []*BTreeNode
	leaf bool
}

func NewBTreeNode(leaf bool) *BTreeNode {
	return &BTreeNode{
		keys = make([]int, 0),
		children = make([]*BTreeNode, 0),
		leaf = leaf
	}
}

type Btree struct {
	root *BTreeNode
}

func NewBTree () *Btree {
	return &Btree{}
}

func (t *Btree) Insert(key int) {
	if t.root == nil {
		t.root = NewBTreeNode(true)
		t.root.keys = append(t.root.keys, key)
	} else {
		
	}
}

func Btree () {
	fmt.Println("in b tree")
}