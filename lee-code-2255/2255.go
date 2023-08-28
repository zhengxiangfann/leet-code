package main

func countPrefixes(words []string, s string) (ans int) {
	for i := 0; i < len(words); i++ {
		if len(words[i]) <= len(s) && words[i] == s[:len(words[i])] {
			ans++
		}
	}
	return ans
}

func main() {
	words := []string{"a", "b", "c", "ab", "bc", "abc"}
	countPrefixes(words, "abc")
}
