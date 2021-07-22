package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// deleteDuplicates is to delete all duplicates such that each element appears only once.
// Return the linked list sorted as well.
func deleteDuplicates(head *ListNode) *ListNode {
	tmp := head
	for tmp != nil && tmp.Next != nil {
		if tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return head
}

func main() {

	vals := []int{1, 1, 2}
	exp := []int{1, 2}
	head := makeListNode(vals)
	fmt.Println("Example 1:")
	fmt.Printf("  Input: head = %v\n", vals)
	fmt.Println("  Expected output:", exp)
	fmt.Println("  Actual output:", deleteDuplicates(head).array())

	vals = []int{1, 1, 2, 3, 3}
	exp = []int{1, 2, 3}
	head = makeListNode(vals)
	fmt.Println("Example 2:")
	fmt.Printf("  Input: head = %v\n", vals)
	fmt.Println("  Expected output:", exp)
	fmt.Println("  Actual output:", deleteDuplicates(head).array())
}

func makeListNode(vals []int) *ListNode {
	list := &ListNode{Val: vals[0], Next: nil}
	for _, v := range vals[1:] {
		list.append(v)
	}
	return list
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

func (l *ListNode) array() []int {
	var arr []int
	for n := l; n != nil; n = n.Next {
		arr = append(arr, n.Val)
	}
	return arr
}

func (l *ListNode) printListNode() {
	ln := l
	for ln != nil {
		fmt.Printf("%d", ln.Val)
		if ln.Next != nil {
			fmt.Printf("->")
		}
		ln = ln.Next
	}
	fmt.Printf("\n")
}
