package main

import "fmt"

// write a function that takes in a targetSum and an array of numbers as arguments
// the function should return a boolean indicating whether or not it is possible to generate the targetSum
// using numbers from the array
// You may use an element of the array as many times as needed.
// You may assume that all input numbers are non-negative

func canSumMemo(targetSum int, numbers []int) bool {
	return canSumMemoHelper(targetSum, numbers, map[int]bool{})
}

func canSumMemoHelper(targetSum int, numbers []int, cache map[int]bool) bool {
	if targetSum == 0 {
		return true
	} else if targetSum < 0 {
		return false
	}

	for _, num := range numbers {
		v, ok := cache[targetSum-num]
		if ok {
			if v {
				return true
			}

			continue
		}

		canSumRes := canSumMemoHelper(targetSum-num, numbers, cache)
		if canSumRes {
			cache[targetSum] = canSumRes
			return true
		}
	}

	cache[targetSum] = false
	return false
}

func canSumTab(targetSum int, numbers []int) bool {
	data := make([]bool, targetSum+1)
	data[0] = true

	for i := 0; i <= targetSum; i++ {
		if data[i] {
			for _, num := range numbers {
				val := i + num
				if val <= targetSum {
					data[val] = true
				}
			}
		}
	}

	return data[targetSum]
}

func main() {
	fmt.Println(canSumMemo(7, []int{5, 3, 4, 7}), canSumTab(7, []int{5, 3, 4, 7}))
	fmt.Println(canSumMemo(7, []int{2, 4}), canSumTab(7, []int{2, 4}))
	fmt.Println(canSumMemo(8, []int{2, 3, 5}), canSumTab(8, []int{2, 3, 5}))
	fmt.Println(canSumMemo(300, []int{7, 14}), canSumTab(300, []int{7, 14}))
}
