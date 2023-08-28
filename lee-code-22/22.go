package main

import (
	"fmt"
)

func generateParenthesis1(n int) []string {

	result := make([][]string, 100, 100)
	result[0] = []string{"()"}

	for i := 1; i < n; i++ {
		for _, v := range result[i-1] {
			result[i] = append(result[i], v+"()", "("+v+")", "()"+v)
		}
		result[i] = removeDuplicates(result[i])
	}

	return result[n-1]
}

func removeDuplicates(input []string) []string {
	uniqueMap := make(map[string]bool)
	result := make([]string, 0)

	for _, str := range input {
		if !uniqueMap[str] {
			uniqueMap[str] = true
			result = append(result, str)
		}
	}

	return result
}

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	generate("", n, n, &result)
	return result
}

func generate(s string, left int, right int, result *[]string) {
	if left == 0 && right == 0 {
		*result = append(*result, s)
		return
	}

	if left > 0 {
		generate(s+"(", left-1, right, result)
	}

	if right > left {
		generate(s+")", left, right-1, result)
	}
}

func main() {
	ans := generateParenthesis(4)
	fmt.Println(ans)
}
