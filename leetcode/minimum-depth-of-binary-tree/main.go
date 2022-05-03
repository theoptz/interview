package main

// Follow link to see details
// https://leetcode.com/problems/minimum-depth-of-binary-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftVal := minDepth(root.Left)
	rightVal := minDepth(root.Right)

	if leftVal == 0 || rightVal == 0 {
		return 1 + leftVal + rightVal
	}

	if leftVal < rightVal {
		return 1 + leftVal
	}

	return 1 + rightVal
}
