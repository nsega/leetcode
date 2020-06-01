/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
    
    previous := head
    current := head
        
    for previous != nil && previous.Next != nil {
        
        if current != nil && current.Next != nil && current.Next.Val == current.Val {
            
            for current != nil && current.Next != nil && current.Next.Val == current.Val {
                current.Next = current.Next.Next                                
            }

            if current == head {
                head = head.Next
                previous = head
                current = head
            } else {
                current = current.Next
                previous.Next = current
            }            
        } else {
            previous = current
            current = current.Next
        }
    }
    return head
    
}
