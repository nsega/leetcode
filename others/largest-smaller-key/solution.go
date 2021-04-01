package main

import "fmt"

type Node struct {
	key int
	left *Node
	right *Node
	parent *Node
}

// implemented by you
func FindLargestSmallerKey(rootNode *Node, num int) int {
	// your code goes here
	var min int
	min = -1
	if (rootNode == nil) {
		return -1
	}
	if (rootNode.key == num) {
		return num
	}

	if (num > rootNode.key) {
		min = FindLargestSmallerKey(rootNode.right, num)
		if (min == -1) {
			return rootNode.key
		}
	} else {
		min = FindLargestSmallerKey(rootNode.left, num)
		if min == -1 {
			return rootNode.key
		}
	}

	return min
}

// implemented by me
func FindLargestSmallerKey2(rootNode *Node, num int) int {
	if rootNode == nil {
		return -1
	} else if rootNode.key < num {
		tmp := FindLargestSmallerKey(rootNode.right,num)
		if tmp != -1 {
			return tmp
		}
		return rootNode.key
	} else {
		return FindLargestSmallerKey(rootNode.left, num)
	}
}


func main() {
	root := &Node{
		key: 20, left: nil, right: nil, parent: nil,
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

	var args = []struct {
		input int
		want  int
	}{
		{17, 14},
		{26, 25},
		{19, 14},
		{6, 5},
		{5, -1},
		{1, -1},
	}

	fmt.Println("-- FindLargestSmallerKey --")
	for _, arg := range args {
		fmt.Println("---------------")
		fmt.Println("input:", arg.input)
		fmt.Println("want:", arg.want)
		got := FindLargestSmallerKey(root, arg.input)
		fmt.Println("got:", got)
		if arg.want == got {
			fmt.Println("test result: OK")
		} else {
			fmt.Println("test result: NG")
		}
	}

	fmt.Println("-- FindLargestSmallerKey2 --")
	for _, arg := range args {
		fmt.Println("---------------")
		fmt.Println("input:", arg.input)
		fmt.Println("want:", arg.want)
		got := FindLargestSmallerKey2(root, arg.input)
		fmt.Println("got:", got)
		if arg.want == got {
			fmt.Println("test result: OK")
		} else {
			fmt.Println("test result: NG")
		}
	}
}
