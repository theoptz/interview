package main

// Follow link to see details
// https://leetcode.com/problems/reverse-linked-list-ii/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := head
	prevNthNode := head
	var index int

	for head != nil {
		index++
		if index > n+1 {
			prevNthNode = prevNthNode.Next
		}

		head = head.Next
	}

	if index > n {
		prevNthNode.Next = prevNthNode.Next.Next
	} else {
		res = res.Next
	}

	return res
}
