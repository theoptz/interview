package main

import "fmt"

// Follow link to see details
// https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	reverse := 0
	source := x

	for x > 0 {
		reverse = reverse*10 + x%10
		x /= 10
	}

	return source == reverse
}

func main() {
	fmt.Println(isPalindrome(121))  // true
	fmt.Println(isPalindrome(-121)) // false
	fmt.Println(isPalindrome(10))   // false
}
