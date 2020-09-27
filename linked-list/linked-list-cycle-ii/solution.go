package main

import "fmt"

// ListNode is definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	vals := []int{3, 2, 0, -4}
	head := makeList(vals)
	lastNode := node(head, len(vals)-1)
	pos := 1
	lastNode.Next = node(head, pos)

	fmt.Println("Example 1:")
	fmt.Printf("  Input: s = %v, pos = %d\n", vals, pos)
	fmt.Println("  Expected output: tail connects to node index", pos)
	out := detectCycle(head)
	if out == nil {
		fmt.Println("  Actual output: no cycle")
	} else {
		fmt.Println("  Actual output: tail connects to node index", indexVal(vals, out))
	}

	vals = []int{1, 2}
	head = makeList(vals)
	lastNode = node(head, len(vals)-1)
	pos = 0
	lastNode.Next = node(head, pos)

	fmt.Println("Example 2:")
	fmt.Printf("  Input: s = %v, pos = %d\n", vals, pos)
	fmt.Println("  Expected output: tail connects to node index", pos)
	out = detectCycle(head)
	if out == nil {
		fmt.Println("  Actual output: no cycle")
	} else {
		fmt.Println("  Actual output: tail connects to node index", indexVal(vals, out))
	}

	vals = []int{1}
	head = makeList(vals)
	lastNode = node(head, len(vals)-1)
	pos = -1
	lastNode.Next = node(head, pos)

	fmt.Println("Example 3:")
	fmt.Printf("  Input: s = %v, pos = %d\n", vals, pos)
	fmt.Println("  Expected output: no cycle")
	out = detectCycle(head)
	if out == nil {
		fmt.Println("  Actual output: no cycle")
	} else {
		fmt.Println("  Actual output: tail connects to node index", indexVal(vals, out))
	}
}

// detectCycle returns the node where the cycle begins. If there is no cycle, return null.
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

func makeList(vals []int) *ListNode {
	head := &ListNode{
		Val:  vals[0],
		Next: nil,
	}
	for val := range vals[1:] {
		ptr := head
		for ptr.Next != nil {
			ptr = ptr.Next
		}
		ptr.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
	}
	return head
}

func node(head *ListNode, pos int) *ListNode {
	if pos != -1 {
		p := 0
		ptr := head
		for p < pos {
			ptr = ptr.Next
			p++
		}
		return ptr
	}
	return nil
}

func indexVal(vals []int, node *ListNode) int {
	for i, v := range vals {
		if v == node.Val {
			return i
		}
	}
	return -1
}
