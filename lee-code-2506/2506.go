package main

import "fmt"

func isSimilar(word1, word2 string) bool {
	m := make(map[rune]bool)
	m1 := make(map[rune]bool)
	for _, v := range word1 {
		m[v] = true
	}

	for _, v := range word2 {
		m1[v] = true
	}

	if len(m) != len(m1) {
		return false
	}

	for k, _ := range m1 {
		if !m[k] {
			return false
		}
	}
	return true
}

func similarPairs(words []string) int {
	ans := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if isSimilar(words[i], words[j]) {
				ans++
			}
		}
	}
	return ans
}

func main() {
	words := []string{"aba", "aabb", "abcd", "bac", "aabc"}
	ans := similarPairs(words)
	fmt.Println(ans)
}
