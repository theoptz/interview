package main

// Follow link to see details
// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTHelper(0, len(nums)-1, nums)
}

func sortedArrayToBSTHelper(from, to int, nums []int) *TreeNode {
	if from > to {
		return nil
	}

	middle := from + (to-from)/2
	val := &TreeNode{
		Val: nums[middle],
	}

	if from == to {
		return val
	}

	val.Left = sortedArrayToBSTHelper(from, middle-1, nums)
	val.Right = sortedArrayToBSTHelper(middle+1, to, nums)

	return val
}
