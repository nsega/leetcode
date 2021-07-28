package main

import "fmt"

// ListNode is Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList is to reverse the list and to return the reversed list.
func reverseList(head *ListNode) *ListNode {
	var reverse *ListNode
	for head != nil {
		temp := head.Next
		head.Next = reverse
		reverse = head
		head = temp
	}
	return reverse
}

func main() {
	type args struct {
		testName string
		head     *ListNode
		exp      *ListNode
	}
	tests := []args{
		{
			testName: "Example 1",
			head:     makeListNode([]int{1, 2, 3, 4, 5}),
			exp:      makeListNode([]int{5, 4, 3, 2, 1}),
		},
		{
			testName: "Example 2",
			head:     makeListNode([]int{1, 2}),
			exp:      makeListNode([]int{2, 1}),
		},
		{
			testName: "Example 3",
			head:     makeListNode([]int{}),
			exp:      makeListNode([]int{}),
		},
	}
	for _, tt := range tests {
		fmt.Println(tt.testName)
		fmt.Printf("  Input: head = %v\n", tt.head.array())
		fmt.Println("  Expected:", tt.exp.array())
		fmt.Println("  Output:", reverseList(tt.head).array())
	}
}

// makeListNode makes ListNode with int array
func makeListNode(vals []int) *ListNode {
	switch len(vals) {
	case 0:
		return nil
	case 1:
		return &ListNode{Val: vals[0], Next: nil}
	default:
		list := &ListNode{Val: vals[0], Next: nil}
		for _, v := range vals[1:] {
			list.append(v)
		}
		return list
	}
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
