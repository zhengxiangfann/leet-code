package main

func numJewelsInStones(jewels string, stones string) int {
	m := make(map[int32]bool)
	for _, v := range jewels {
		m[v] = true
	}

	ans := 0
	for _, v := range stones {
		if m[v] {
			ans++
		}
	}

	return ans

}

func main() {

}
