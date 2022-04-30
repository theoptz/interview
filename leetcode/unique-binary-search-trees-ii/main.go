package main

import "fmt"

// Follow link to see details
// https://leetcode.com/problems/unique-binary-search-trees/
// https://www.geeksforgeeks.org/construct-all-possible-bsts-for-keys-1-to-n/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return generateTreesHelper(1, n)
}

func generateTreesHelper(from, to int) []*TreeNode {
	var res []*TreeNode

	if from > to {
		res = append(res, nil)
		return res
	} else if from == to {
		res = append(res, &TreeNode{
			Val: from,
		})
		return res
	}

	for i := from; i <= to; i++ {
		leftTrees := generateTreesHelper(from, i-1)
		rightTrees := generateTreesHelper(i+1, to)

		for _, left := range leftTrees {
			for _, right := range rightTrees {
				res = append(res, &TreeNode{
					Val:   i,
					Left:  left,
					Right: right,
				})
			}
		}
	}

	return res
}

func main() {
	fmt.Println(generateTrees(2))
}
