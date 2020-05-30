/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    
    visited := make(map[*ListNode]bool)
    var count int
    for n := head; n != nil; n = n.Next {
        if visited[n] {
            return n
        }
        visited[n] = true
        count++
    }    
    return nil
}
