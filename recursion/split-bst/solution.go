package main

import "fmt"

func main() {

	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
			Right: &TreeNode{Val: 3, Left: nil, Right: nil},
		},
		Right: &TreeNode{
			Val:   6,
			Left:  &TreeNode{Val: 5, Left: nil, Right: nil},
			Right: &TreeNode{Val: 7, Left: nil, Right: nil},
		},
	}
	V := 2
	fmt.Println("Example 1:")
	fmt.Printf("Input:\n")
	print(root, 2, 'M')
	fmt.Printf("  V: %d\n", V)
	fmt.Println("Expected output: [[2,1],[4,3,6,null,null,5,7]]")
	fmt.Println("Actual output: ")
	for _, bs := range splitBST(root, V) {
		print(bs, 2, 'M')
	}

}

// TreeNode is a struct for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func splitBST(root *TreeNode, V int) []*TreeNode {

	if root == nil {
		return []*TreeNode{nil, nil}
	} else if root.Val <= V {
		bs := splitBST(root.Right, V)
		root.Right = bs[0]
		bs[0] = root
		return bs
	} else {
		bs := splitBST(root.Left, V)
		root.Left = bs[1]
		bs[1] = root
		return bs
	}
}

func print(node *TreeNode, ns int, ch rune) {
	if node == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%c:%v\n", ch, node.Val)
	print(node.Left, ns+2, 'L')
	print(node.Right, ns+2, 'R')
}
