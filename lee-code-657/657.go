package main

import "fmt"

func judgeCircle(moves string) bool {
	if len(moves)%2 == 1 {
		return false
	}
	pointUD := 0
	pointLR := 0
	for _, v := range moves {
		switch v {
		case 'U':
			pointUD++
		case 'D':
			pointUD--
		case 'L':
			pointLR++
		case 'R':
			pointLR--
		}
	}
	return pointUD == 0 && pointLR == 0
}

func main() {
	moves := "LDRRLRUULR"
	ans := judgeCircle(moves)
	fmt.Println(ans)
}
