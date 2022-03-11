package main

import "fmt"

// Write a function that takes in a targetSum and an array of numbers as arguments
// The function should return an array containing the shortest combination of numbers
// that add up to exactly the targetSum
// If there is a tie for the shortest combination, you may return any one of the shortest

func bestSumMemo(targetSum int, numbers []int) []int {
	return bestSumMemoHelper(targetSum, numbers, map[int][]int{})
}

func bestSumMemoHelper(targetSum int, numbers []int, cache map[int][]int) []int {
	if targetSum == 0 {
		return []int{}
	} else if targetSum < 0 {
		return nil
	}

	v, ok := cache[targetSum]
	if ok {
		return v
	}

	var shortestSum []int

	for _, num := range numbers {
		bSum := bestSumMemoHelper(targetSum-num, numbers, cache)
		if bSum != nil {
			combination := append([]int{num}, bSum...)

			if len(shortestSum) == 0 || len(shortestSum) > len(combination) {
				shortestSum = combination
			}
		}
	}

	cache[targetSum] = shortestSum
	return shortestSum
}

func bestSumTab(targetSum int, numbers []int) []int {
	data := make([][]int, targetSum+1)
	data[0] = []int{}

	for i := 0; i <= targetSum; i++ {
		if data[i] != nil {
			for _, num := range numbers {
				val := num + i
				if val <= targetSum && (len(data[val]) == 0 || len(data[val]) > (len(data[i])+1)) {
					data[val] = append([]int{num}, data[i]...)
				}
			}
		}
	}

	return data[targetSum]
}

func main() {
	fmt.Println(bestSumMemo(7, []int{2, 3}), bestSumTab(7, []int{2, 3}))
	fmt.Println(bestSumMemo(7, []int{5, 3, 4, 7}), bestSumTab(7, []int{5, 3, 4, 7}))
	fmt.Println(bestSumMemo(7, []int{2, 4}), bestSumTab(7, []int{2, 4}))
	fmt.Println(bestSumMemo(8, []int{1, 4, 5}), bestSumTab(8, []int{1, 4, 5}))
	fmt.Println(bestSumMemo(300, []int{7, 14}), bestSumTab(300, []int{7, 14}))
	fmt.Println(bestSumMemo(100, []int{1, 2, 5, 25}), bestSumTab(100, []int{1, 2, 5, 25}))
}
