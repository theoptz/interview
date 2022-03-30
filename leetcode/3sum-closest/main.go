package main

import (
	"fmt"
	"sort"
)

// Follow link to see details
// https://leetcode.com/problems/3sum-closest/

func checkAbs(target, closest, value int) bool {
	diff := target - closest
	if diff < 0 {
		diff *= -1
	}

	diff2 := target - value
	if diff2 < 0 {
		diff2 *= -1
	}

	return diff2 < diff
}

func threeSumClosest(nums []int, target int) int {
	closest := nums[0] + nums[1] + nums[2]
	if closest == target {
		return target
	}

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
			if sum == target {
				return target
			}

			if checkAbs(target, closest, sum) {
				closest = sum
			}

			if sum > target {
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

	return closest
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1)) // 2
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))      // 0
}
