package main

import "fmt"

// write a function that takes in a targetSum and an array of numbers as arguments.
// The function should return an array containing any combination of elements that add up to exactly the targetSum.
// If there is no combination that adds up to the targetSum, then return nil.
// If there are multiple combinations possible, you may return any single one.

func howSumMemo(targetSum int, numbers []int) []int {
	return howSumMemoHelper(targetSum, numbers, map[int][]int{})
}

func howSumMemoHelper(targetSum int, numbers []int, cache map[int][]int) []int {
	if targetSum == 0 {
		return []int{}
	} else if targetSum < 0 {
		return nil
	}

	v, ok := cache[targetSum]
	if ok {
		return v
	}

	for _, num := range numbers {
		val := howSumMemoHelper(targetSum-num, numbers, cache)
		if val != nil {
			res := append([]int{num}, val...)
			cache[targetSum] = res
			return res
		}
	}

	cache[targetSum] = nil
	return nil
}

func howSumTab(targetSum int, numbers []int) []int {
	data := make([][]int, targetSum+1)
	data[0] = []int{}

	for i := 0; i <= targetSum; i++ {
		if data[i] != nil {
			for _, num := range numbers {
				val := num + i
				if val <= targetSum {
					data[val] = append([]int{num}, data[i]...)
				}
			}
		}
	}

	return data[targetSum]
}

func main() {
	fmt.Println(howSumMemo(7, []int{2, 3}), howSumTab(7, []int{2, 3}))
	fmt.Println(howSumMemo(7, []int{5, 3, 4, 7}), howSumTab(7, []int{5, 3, 4, 7}))
	fmt.Println(howSumMemo(7, []int{2, 4}), howSumTab(7, []int{2, 4}))
	fmt.Println(howSumMemo(8, []int{2, 3, 5}), howSumTab(8, []int{2, 3, 5}))
	fmt.Println(howSumMemo(300, []int{7, 14}), howSumTab(300, []int{7, 14}))
}
