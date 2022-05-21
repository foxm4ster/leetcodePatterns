package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	next := head.Next
	for next != nil && next.Next != nil {
		if head == next {
			return true
		}
		head = head.Next
		next = next.Next.Next
	}

	return false
}
