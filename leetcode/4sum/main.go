package main

import (
	"fmt"
	"sort"
)

// Follow link to see details
// https://leetcode.com/problems/4sum/

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := len(nums) - 1

			for left < right {
				valLeft := nums[left]
				valRight := nums[right]

				sum := nums[i] + nums[j] + valLeft + valRight
				if sum == target {
					res = append(res, []int{nums[i], nums[j], valLeft, valRight})
					for left < right && nums[left] == valLeft {
						left++
					}
				} else if sum > target {
					for left < right && nums[right] == valRight {
						right--
					}
				} else {
					for left < right && nums[left] == valLeft {
						left++
					}
				}
			}
		}
	}

	return res
}

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0)) // [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))      // [[2,2,2,2]]
}
