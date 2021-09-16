package main

import "fmt"

// TreeNode is a struct of a node of a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// sortedArrayToBST is to convert an integer array nums where the elements
// are sorted in ascending order to a height-balanced binary search tree
func sortedArrayToBST(nums []int) *TreeNode {
	return bst(nums, 0, len(nums)-1)
}

// bst is to convert an integer array nums where the elements
// are sorted in ascending order to a height-balanced binary search tree
func bst(nums []int, left, right int) *TreeNode {
	if left > right || left < 0 || right < 0 || right > len(nums) {
		return nil
	}
	if left == right {
		return &TreeNode{Val: nums[left], Left: nil, Right: nil}
	}
	mid := (left + right) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  bst(nums, left, mid-1),
		Right: bst(nums, mid+1, right),
	}
}

func main() {
	type args struct {
		nums []int
	}
	type test struct {
		name string
		arg  args
		exp  *TreeNode
	}

	exp1 := &BinaryTree{}
	exp1.insert(0).insert(-10).insert(5).insert(-3).insert(9)
	exp2 := &BinaryTree{}
	exp2.insert(1).insert(3)

	tests := []test{
		{
			name: "Example 1:",
			arg: args{
				nums: []int{-10, -3, 0, 5, 9},
			},
			exp: exp1.root,
		},
		{
			name: "Example 2:",
			arg: args{
				nums: []int{1, 3},
			},
			exp: exp2.root,
		},
	}
	for _, tt := range tests {
		fmt.Println("Case:", tt.name)
		fmt.Println("Actual: ")
		print(sortedArrayToBST(tt.arg.nums), 0, 'M')
		fmt.Println("Exp: ")
		print(tt.exp, 0, 'M')
	}
}

// BinaryTree is a struct of a binary tree
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
