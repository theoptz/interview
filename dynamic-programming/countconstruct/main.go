package main

import (
	"fmt"
	"strings"
)

// Write a function that accepts a target string and an array of strings.
// The function should return the number of ways that the target can be constructed by
// concatenating elements of the wordBank array.
// You may reuse elements of wordBank as many times as needed.

func countConstructMemo(target string, wordBank []string) int {
	return countConstructMemoHelper(target, wordBank, map[string]int{})
}

func countConstructMemoHelper(target string, wordBank []string, cache map[string]int) int {
	if target == "" {
		return 1
	}

	v, ok := cache[target]
	if ok {
		return v
	}

	var count int

	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			count += countConstructMemoHelper(target[len(word):], wordBank, cache)
		}
	}

	cache[target] = count
	return count
}

func countConstructTab(target string, wordBank []string) int {
	totalLength := len(target)
	data := make([]int, totalLength+1)
	data[0] = 1

	for i := 0; i <= totalLength; i++ {
		if data[i] > 0 {
			for _, word := range wordBank {
				if len(word)+i <= totalLength && strings.HasPrefix(target[i:], word) {
					data[i+len(word)] += data[i]
				}
			}
		}
	}

	return data[totalLength]
}

func main() {
	fmt.Println(
		countConstructMemo("purple", []string{"purp", "p", "ur", "le", "purpl"}),
		countConstructTab("purple", []string{"purp", "p", "ur", "le", "purpl"}),
	)
	fmt.Println(
		countConstructMemo("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}),
		countConstructTab("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}),
	)
	fmt.Println(
		countConstructMemo("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}),
		countConstructTab("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}),
	)
	fmt.Println(
		countConstructMemo("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}),
		countConstructTab("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}),
	)
	fmt.Println(
		countConstructMemo("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee"}),
		countConstructTab("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee"}),
	)
}
