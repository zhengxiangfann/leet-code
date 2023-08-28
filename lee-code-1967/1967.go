package main

import "strings"

func numOfStrings(patterns []string, word string) (ans int) {
	for _, v := range patterns {
		if strings.Contains(word, v) {
			ans++
		}
	}
	return
}

func main() {

}
