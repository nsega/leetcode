package main

import "fmt"

func main() {

	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	exp := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}

	fmt.Println("Example:")
	fmt.Printf("  Input: l1 = %v l2 = %v \n", l1, l2)
	fmt.Printf("  Expected output: %v \n", exp)
	fmt.Println("  Actual output: ", addTwoNumbers(l1, l2))

}

// ListNode is a struct of definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var head ListNode
	current := &head
	for l1 != nil || l2 != nil || carry != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		num := n1 + n2 + carry
		carry = num / 10
		current.Next = &ListNode{
			Val:  num % 10,
			Next: nil,
		}
		current = current.Next
	}
	return head.Next
}
