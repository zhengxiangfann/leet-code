package main

import "fmt"

func squareIsWhite(coordinates string) bool {

	return (coordinates[0]+coordinates[1])%2 == 1

}

func prefixCount(words []string, pref string) (ans int) {

	for _, word := range words {
		if word[:len(pref)] == pref {
			ans++
		}
	}

	return ans
}

func commonFactors1(a int, b int) (ans int) {
	minScore := 0
	if a > b {
		minScore = b
	} else {
		minScore = a
	}

	for i := 1; i <= minScore; i++ {

		if a%i == 0 && b%i == 0 {
			ans++
		}
	}

	return ans
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func countDivisors(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}

func commonFactors(a int, b int) (ans int) {
	// Find the GCD of a and b using the Euclidean algorithm
	gcdValue := gcd(a, b)
	fmt.Println(gcdValue)
	// Count the number of divisors of the GCD, which are common factors of a and b
	ans = countDivisors(gcdValue)
	fmt.Println(ans)
	return ans
}

func main() {
	//
	//grid := "b1"
	//
	//ans := squareIsWhite(grid)
	//fmt.Println(ans)

	ans := commonFactors(20, 25)
	fmt.Println(ans)

}
