package main

import "fmt"

func reversePrefix(word string, ch byte) string {
	stopIndex := -1
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			stopIndex = i
			break
		}
	}
	if stopIndex > 0 {

		halfWord := []byte(word[0 : stopIndex+1])

		l, r := 0, stopIndex

		for l < r {
			halfWord[l], halfWord[r] = halfWord[r], halfWord[l]
			l++
			r--
		}
		return string(halfWord) + word[stopIndex+1:]
	}
	return word
}

func main() {
	ans := reversePrefix("abcdefd", 'd')
	fmt.Println(ans == "dcbaefd")

	ans = reversePrefix("xyxzxe", 'z')
	fmt.Println(ans == "zxyxxe")

	ans = reversePrefix("abcd", 'z')
	fmt.Println(ans == "abcd")
}
