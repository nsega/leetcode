package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	righDepth := maxDepth(root.Right)
	if leftDepth > righDepth {
		return leftDepth + 1
	} else {
		return righDepth + 1
	}
}

func main() {
	type args struct {
		root *TreeNode
	}
	type test struct {
		name string
		arg  args
		exp  int
	}
	tests := []test{
		{
			name: "Example 1:",
			arg: args{
				arrayToBinaryTree([]int{3, 9, 20, 0, 0, 15, 7}),
			},
			exp: 3,
		},
		{
			name: "Example 2:",
			arg: args{
				arrayToBinaryTree([]int{1, 0, 2}),
			},
			exp: 2,
		},
		{
			name: "Example 3:",
			arg: args{
				arrayToBinaryTree([]int{}),
			},
			exp: 0,
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		print(tt.arg.root, 0, 'M')
		fmt.Println(" Actual: ", maxDepth(tt.arg.root))
		fmt.Println(" Exp: ", tt.exp)
	}
}

func arrayToBinaryTree(tree []int) *TreeNode {
	if len(tree) == 0 {
		return nil
	}
	root := &TreeNode{Val: tree[0], Left: nil, Right: nil}
	l := list.New()
	l.PushFront(root)
	for i := 1; i < len(tree); i++ {
		node := l.Front().Value.(*TreeNode)
		if node.Left == nil {
			node.Left = &TreeNode{Val: tree[i], Left: nil, Right: nil}
			if tree[i] > 0 {
				l.PushBack(node.Left)
			}
		} else if node.Right == nil {
			node.Right = &TreeNode{Val: tree[i], Left: nil, Right: nil}
			if tree[i] > 0 {
				l.PushBack(node.Right)
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
