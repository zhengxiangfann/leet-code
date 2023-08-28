package main

func countPoints1(rings string) int {
	pole := make(map[uint8]map[uint8]bool)
	for i := 0; i < len(rings); i += 2 {
		if _, ok := pole[rings[i+1]]; !ok {
			pole[rings[i+1]] = make(map[uint8]bool)
		}
		pole[rings[i+1]][rings[i]] = true
	}
	ans := 0
	for _, v := range pole {
		if len(v) == 3 {
			ans++
		}
	}
	return ans
}

type RGB struct {
	R, G, B bool
}

func countPoints(rings string) int {
	pole := make(map[uint8]*RGB)
	ans := 0

	for i := 0; i < len(rings); i += 2 {
		color := rings[i]
		radius := rings[i+1]

		if _, ok := pole[radius]; !ok {
			pole[radius] = &RGB{}
		}

		switch color {
		case 'R':
			pole[radius].R = true
		case 'G':
			pole[radius].G = true
		case 'B':
			pole[radius].B = true
		}
	}

	for _, v := range pole {
		if v.R && v.G && v.B {
			ans++
		}
	}
	return ans
}

func main() {
	countPoints("B0B6G0R6R0R6G9")
}
