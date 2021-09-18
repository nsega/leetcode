package main

import (
	"container/list"
	"fmt"
)

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// hasPathSum is to return true if the tree has a root-to-leaf path
// such that adding up all the values along the path equals targetSum
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func main() {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	type test struct {
		name string
		arg  args
		exp  bool
	}
	tests := []test{
		{
			name: "Example 1:",
			arg: args{
				root:      arrayToBinaryTree([]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 1}),
				targetSum: 22,
			},
			exp: true,
		},
		{
			name: "Example 2:",
			arg: args{
				root:      arrayToBinaryTree([]int{1, 2, 3}),
				targetSum: 5,
			},
			exp: false,
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		print(tt.arg.root, 0, 'M')
		fmt.Println(" Actual: ", hasPathSum(tt.arg.root, tt.arg.targetSum))
		fmt.Println(" Exp: ", tt.exp)
	}
}

func arrayToBinaryTree(tree []int) *TreeNode {
	if len(tree) == 0 {
		return nil
	}
	root := &TreeNode{Val: tree[0], Left: nil, Right: nil}
	l := list.New()
	l.PushBack(root)
	for i := 1; i < len(tree); i++ {
		node := l.Front().Value.(*TreeNode)
		if ((i + 1) % 2) == 0 {
			if tree[i] > 0 {
				node.Left = &TreeNode{Val: tree[i], Left: nil, Right: nil}
				l.PushBack(node.Left)
			} else {
				node.Left = nil
			}
		} else if ((i + 1) % 2) == 1 {
			if tree[i] > 0 {
				node.Right = &TreeNode{Val: tree[i], Left: nil, Right: nil}
				l.PushBack(node.Right)
			} else {
				node.Right = nil
			}
			l.Remove(l.Front())
		}
	}
	return root
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
