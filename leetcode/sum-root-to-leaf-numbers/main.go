package main

import "fmt"

// Follow link to see details
// https://leetcode.com/problems/binary-tree-right-side-view/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	sum := 0
	current := 0

	sumNumbersHelper(root, current, &sum)

	return sum
}

func sumNumbersHelper(node *TreeNode, currentSum int, sum *int) {
	if node == nil {
		*sum += currentSum
		return
	}

	currentSum = currentSum*10 + node.Val

	if node.Left == nil && node.Right == nil {
		*sum += currentSum
		return
	}

	if node.Left != nil {
		sumNumbersHelper(node.Left, currentSum, sum)
	}

	if node.Right != nil {
		sumNumbersHelper(node.Right, currentSum, sum)
	}
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(sumNumbers(tree))
}
