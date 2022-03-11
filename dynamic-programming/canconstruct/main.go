package main

import (
	"fmt"
	"strings"
)

// Write a function that accepts a target string and an array of strings.
// The function should return a boolean indicating whether or not the target can be constructed by
// concatenating elements of the wordBank array.
// You may reuse elements of wordBank as many times as needed.

func canConstructMemo(target string, wordBank []string) bool {
	return canConstructMemoHelper(target, wordBank, map[string]bool{})
}

func canConstructMemoHelper(target string, wordBank []string, cache map[string]bool) bool {
	if target == "" {
		return true
	}

	v, ok := cache[target]
	if ok {
		return v
	}

	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			res := canConstructMemoHelper(target[len(word):], wordBank, cache)
			if res {
				cache[target] = true
				return true
			}
		}
	}

	cache[target] = false
	return false
}

func canConstructTab(target string, wordBank []string) bool {
	totalLength := len(target)
	data := make([]bool, totalLength+1)
	data[0] = true

	for i := 0; i <= totalLength; i++ {
		if data[i] {
			for _, word := range wordBank {
				if len(word)+i <= totalLength {
					if strings.HasPrefix(target[i:], word) {
						data[i+len(word)] = true
					}
				}
			}
		}
	}

	return data[totalLength]
}

func main() {
	fmt.Println(
		canConstructMemo("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}),
		canConstructTab("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}),
	)
	fmt.Println(
		canConstructMemo("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}),
		canConstructTab("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}),
	)
	fmt.Println(
		canConstructMemo("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}),
		canConstructTab("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}),
	)
	fmt.Println(
		canConstructMemo("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee"}),
		canConstructTab("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee"}),
	)
}
