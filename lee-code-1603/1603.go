package main

type ParkingSystem struct {
	bigSiteCap    int
	mediumSiteCap int
	smallSiteCap  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		bigSiteCap:    big,
		mediumSiteCap: medium,
		smallSiteCap:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.bigSiteCap > 0 {
			this.bigSiteCap--
			return true
		} else {
			return false
		}
	case 2:
		if this.mediumSiteCap > 0 {
			this.mediumSiteCap--
			return true
		} else {
			return false
		}

	case 3:
		if this.smallSiteCap > 0 {
			this.smallSiteCap--
			return true
		} else {
			return false
		}
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */

func main() {

}
