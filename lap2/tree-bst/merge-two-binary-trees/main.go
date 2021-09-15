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

// mergeTrees is to merge root1 and root2 and return the merged tree node.
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	return &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  mergeTrees(root1.Left, root2.Left),
		Right: mergeTrees(root1.Right, root2.Right),
	}
}

func main() {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	type test struct {
		name string
		arg  args
		exp  *TreeNode
	}
	tests := []test{
		{
			name: "Example 1:",
			arg: args{
				root1: arrayToBinaryTree([]int{1, 3, 2, 5}),
				root2: arrayToBinaryTree([]int{2, 1, 3, 0, 4, 0, 7}),
			},
			exp: arrayToBinaryTree([]int{3, 4, 5, 5, 4, 0, 7}),
		},
		{
			name: "Example 2:",
			arg: args{
				root1: arrayToBinaryTree([]int{1}),
				root2: arrayToBinaryTree([]int{1, 2}),
			},
			exp: arrayToBinaryTree([]int{2, 2}),
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		print(tt.arg.root1, 0, 'M')
		print(tt.arg.root2, 0, 'M')
		fmt.Println(" Actual: ")
		print(mergeTrees(tt.arg.root1, tt.arg.root2), 0, 'M')
		fmt.Println(" Exp: ")
		print(tt.exp, 0, 'M')
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
