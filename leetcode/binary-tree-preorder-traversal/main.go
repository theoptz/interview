package main

// Follow link to see details
// https://leetcode.com/problems/binary-tree-preorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	res = append(res, root.Val)

	if root.Left != nil {
		res = append(res, preorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		res = append(res, preorderTraversal(root.Right)...)
	}

	return res
}
