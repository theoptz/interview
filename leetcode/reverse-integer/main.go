package main

import (
	"fmt"
	"math"
)

// Follow link to see details
// https://leetcode.com/problems/reverse-integer/

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	var res int
	for x > 0 {
		res = res*10 + x%10
		if res > math.MaxInt32 {
			return 0
		}

		x /= 10
	}

	return res * sign
}

func main() {
	fmt.Println(reverse(123))  // 321
	fmt.Println(reverse(-123)) // -321
	fmt.Println(reverse(120))  // 21
}
