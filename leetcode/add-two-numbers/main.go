package main

import "fmt"

// Follow link to see details
// https://leetcode.com/problems/add-two-numbers/

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(items []int) *ListNode {
	if len(items) == 0 {
		return nil
	}

	head := &ListNode{
		Val: items[0],
	}

	iter := head
	for i := 1; i < len(items); i++ {
		iter.Next = &ListNode{
			Val: items[i],
		}

		iter = iter.Next
	}

	return head
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Print(" -> ")
		}

		node = node.Next
	}
	fmt.Println()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	res := &ListNode{}
	var iter *ListNode
	var additional int

	for l1 != nil && l2 != nil {
		total := l1.Val + l2.Val + additional

		if total >= 10 {
			total = total % 10
			additional = 1
		} else {
			additional = 0
		}

		if iter == nil {
			res.Val = total
			iter = res
		} else {
			iter.Next = &ListNode{
				Val: total,
			}
			iter = iter.Next
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		total := l1.Val + additional

		if total >= 10 {
			total = total % 10
			additional = 1
		} else {
			additional = 0
		}

		iter.Next = &ListNode{
			Val: total,
		}
		iter = iter.Next

		l1 = l1.Next
	}

	for l2 != nil {
		total := l2.Val + additional

		if total >= 10 {
			total = total % 10
			additional = 1
		} else {
			additional = 0
		}

		iter.Next = &ListNode{
			Val: total,
		}
		iter = iter.Next

		l2 = l2.Next
	}

	if additional > 0 {
		iter.Next = &ListNode{Val: additional}
	}

	return res
}

func main() {
	printList(addTwoNumbers(NewList([]int{2, 4, 3}), NewList([]int{5, 6, 4})))
	printList(addTwoNumbers(NewList([]int{0}), NewList([]int{0})))
	printList(addTwoNumbers(NewList([]int{9, 9, 9, 9, 9, 9, 9}), NewList([]int{9, 9, 9, 9})))
}
