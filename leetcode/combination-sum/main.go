package main

import (
	"fmt"
	"sort"
)

// Follow link to see details
// https://leetcode.com/problems/combination-sum/

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)

	sort.Ints(candidates)
	combinationSumHelper(&res, candidates, []int{}, target, 0)

	return res
}

func combinationSumHelper(res *[][]int, candidates []int, current []int, target, minVal int) {
	for _, num := range candidates {
		if num > target || (minVal > 0 && num < minVal) {
			continue
		}

		combination := make([]int, 0, len(current)+1)
		combination = append(combination, current...)
		combination = append(combination, num)

		if num == target {
			*res = append(*res, combination)
			continue
		}

		combinationSumHelper(res, candidates, combination, target-num, num)
	}
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7)) // [[2, 2, 3], [7]]
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))    // [[2, 2, 2, 2], [2, 3, 3], [3, 5]]
	fmt.Println(combinationSum([]int{2}, 1))          // []
}
