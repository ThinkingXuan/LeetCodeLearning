package main

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseKGroup k个一组链表反转
func reverseKGroup(head *ListNode, k int) *ListNode {

	hair := &ListNode{Next: head}

	pre := hair

	for head != nil {
		tail := pre

		for i := 0; i < k; i++ {
			tail = tail.Next

			if tail == nil {
				return hair.Next
			}
		}

		next := tail.Next
		head, tail = reverse(head, tail)
		pre.Next = head
		tail.Next = next

		pre = tail
		head = pre.Next

	}
	return hair.Next
}
func reverse(head, tail *ListNode) (*ListNode, *ListNode) {

	var pre *ListNode
	cur := head

	for pre != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return tail, head
}
