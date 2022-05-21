package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var (
		prev *ListNode
		curr = head
		next = head
	)

	for next != nil {
		next = next.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
