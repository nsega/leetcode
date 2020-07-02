package main

import (
	"fmt"
)

// TreeNode is a struct of a node of a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BinaryTree is a struct of a binary tree
type BinaryTree struct {
	root *TreeNode
}

func (t *BinaryTree) insert(val int) *BinaryTree {
	if t.root == nil {
		t.root = &TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
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
			n.Left = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
		} else {
			n.Left.insert(val)
		}
	} else {
		if n.Right == nil {
			n.Right = &TreeNode{
				Val:   val,
				Left:  nil,
				Right: nil,
			}
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

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil {
		return t2
	}

	if t2 == nil {
		return t1
	}

	return &TreeNode{
		Val:   t1.Val + t2.Val,
		Left:  mergeTrees(t1.Left, t2.Left),
		Right: mergeTrees(t1.Right, t2.Right),
	}
}

func main() {

	fmt.Println("------TestCase1-------")
	t1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	print(t1, 0, 'M')

	t2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	print(t2, 0, 'M')
	print(mergeTrees(t1, t2), 0, 'M')

	// other test case by using a binary tree
	fmt.Println("------TestCase2-------")
	tree1 := &BinaryTree{}
	tree1.insert(1).insert(3).insert(2).insert(5)
	print(tree1.root, 0, 'M')
	tree2 := &BinaryTree{}
	tree2.insert(2).insert(1).insert(3).insert(4).insert(7)
	print(tree2.root, 0, 'M')
	print(mergeTrees(tree1.root, tree2.root), 0, 'M')
}
