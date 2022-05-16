package main

// Follow link to see details
// https://leetcode.com/problems/reverse-linked-list-ii/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	res := head
	var last *ListNode

	index := 0
	for head != nil {
		index++
		if index >= left {
			index--
			break
		}

		last = head
		head = head.Next
	}

	var reversed, lastReversed, prev *ListNode

	for head != nil {
		index++
		if index > right {
			break
		}
		next := head.Next

		reversed = head
		reversed.Next = prev

		prev = reversed

		if lastReversed == nil {
			lastReversed = prev
		}

		head = next
	}

	if last == nil {
		res = reversed
	} else {
		last.Next = reversed
	}

	lastReversed.Next = head

	return res
}
