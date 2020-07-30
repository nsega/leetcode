package main

import "fmt"

// TreeNode is a struct for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	return bfs(root)
}

func bfs(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int
	var queue []*TreeNode
	queue = append(queue, root)
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
		queue = queue[currentLength:]
		ret = append(ret, tmp)
	}
	return ret
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
	fmt.Println("output:", levelOrder(root))

}
