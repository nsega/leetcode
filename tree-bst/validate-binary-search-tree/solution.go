package main

import (
	"fmt"
	"math"
)

// TreeNode is a struct for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {

	var stack []*TreeNode
	max := math.MinInt64
	node := root
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		if len(stack) != 0 {
			node = stack[len(stack)-1]
			if node.Val > max {
				max = node.Val
			} else {
				return false
			}
			stack = stack[:len(stack)-1]
			node = node.Right
		}
	}
	return true
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

func main() {
	fmt.Println("Example1")
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	print(root, 0, 'M')
	fmt.Println("output:", isValidBST(root))

	fmt.Println("Example2")
	root = &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
	}
	print(root, 0, 'M')
	fmt.Println("output:", isValidBST(root))
}
