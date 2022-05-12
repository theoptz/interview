package main

import "fmt"

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	low, hi := 0, len(a)-1
	pivot := a[len(a)-1]

	for i := 0; i < hi; i++ {
		if a[i] < pivot {
			a[i], a[low] = a[low], a[i]
			low++
		}
	}

	a[low], a[hi] = a[hi], a[low]

	quickSort(a[:low])
	quickSort(a[low:])

	return a
}

func main() {
	fmt.Println(quickSort([]int{100, 20, 80, 90, 23, 500, 1000, 101, 9999}))
}
