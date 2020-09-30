package main

import "fmt"

// ListNode is definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l ListNode
	l = ListNode{Val: 1, Next: nil}
	vals := []int{1, 2}
	for _, v := range vals {
		l.Append(v)
	}
	expOut := []int{1, 2}
	fmt.Println("Example 1:")
	fmt.Printf("  Input: s = %v\n", []int{1, 1, 2})
	fmt.Println("  Expected output: ", expOut)
	out := deleteDuplicates(&l)
	var ans []int
	for n := out; n != nil; n = n.Next {
		ans = append(ans, n.Val)
	}
	fmt.Println("  Actual output: ", ans)

	l = ListNode{Val: 1, Next: nil}
	vals = []int{1, 2, 3, 3}
	for _, v := range vals {
		l.Append(v)
	}
	expOut = []int{1, 2, 3}
	fmt.Println("Example 2:")
	fmt.Printf("  Input: s = %v\n", []int{1, 1, 2, 3, 3})
	fmt.Println("  Expected output: ", expOut)
	out = deleteDuplicates(&l)
	var a []int
	for n := out; n != nil; n = n.Next {
		a = append(a, n.Val)
	}
	fmt.Println("  Actual output: ", a)

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		if current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

// Append adds an Item to the end of the linked list
func (l *ListNode) Append(v int) {
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
