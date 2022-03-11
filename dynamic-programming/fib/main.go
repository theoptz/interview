package main

import "fmt"

// Write a function that takes in a number as an argument.
// The function should return the n-th number of the Fibonacci sequence.

// The 0th number of the sequence is 0.
// The 1st number of the sequence is 1.

// To generate the next number of the sequence, we sum the previous two.

func fibMemoHelper(n int, m map[int]int) int {
	if n <= 2 {
		return 1
	}

	v, ok := m[n]
	if ok {
		return v
	}

	res := fibMemoHelper(n-2, m) + fibMemoHelper(n-1, m)
	m[n] = res

	return res
}

func fibMemo(n int) int {
	return fibMemoHelper(n, map[int]int{})
}

func fibTab(n int) int {
	numbers := make([]int, n+1)
	if n > 1 {
		numbers[1] = 1
	}

	for i := 2; i <= n; i++ {
		numbers[i] = numbers[i-1] + numbers[i-2]
	}

	return numbers[n]
}

func main() {
	fmt.Println(fibMemo(6), fibTab(6))
	fmt.Println(fibMemo(7), fibTab(7))
	fmt.Println(fibMemo(8), fibTab(8))
	fmt.Println(fibMemo(50), fibTab(50))
}
