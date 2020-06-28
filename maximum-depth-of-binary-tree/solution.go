package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	root *TreeNode
}

func (t *BinaryTree) insert(val int) *BinaryTree {
	if t.root == nil {
		t.root = &TreeNode{Val: val, Left: nil, Right: nil}
	} else {
		t.root.insert(val)
	}
	return t
}

func (n *TreeNode) insert(val int) {
	if n == nil {
		return
	} else if val <= n.Val {
		if n.Left == nil {
			n.Left = &TreeNode{Val: val, Left: nil, Right: nil}
		} else {
			n.Left.insert(val)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{Val: val, Left: nil, Right: nil}
		} else {
			n.Right.insert(val)
		}
	}
}

func print(node *TreeNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Printf(" ")
	}
	fmt.Printf("%c:%v\n", ch, node.Val)
	print(node.Left, ns+2, 'L')
	print(node.Right, ns+2, 'R')
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func main() {
	tree := &BinaryTree{}
	tree.insert(9).insert(3).insert(20).insert(15).insert(7)
	fmt.Println(maxDepth(tree.root))
	print(tree.root, 0, 'M')
}
