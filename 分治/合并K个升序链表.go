package main

//Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return mergesort(lists, 0, len(lists)-1)
}

func mergesort(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	if left > right {
		return nil
	}
	mid := left + (right-left)>>1
	return mergeTwoList(mergesort(lists, left, mid), mergesort(lists, mid+1, right))
}

func mergeTwoList(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
			p = p.Next
		} else {
			p.Next = list2
			list2 = list2.Next
			p = p.Next
		}
	}

	if list1 == nil {
		p.Next = list2
	} else {
		p.Next = list1
	}
	return head.Next
}
