package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	mid := head
	for head != nil && head.Next != nil {
		mid = mid.Next
		head = head.Next.Next
	}
	return mid
}
