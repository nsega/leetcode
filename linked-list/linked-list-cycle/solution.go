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
	lastNode := node(head, 3)
	pos := 1
	lastNode.Next = node(head, pos)

	fmt.Println("Example 1:")
	fmt.Printf("  Input: s = %v, pos = %d\n", vals, pos)
	fmt.Println("  Expected output: ", true)
	fmt.Println("  Actual output:   ", hasCycle(head))

	vals = []int{1, 2}
	head = makeList(vals)
	lastNode = node(head, 1)
	pos = 0
	lastNode.Next = node(head, pos)

	fmt.Println("Example 2:")
	fmt.Printf("  Input: s = %v, pos = %d\n", vals, pos)
	fmt.Println("  Expected output: ", true)
	fmt.Println("  Actual output:   ", hasCycle(head))

	vals = []int{1}
	head = makeList(vals)
	lastNode = node(head, 0)
	pos = -1
	lastNode.Next = node(head, pos)

	fmt.Println("Example 3:")
	fmt.Printf("  Input: s = %v, pos = %d\n", vals, pos)
	fmt.Println("  Expected output: ", false)
	fmt.Println("  Actual output:   ", hasCycle(head))
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
			p += 1
		}
		return ptr
	}
	return nil
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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
