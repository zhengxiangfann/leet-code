package main

func maximumWealth(accounts [][]int) (ans int) {
	for _, account := range accounts {

		wealth := 0
		for _, money := range account {
			wealth += money
		}
		if ans < wealth {
			ans = wealth
		}
	}
	return
}

func main() {

}
