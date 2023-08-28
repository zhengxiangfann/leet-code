package main

import "fmt"

func restoreString(s string, indices []int) string {

	ans := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		ans[i] = s[indices[i]]
	}

	return string(ans)

}

func main() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	ans := restoreString(s, indices)
	fmt.Println(ans)
}
