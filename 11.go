package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func isPalindrome(head *ListNode) bool {
	mid := middleNode(head)
	mid = reverseList(mid)

	for head != nil && mid != nil {
		if head.Val != mid.Val {
			return false
		}
		head = head.Next
		mid = mid.Next
	}

	return true
}

func middleNode(head *ListNode) *ListNode {
	mid := head
	for head != nil && head.Next != nil {
		mid = mid.Next
		head = head.Next.Next
	}
	return mid
}

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
