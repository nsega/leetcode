package main

import "fmt"

type Node struct {
	key    int
	left   *Node
	right  *Node
	parent *Node
}

func FindLargestSmallerKey(rootNode *Node, num int) int {
	// your code goes here
	if num == 0 || rootNode == nil {
		return -1
	}

	if rootNode.key >= num {
		// move on left tree
		if rootNode.left == nil {
			return rootNode.key
		}
		return FindLargestSmallerKey(rootNode.left, num)
	} else {
		// move on right tree
		if rootNode.right == nil {
			return rootNode.key
		}
		return FindLargestSmallerKey(rootNode.right, num)
	}
}

// For example: input: 17, result 14
//     20
//    /  \
//   9    25
//  / \
// 5  12
//    / \
//  11  14

func main() {
	root := &Node{
		key:    20,
		left:   nil,
		right:  nil,
		parent: nil,
	}

	node9 := &Node{key: 9, left: nil, right: nil, parent: root}
	node25 := &Node{key: 25, left: nil, right: nil, parent: root}
	root.left = node9
	root.right = node25

	node5 := &Node{key: 5, left: nil, right: nil, parent: node9}
	node9.left = node5

	node12 := &Node{key: 12, left: nil, right: nil, parent: node9}
	node9.right = node12

	node11 := &Node{key: 11, left: nil, right: nil, parent: node12}
	node14 := &Node{key: 14, left: nil, right: nil, parent: node12}
	node12.left = node11
	node12.right = node14

	num := 17
	fmt.Println("input:", num)
	fmt.Println("exectation:", 14)
	fmt.Println("result:", FindLargestSmallerKey(root, num))
}
