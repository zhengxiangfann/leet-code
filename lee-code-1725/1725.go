package main

func countGoodRectangles1(rectangles [][]int) (ans int) {
	m := make(map[int]int)

	for _, v := range rectangles {
		minSite := 0
		if v[0] <= v[1] {
			minSite = v[0]
		} else {
			minSite = v[1]
		}
		m[minSite] += 1
	}

	maxK := 0
	for k, v := range m {
		if k > maxK {
			maxK = k
			ans = v
		}
	}
	return

}

func countGoodRectangles(rectangles [][]int) int {
	maxLen := 0
	count := 0
	for _, rect := range rectangles {
		minSide := min(rect[0], rect[1])
		if minSide > maxLen {
			maxLen = minSide
			count = 1
		} else if minSide == maxLen {
			count++
		}
	}
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//nums := [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}} //3
	//nums := [][]int{{5, 8}, {3, 9}, {3, 12}}
	// k <= li 和 k <= wi
	nums := [][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}} //3
	ans := countGoodRectangles(nums)
	println(ans)
}
