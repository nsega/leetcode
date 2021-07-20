package main

import "fmt"

// ListNode is definition for singly-linked list.
type ListNode struct {
	Val  int
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
	vals := []int{3, 2, 0, 4}
	pos := 1
	head := makeListNode(vals, pos)
	fmt.Println("Example 1:")
	fmt.Printf("  Input: head = %v, pos = %d\n", vals, pos)
	fmt.Println("  Expected output: ", true)
	fmt.Println("  Actual output:   ", hasCycle(head))

	// Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
	vals = []int{1, 2}
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

func makeListNode(vals []int, pos int) *ListNode {
	list := &ListNode{Val: vals[0], Next: nil}
	for _, v := range vals[1:] {
		list.append(v)
	}

	tmp := list
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	lastNode := tmp
	lastNode.Next = node(list, pos)

	return list
}

func node(ln *ListNode, pos int) *ListNode {
	if pos != -1 {
		p := 0
		ptr := ln
		for p < pos {
			ptr = ptr.Next
			p += 1
		}
		return ptr
	}
	return nil
}

// Append adds an Item to the end of the linked list
func (l *ListNode) append(v int) {
	node := ListNode{Val: v, Next: nil}
	if l.Next == nil {
		l.Next = &node
	} else {
		last := l.Next
		for {
			if last.Next == nil {
				break
			}
			last = last.Next
		}
		last.Next = &node
	}
}
