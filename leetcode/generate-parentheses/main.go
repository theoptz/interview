package main

import "fmt"

// Follow link to see details
// https://leetcode.com/problems/generate-parentheses/

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	} else if n == 1 {
		return []string{"()"}
	}

	var res []string
	data := []byte{'('}
	doGenerateParenthesis(data, n-1, n, &res)
	return res
}

func doGenerateParenthesis(data []byte, leftRemain, rightRemain int, res *[]string) {
	if leftRemain == 0 && rightRemain == 0 {
		*res = append(*res, string(data))
		return
	}

	if leftRemain > 0 {
		doGenerateParenthesis(append(data, '('), leftRemain-1, rightRemain, res)
	}

	if rightRemain > leftRemain {
		doGenerateParenthesis(append(data, ')'), leftRemain, rightRemain-1, res)
	}
}

func main() {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
}
