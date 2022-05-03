package main

// Follow link to see details
// https://leetcode.com/problems/binary-tree-postorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int

	if root.Left != nil {
		res = append(res, postorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		res = append(res, postorderTraversal(root.Right)...)
	}

	res = append(res, root.Val)

	return res
}
