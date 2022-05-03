package main

// Follow link to see details
// https://leetcode.com/problems/binary-tree-level-order-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)

	levelOrderHelper(root, 1, &res)

	return res
}

func levelOrderHelper(node *TreeNode, level int, res *[][]int) {
	if node == nil {
		return
	}

	if len(*res) < level {
		*res = append(*res, make([]int, 0))
	}

	(*res)[level-1] = append((*res)[level-1], node.Val)

	levelOrderHelper(node.Left, level+1, res)
	levelOrderHelper(node.Right, level+1, res)
}
