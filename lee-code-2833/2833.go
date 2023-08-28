package main

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func furthestDistanceFromOrigin(moves string) (ans int) {
	underlineCount := 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'L':
			ans--
		case 'R':
			ans++
		case '_':
			underlineCount++
		}
	}
	return abs(ans) + underlineCount
}

func main() {
	furthestDistanceFromOrigin("_R__LL_")
}
