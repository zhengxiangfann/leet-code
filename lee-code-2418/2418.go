package main

import (
	"fmt"
	"sort"
)

func sortPeople1(names []string, heights []int) []string {
	m := make(map[int]string)
	for i := 0; i < len(names); i++ {
		m[heights[i]] = names[i]
	}

	sort.Ints(heights)

	ans := make([]string, len(heights))
	for i := len(heights) - 1; i >= 0; i-- {
		ans[len(heights)-1-i] = m[heights[i]]
	}
	return ans
}

func sortPeople(names []string, heights []int) []string {
	data := make([]struct {
		name   string
		height int
	}, len(names))

	for i := range names {
		data[i] = struct {
			name   string
			height int
		}{names[i], heights[i]}
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].height > data[j].height
	})

	sortedNames := make([]string, len(data))
	for i, entry := range data {
		sortedNames[i] = entry.name
	}

	return sortedNames
}

func main() {
	names := []string{"Mary", "John", "Emma"}
	heights := []int{180, 165, 170}

	ans := sortPeople(names, heights)
	fmt.Println(ans)
}
