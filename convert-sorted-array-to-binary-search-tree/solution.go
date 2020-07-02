package main

import (
	"fmt"
)

// TreeNode is a struct for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return bst(nums, 0, len(nums)-1)
}

func bst(nums []int, left, right int) *TreeNode {

	if left > right || left < 0 || right < 0 || right > len(nums) {
		return nil
	}

	if left == right {
		return &TreeNode{
			Val:   nums[left],
			Left:  nil,
			Right: nil,
		}
	}

	mid := (left + right) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  bst(nums, left, mid-1),
		Right: bst(nums, mid+1, right),
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

func main() {

	nums := []int{-10, -3, 0, 5, 9}
	print(sortedArrayToBST(nums), 0, 'M')
}
