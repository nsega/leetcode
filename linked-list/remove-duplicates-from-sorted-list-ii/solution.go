package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	vals := []int{2, 3, 3, 4, 4, 5}
	var l ListNode
	l = ListNode{Val: 1, Next: nil}
	for _, v := range vals {
		l.Append(v)
	}
	fmt.Println("Example 1:")
	fmt.Printf("  Input: ")
	printListNode(&l)
	fmt.Println("  Expected output:  1->2->5")
	out := deleteDuplicates(&l)
	fmt.Printf("  Actual output:  ")
	printListNode(out)

	vals = []int{1, 1, 2, 3}
	l = ListNode{Val: 1, Next: nil}
	for _, v := range vals {
		l.Append(v)
	}
	fmt.Println("Example 2:")
	fmt.Printf("  Input: ")
	printListNode(&l)
	fmt.Println("  Expected output:  2->3")
	out = deleteDuplicates(&l)
	fmt.Printf("  Actual output:  ")
	printListNode(out)
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

func printListNode(in *ListNode) {
	for in != nil {
		fmt.Printf("%d", in.Val)
		if in.Next != nil {
			fmt.Printf("->")
		}
		in = in.Next
	}
	fmt.Printf("\n")
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
	var temp *ListNode
	for current != nil {
		if current.Next != nil && current.Val == current.Next.Val {
			if temp != nil {
				for current.Next != nil && current.Val == current.Next.Val {
					current = current.Next
				}
				temp.Next = current.Next
				current = current.Next
				continue
			} else {
				for current.Next != nil && current.Val == current.Next.Val {
					current = current.Next
				}
				head = current.Next
				current = current.Next
				continue
			}
		}
		temp = current
		current = current.Next
	}
	return head
}
