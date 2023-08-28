package main

import "fmt"

func flipAndInvertImage(image [][]int) [][]int {
	for i := 0; i < len(image); i++ {
		l, r := 0, len(image[i])-1
		for l <= r {
			image[i][l], image[i][r] = image[i][r]^1, image[i][l]^1
			l++
			r--
		}
	}
	return image
}

func main() {
	image := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	flipAndInvertImage(image)

	fmt.Println(1^1, 0^1)

	fmt.Println(1&1, 0&1)

	fmt.Println(1|1, 0|1)

}
