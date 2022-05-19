package main

import "fmt"

// Follow link to see details
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	index := -1

	for left <= right {
		middle := left + (right-left)/2
		if nums[middle] == target {
			index = middle
			break
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	if index == -1 {
		return []int{-1, -1}
	}

	left = index
	for i := index - 1; i >= 0; i-- {
		if nums[i] < target {
			break
		}

		left = i
	}

	right = index
	for i := index + 1; i < len(nums); i++ {
		if nums[i] > target {
			break
		}

		right = i
	}

	return []int{left, right}
}

func main() {
	res := searchRange([]int{5, 7, 7, 8, 8, 10}, 6) // [-1, -1]
	res2 := searchRange([]int{1, 4}, 4)             // [1, 1]
	fmt.Println(res, res2)
}
