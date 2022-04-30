package main

import "fmt"

// Follow link to see details
// https://leetcode.com/problems/unique-binary-search-trees/
// https://www.geeksforgeeks.org/construct-all-possible-bsts-for-keys-1-to-n/

func numTrees(n int) int {
	if n <= 1 {
		return 1
	}

	data := make([]int, n+1)
	data[0] = 1
	data[1] = 1

	for i := 2; i <= n; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			sum += data[j] * data[i-j-1]
		}

		data[i] = sum
	}

	return data[n]
}

func main() {
	fmt.Println(numTrees(3)) // 5
	fmt.Println(numTrees(1)) // 1
}
