package main

import (
	"fmt"
	"sort"
)

func equalFrequency(word string) bool {
	charCount := [26]int{}
	for _, c := range word {
		charCount[c-'a']++
	}
	fmt.Println(charCount)

	for i := 0; i < 26; i++ {
		if charCount[i] == 0 {
			continue
		}
		charCount[i]--
		frequency := make(map[int]bool)
		for _, f := range charCount {
			if f > 0 {
				frequency[f] = true
			}
		}
		if len(frequency) == 1 {
			return true
		}
		charCount[i]++
	}
	return false
}

func equalFrequency2(word string) bool {
	// Create a map to store the frequency of occurrence of each character.
	cnt := make(map[int]int)
	for _, char := range word {
		cnt[int(char)]++
	}

	// Create a slice to store the unique frequencies.
	v := []int{}
	for _, freq := range cnt {
		v = append(v, freq)
	}

	sort.Ints(v)

	if len(v) == 1 && (v[0] == 1 || cnt[v[0]] == 1) {
		return true
	}

	if len(v) == 2 {
		x, y := v[0], v[1]

		if x == 1 && cnt[x] == 1 {
			return true
		}

		if x+1 == y && cnt[y] == 1 {
			return true
		}
	}

	return false
}

func equalFrequency1(word string) bool {

	m := make(map[int32]int)
	for _, v := range word {
		m[v] += 1
	}

	m1 := make(map[int]int)
	sum := 0
	for _, v := range m {
		m1[v] += 1
		sum += v
	}

	if sum == len(m) {
		return true
	}

	if len(m1) == 1 || len(m1) > 2 {
		return false
	}

	k1, k2 := 0, 0
	for k, _ := range m1 {
		k1 = k
		if k < k2 {
			k2 = k
		}
	}

	if (k1 == 1 && m1[k2] == 1) || (m1[k2]-m1[k1] == 1) {
		return true
	}

	return false
}

func main() {
	fmt.Println(equalFrequency("aabbcc"))
}
