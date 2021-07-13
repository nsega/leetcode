package main

import "fmt"

// ListNode is definition for singly-linked list.
type ListNode struct {
   Val int
   Next *ListNode
}

func hasCycle(head *ListNode) bool {
	visited := make(map[*ListNode]bool)
	for head != nil {
		if visited[head] {
			return true
		} else {
			visited[head] = true
		}
		head = head.Next
	}
	return false
}

func main() {
	// Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
	vals := []int{3,2,0,4}
	pos := 1
	head := makeListNode(vals, pos)
	fmt.Println("Example 1:")
	fmt.Printf("  Input: head = %v, pos = %d\n", vals, pos)
	fmt.Println("  Expected output: ", true)
	fmt.Println("  Actual output:   ", hasCycle(head))

	// Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
	vals = []int{1,2}
	pos = 0
	head = makeListNode(vals, pos)
	fmt.Println("Example 2:")
	fmt.Printf("  Input: head = %v, pos = %d\n", vals, pos)
	fmt.Println("  Expected output: ", true)
	fmt.Println("  Actual output:   ", hasCycle(head))

	// Explanation: There is no cycle in the linked list.
	vals = []int{1}
	pos = -1
	head = makeListNode(vals, pos)
	fmt.Println("Example 3:")
	fmt.Printf("  Input: head = %v, pos = %d\n", vals, pos)
	fmt.Println("  Expected output: ", false)
	fmt.Println("  Actual output:   ", hasCycle(head))
}

func makeListNode(vals []int, pos int) *ListNode{
	head := &ListNode{
		Val: vals[0],
		Next: nil,
	}
	lastNode := &ListNode{}
	for val := range vals[1:] {
		tmp := head
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		tmp.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		lastNode = tmp.Next
	}
	if pos != -1 {
		lastNode.Next = node(head, pos)
	}
	return head
}

func node(ln *ListNode, pos int) *ListNode {
	p := 0
	ptr := ln
	for p < pos {
		ptr = ptr.Next
		p += 1
	}
	return ptr
}