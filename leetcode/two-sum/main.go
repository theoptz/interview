package main

import "fmt"

// Follow link to see details
// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	arrLength := len(nums)
	cache := make(map[int]int)

	for i := 0; i < arrLength; i++ {
		val := target - nums[i]
		if index, ok := cache[val]; ok && index != i {
			return []int{index, i}
		}

		cache[nums[i]] = i
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0, 1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1, 2]
	fmt.Println(twoSum([]int{3, 3}, 6))         // [0, 1]
}
