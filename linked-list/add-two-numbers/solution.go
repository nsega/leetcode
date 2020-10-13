package main

import "fmt"

func main() {

	l1 := &ListNode{Val: 2, Next: nil}
	for _, v := range [2]int{4, 3} {
		l1.Append(v)
	}
	l2 := &ListNode{Val: 5, Next: nil}
	for _, v := range [2]int{6, 4} {
		l2.Append(v)
	}
	exp := []int{7, 0, 8}
	fmt.Println("Example 1:")
	fmt.Printf("  Input: l1 = %v l2 = %v \n", []int{2, 4, 3}, []int{5, 6, 4})
	fmt.Printf("  Expected output: %v \n", exp)
	out := addTwoNumbers(l1, l2)
	var ans1 []int
	for n := out; n != nil; n = n.Next {
		ans1 = append(ans1, n.Val)
	}
	fmt.Println("  Actual output:  ", ans1)

	l1 = &ListNode{Val: 0, Next: nil}
	l2 = &ListNode{Val: 0, Next: nil}
	exp = []int{0}
	fmt.Println("Example 2:")
	fmt.Printf("  Input: l1 = %v l2 = %v \n", [1]int{0}, [1]int{0})
	fmt.Printf("  Expected output: %v \n", exp)
	out = addTwoNumbers(l1, l2)
	var ans2 []int
	for n := out; n != nil; n = n.Next {
		ans2 = append(ans2, n.Val)
	}
	fmt.Println("  Actual output:  ", ans2)

	l1 = &ListNode{Val: 9, Next: nil}
	for _, v := range [6]int{9, 9, 9, 9, 9, 9} {
		l1.Append(v)
	}
	l2 = &ListNode{Val: 9, Next: nil}
	for _, v := range [3]int{9, 9, 9} {
		l2.Append(v)
	}
	exp = []int{8, 9, 9, 9, 0, 0, 0, 1}
	fmt.Println("Example 3:")
	fmt.Printf("  Input: l1 = %v l2 = %v \n", [7]int{9, 9, 9, 9, 9, 9, 9}, [4]int{9, 9, 9, 9})
	fmt.Printf("  Expected output: %v \n", exp)
	out = addTwoNumbers(l1, l2)
	var ans3 []int
	for n := out; n != nil; n = n.Next {
		ans3 = append(ans3, n.Val)
	}
	fmt.Println("  Actual output:  ", ans3)
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
