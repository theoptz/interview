package main

import "fmt"

// Follow link to see details
// https://leetcode.com/problems/binary-tree-right-side-view/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)

	if root == nil {
		return res
	}

	levelStack := make([]*TreeNode, 0, 1)
	levelStack = append(levelStack, root)

	for len(levelStack) > 0 {
		currentLength := len(levelStack)

		res = append(res, levelStack[0].Val)

		for i := 0; i < currentLength; i++ {
			if levelStack[i].Right != nil {
				levelStack = append(levelStack, levelStack[i].Right)
			}

			if levelStack[i].Left != nil {
				levelStack = append(levelStack, levelStack[i].Left)
			}
		}

		levelStack = levelStack[currentLength:]
	}

	return res
}

func main() {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
		},
	}

	fmt.Println(rightSideView(tree))
}
