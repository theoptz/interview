package main

import "math"

// Follow link to see details
// https://leetcode.com/problems/validate-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return validate(root.Left, min, root.Val) && validate(root.Right, root.Val, max)
}
