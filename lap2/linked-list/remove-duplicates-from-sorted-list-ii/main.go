package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteDuplicates deletes all nodes that have duplicate numbers,
// leaving only distinct numbers from the original list.
// Return the linked list sorted as well.
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	var tmp *ListNode
	for current != nil {
		if current.Next != nil && current.Val == current.Next.Val {
			if tmp != nil {
				for current.Next != nil && current.Val == current.Next.Val {
					current = current.Next
				}
				tmp.Next = current.Next
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
		tmp = current
		current = current.Next
	}
	return head
}

func main() {
	vals := []int{1, 2, 3, 3, 4, 4, 5}
	exp := []int{1, 2, 5}
	head := makeListNode(vals)
	fmt.Println("Example 1:")
	fmt.Printf("  Input: head = %v\n", vals)
	fmt.Println("  Expected:", exp)
	fmt.Println("  Output:", deleteDuplicates(head).array())

	vals = []int{1, 1, 1, 2, 3}
	exp = []int{2, 3}
	head = makeListNode(vals)
	fmt.Println("Example 2:")
	fmt.Printf("  Input: head = %v\n", vals)
	fmt.Println("  Expected:", exp)
	fmt.Println("  Output:", deleteDuplicates(head).array())
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
