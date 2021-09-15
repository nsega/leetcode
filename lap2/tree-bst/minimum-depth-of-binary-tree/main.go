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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)
	if leftDepth == 0 || rightDepth == 0 {
		return max(leftDepth, rightDepth) + 1
	}
	return min(leftDepth, rightDepth) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
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
			exp: 2,
		},
		{
			name: "Example 2:",
			arg: args{
				arrayToBinaryTree([]int{2, 0, 3, 0, 4, 0, 5, 0, 6}),
			},
			exp: 5,
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		print(tt.arg.root, 0, 'M')
		fmt.Println(" Actual: ", minDepth(tt.arg.root))
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
