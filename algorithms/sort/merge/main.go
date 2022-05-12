package main

import "fmt"

func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	mid := len(a) / 2
	return merge(mergeSort(a[:mid]), mergeSort(a[mid:]))
}

func merge(left, right []int) []int {
	k := len(left) + len(right)
	res := make([]int, k)

	l := 0
	r := 0

	for i := 0; i < k; i++ {
		if l >= len(left) {
			res[i] = right[r]
			r++
		} else if r >= len(right) {
			res[i] = left[l]
			l++
		} else if left[l] < right[r] {
			res[i] = left[l]
			l++
		} else {
			res[i] = right[r]
			r++
		}
	}

	return res
}

func main() {
	fmt.Println(mergeSort([]int{100, 20, 80, 90, 23, 500, 1000, 101, 9999}))
}
