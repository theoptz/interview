package main

import (
	"fmt"
	"strings"
)

// Write a function that accepts a target string and an array of strings.
// The function should return a 2D array containing all of the ways that the target can be constructed by concatenating
// elements of the wordBank array. Each element of the 2D array should represent one combination
// that constructs the target.
// You may reuse elements of wordBank as many times as needed.

func allConstructMemo(targetWord string, wordBank []string) [][]string {
	return allConstructMemoHelper(targetWord, wordBank, map[string][][]string{})
}

func allConstructMemoHelper(targetWord string, wordBank []string, cache map[string][][]string) [][]string {
	if targetWord == "" {
		return [][]string{
			{},
		}
	}

	v, ok := cache[targetWord]
	if ok {
		return v
	}

	var res [][]string

	for _, word := range wordBank {
		if strings.HasPrefix(targetWord, word) {
			combinations := allConstructMemoHelper(targetWord[len(word):], wordBank, cache)
			for _, combination := range combinations {
				comb := append([]string{word}, combination...)
				res = append(res, comb)
			}
		}
	}

	cache[targetWord] = res
	return res
}

func allConstructTab(targetWord string, wordBank []string) [][]string {
	totalLength := len(targetWord)
	data := make([][][]string, totalLength+1)
	data[0] = make([][]string, 1)

	for i := 0; i <= totalLength; i++ {
		if data[i] != nil {
			for _, word := range wordBank {
				if len(word)+i <= totalLength && strings.HasPrefix(targetWord[i:], word) {
					for _, comb := range data[i] {
						items := make([]string, 0, len(comb)+1)
						items = append(items, comb...)
						items = append(items, word)
						data[len(word)+i] = append(data[len(word)+i], items)
					}
				}
			}
		}
	}

	return data[totalLength]
}

func main() {
	fmt.Println(allConstructMemo("purple", []string{"purp", "p", "ur", "le", "purpl"}))
	fmt.Println(allConstructTab("purple", []string{"purp", "p", "ur", "le", "purpl"}))

	fmt.Println(allConstructMemo("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	fmt.Println(allConstructTab("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))

	fmt.Println(allConstructMemo("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))
	fmt.Println(allConstructTab("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))

	fmt.Println(allConstructMemo("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))
	fmt.Println(allConstructTab("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))

	fmt.Println(
		allConstructMemo("eeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee"}),
	)
	fmt.Println(
		allConstructTab("eeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee"}),
	)
}
