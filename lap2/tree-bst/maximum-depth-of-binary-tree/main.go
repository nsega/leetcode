package main

import "fmt"

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
	tree := &BinaryTree{}
	// FIXME: convert Tree from int array.
	tree.insert(9).insert(3).insert(20).insert(15).insert(7)
	tests := []test{
		{
			name: "Example 1:",
			arg: args{
				tree.root,
			},
			exp: 3,
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		print(tt.arg.root, 0, 'M')
		fmt.Println(" Actual: ", maxDepth(tt.arg.root))
		fmt.Println(" Exp: ", tt.exp)
	}
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
