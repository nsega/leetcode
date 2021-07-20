package main

import "fmt"

// ListNode is Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// detectCycle is return the node where the cycle begins.
// If there is no cycle, return null.
func detectCycle(head *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)
	for n := head; n != nil; n = n.Next {
		if visited[n] {
			return n
		}
		visited[n] = true
	}
	return nil
}

func main() {
	// Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
	vals := []int{3, 2, 0, -4}
	pos := 1
	head := makeLinkedList(vals, pos)
	fmt.Println("Example 1:")
	fmt.Printf("  Input: head = %v, pos = %d\n", vals, pos)
	fmt.Println("  Expected output: tail connects to node index", 1)
	got := detectCycle(head)
	if got != nil {
		fmt.Println("  Actual output: tail connects to node index", index(vals, got.Val))
	} else {
		fmt.Println("  Actual output: ", "nocycle")
	}

	// Explanation: There is a cycle in the linked list, where tail connects to the first node.
	vals = []int{1, 2}
	pos = 0
	head = makeLinkedList(vals, pos)
	fmt.Println("Example 2:")
	fmt.Printf("  Input: head = %v, pos = %d\n", vals, pos)
	fmt.Println("  Expected output: tail connects to node index", 0)
	got = detectCycle(head)
	if got != nil {
		fmt.Println("  Actual output: tail connects to node index", index(vals, got.Val))
	} else {
		fmt.Println("  Actual output: ", "nocycle")
	}

	// Explanation: There is no cycle in the linked list.
	vals = []int{1}
	pos = -1
	head = makeLinkedList(vals, pos)
	fmt.Println("Example 3:")
	fmt.Printf("  Input: head = %v, pos = %d\n", vals, pos)
	fmt.Println("  Expected output: nocycle")
	got = detectCycle(head)
	if got != nil {
		fmt.Println("  Actual output: tail connects to node index", index(vals, got.Val))
	} else {
		fmt.Println("  Actual output: nocycle")
	}
}

func makeLinkedList(vals []int, pos int) *ListNode {
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

func index(vals []int, val int) int {
	idx := -1
	for i, v := range vals {
		if v == val {
			idx = i
		}
	}
	return idx
}

func node(head *ListNode, pos int) *ListNode {
	if pos != -1 {
		p := 0
		ptr := head
		for p < pos {
			ptr = ptr.Next
			p += 1
		}
		return ptr
	}
	return nil
}

// append adds an Item to the end of the linked list
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
