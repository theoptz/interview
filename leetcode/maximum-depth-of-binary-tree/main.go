package main

// Follow link to see details
// https://leetcode.com/problems/maximum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftVal := maxDepth(root.Left)
	rightVal := maxDepth(root.Right)

	if leftVal > rightVal {
		return 1 + leftVal
	}

	return 1 + rightVal
}
