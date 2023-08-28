package main

import "fmt"

func repeatedCharacter1(s string) byte {
	m := make(map[int32]int)
	var ans byte
	for _, b := range s {
		_, ok := m[b]
		if !ok {
			m[b] = 1
		} else {
			ans = byte(b)
			break
		}
	}
	return ans
}

func repeatedCharacter(s string) byte {
	seen := 0
	for _, c := range s {
		if seen>>(c-'a')&1 > 0 {
			return byte(c)
		}
		seen |= 1 << (c - 'a')
	}
	return 0 // impossible
}

func main() {
	ans := repeatedCharacter("abccbaacz")
	fmt.Println(string(ans))
}
