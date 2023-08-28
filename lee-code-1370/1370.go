package main

import (
	"fmt"
	"sort"
	"strings"
)

type Letter struct {
	letter string
	num    int
}

func sortString(s string) string {
	var ans strings.Builder
	m := make(map[rune]int)

	for _, letter := range s {
		m[letter] += 1
	}

	keys := make([]rune, 0, len(m))
	for k, _ := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	sum := len(s)
	for sum > 0 {
		for _, k := range keys {
			if m[k] > 0 {
				ans.WriteRune(k)
				m[k]--
				sum--
			}
		}

		for i := len(keys) - 1; i >= 0; i-- {
			if m[keys[i]] > 0 {
				ans.WriteRune(keys[i])
				m[keys[i]]--
				sum--
			}
		}
	}

	return ans.String()
}

func main() {
	str := "gtqxozxzrssrzxzoxqtg"
	//str := "abzzxx"
	ans := sortString(str)
	fmt.Println(ans)
	//goqrstxz zxtsrqog xz x
	//goqrstxz zxtsrqog xz zx
}
