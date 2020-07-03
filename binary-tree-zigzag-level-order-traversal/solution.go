package main

import "fmt"

// TreeNode is a struct for a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	return bfs(root)
}

func bfs(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	level := 0
	for len(queue) > 0 {
		currentLength := len(queue)
		var tmp []int
		for i := 0; i < currentLength; i++ {
			n := queue[i]
			tmp = append(tmp, n.Val)
			if n.Left != nil {
				queue = append(queue, n.Left)
			}
			if n.Right != nil {
				queue = append(queue, n.Right)
			}
		}
		if level%2 == 1 {
			reverse(tmp)
		}
		queue = queue[currentLength:]
		ret = append(ret, tmp)
		level++
	}

	return ret
}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println("output:", zigzagLevelOrder(root))
}
