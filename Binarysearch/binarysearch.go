package main

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

type BinarySearchTree struct {
	Root *TreeNode
}

func (tree BinarySearchTree) Insert(v int) {
	tree.Root.Insert(v)
}

func (node *TreeNode) Insert(v int) {
	if v < node.Value {
		if node.Left != nil {
			node.Left.Insert(v)
		} else {
			node.Left = &TreeNode{v,nil,nil}
		}
	} else {
		if node.Right != nil {
			node.Right.Insert(v)
		} else {
			node.Right = &TreeNode{v,nil,nil}
		}
	}
}

func (tree 	BinarySearchTree) InOrder() []int {
	var res []int
	tree.Root.InOrder(&res)
	return res
}

func (node *TreeNode) InOrder(result *[]int) {
	if node.Left != nil {
		node.Left.InOrder(result)
	}
	*result = append(*result, node)
	if node.Right != nil {
		node.Right.InOrder(result)
	}
} 

func (tree BinarySearchTree) FindMin() int {
	node = tree.Root
	for {
		if node.Left != nil {
			node = node.Left
		} else {
			return node.Value
		}
	}
}

func (tree BinarySearchTree) FindMax() {
	node = tree.Root
	for {
		if node.Right != nil {
			node = node.Right
		} else {
			return node.Value
		}
	}
}

func (tree BinarySearch) Find(v int) bool {
	node = tree.Node
	for {
		if node == nil {
			return false
		} else if node.Value == v {
			return true
		} else if node.Value > v {
			node = node.Left
		} else {
			node = node.Right
		}
	}
}