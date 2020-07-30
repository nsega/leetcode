/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
        
    visited := make(map[*ListNode]bool)
    for head != nil {
        if visited[head] {
            return true
        } else {
            visited[head] = true            
        }
        head = head.Next           
    }
    return false           
}
