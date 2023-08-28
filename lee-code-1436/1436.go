package main

import "fmt"

func destCity(paths [][]string) string {
	if len(paths) < 1 {
		return ""
	}

	cityCounts := make(map[string]int, 0)

	for _, path := range paths {
		depart, destination := path[0], path[1]
		cityCounts[depart]++
		if _, ok := cityCounts[destination]; !ok {
			cityCounts[destination] = 0
		}
	}

	for depart, destination := range cityCounts {
		if destination == 0 {
			return depart
		}
	}
	return ""
}

func main() {
	//paths := [][]string{
	//	{"A", "B"},
	//	{"B", "D"},
	//	{"D", "E"},
	//}

	//paths := [][]string{
	//	{"London", "New York"},
	//	{"New York", "Lima"},
	//	{"Lima", "Sao Paulo"},
	//}
	paths := [][]string{
		{"B", "C"},
		{"D", "B"},
		{"C", "A"},
	}

	ans := destCity(paths)
	fmt.Println(ans)
}
