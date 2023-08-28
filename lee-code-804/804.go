package main

import (
	"fmt"
	"strings"
)

func uniqueMorseRepresentations1(words []string) int {
	msCode := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]bool)
	for _, word := range words {
		morseCode := ""
		for _, w := range word {
			morseCode = morseCode + msCode[w-'a']
		}
		m[morseCode] = true
	}
	return len(m)
}

func uniqueMorseRepresentations(words []string) int {
	msCode := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]struct{})
	for _, word := range words {
		var morseCode strings.Builder
		for _, w := range word {
			morseCode.WriteString(msCode[w-'a'])
		}
		m[morseCode.String()] = struct{}{}
	}
	return len(m)
}

func main() {

	s := []string{"gin", "zen", "gig", "msg"}
	ans := uniqueMorseRepresentations(s)
	fmt.Println(ans)
}
