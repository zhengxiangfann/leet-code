package main

import "fmt"

func match(digits []string) []string {
	var ans []string
	for i := 0; i < len(digits); i++ {
		for j := i + 1; j < len(digits); j++ {
			ans = append(ans, fmt.Sprintf("%s%s", digits[i], digits[j]))
		}
	}
	return ans
}

func letterCombinations1(digits string) (ans []string) {
	m := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	for _, digit := range digits {
		fmt.Println(digit)
		fmt.Println(m[string(digit)])
	}
	return
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	m := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	var generate func(idx int, current string)
	ans := make([]string, 0)

	generate = func(idx int, current string) {
		if idx == len(digits) {
			ans = append(ans, current)
			return
		}

		digit := string(digits[idx])
		letters := m[digit]

		for _, letter := range letters {
			generate(idx+1, current+letter)
		}
	}

	generate(0, "")

	return ans
}

func main() {
	fmt.Println(letterCombinations("23"))
}
