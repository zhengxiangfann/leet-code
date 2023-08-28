package main

import "fmt"

func maxDepth(s string) (ans int) {
	l := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			l++
			if l > ans {
				ans = l
			}
		} else if string(s[i]) == ")" {
			l--
		}
	}
	return
}

func main() {
	s := "(1+(2*3)+((8)/4))+1"
	ans := maxDepth(s)
	fmt.Println(ans)
}
