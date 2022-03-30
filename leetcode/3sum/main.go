package main

import (
	"fmt"
	"sort"
)

// Follow link to see details
// https://leetcode.com/problems/3sum/

func threeSum(nums []int) [][]int {
	var res = make([][]int, 0)

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			valLeft := nums[left]
			valRight := nums[right]

			sum := valLeft + valRight + nums[i]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for nums[left] == valLeft && left < right {
					left++
				}

				for nums[right] == valRight && left < right {
					right--
				}
			} else if sum > 0 {
				for nums[right] == valRight && left < right {
					right--
				}
			} else {
				for nums[left] == valLeft && left < right {
					left++
				}
			}
		}
	}

	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))                             // [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6})) // [[-4 -2 6],[-4 0 4],[-4 1 3],[-4 2 2],[-2 -2 4],[-2 0 2]]
	fmt.Println(threeSum([]int{}))                                                // []
	fmt.Println(threeSum([]int{0}))                                               // []
}
