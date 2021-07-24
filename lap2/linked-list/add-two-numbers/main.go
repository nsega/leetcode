package main

import "fmt"

// ListNode is definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// addTwoNumbers is to add the two numbers and to return the sum as a linked list.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := ListNode{}
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

func main() {
	l1 := makeListNode([]int{2, 4, 3})
	l2 := makeListNode([]int{5, 6, 4})
	exp := []int{7, 0, 8}
	fmt.Println("Example 1:")
	fmt.Printf("  Input: l1 = %v, l2 = %v \n", l1.array(), l2.array())
	fmt.Println("  Expected:", exp)
	fmt.Println("  Output:", addTwoNumbers(l1, l2).array())

	l1 = makeListNode([]int{0})
	l2 = makeListNode([]int{0})
	exp = []int{0}
	fmt.Println("Example 1:")
	fmt.Printf("  Input: l1 = %v, l2 = %v \n", l1.array(), l2.array())
	fmt.Println("  Expected:", exp)
	fmt.Println("  Output:", addTwoNumbers(l1, l2).array())

	l1 = makeListNode([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = makeListNode([]int{9, 9, 9, 9})
	exp = []int{8, 9, 9, 9, 0, 0, 0, 1}
	fmt.Println("Example 1:")
	fmt.Printf("  Input: l1 = %v, l2 = %v \n", l1.array(), l2.array())
	fmt.Println("  Expected:", exp)
	fmt.Println("  Output:", addTwoNumbers(l1, l2).array())
}

// makeListNode makes ListNode with int array
func makeListNode(vals []int) *ListNode {
	list := &ListNode{Val: vals[0], Next: nil}
	for _, v := range vals[1:] {
		list.append(v)
	}
	return list
}

// append adds an Item to the end of the ListNode
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

// array returns int array from the ListNode
func (l *ListNode) array() []int {
	var arr []int
	for n := l; n != nil; n = n.Next {
		arr = append(arr, n.Val)
	}
	return arr
}
