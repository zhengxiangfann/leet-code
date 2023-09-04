package main

import (
	"log"
	"time"
)

func TickerDemo() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("Ticker tick.")
		}
	}
}

//
//func TickerLaunch() {
//	ticker := time.NewTicker(5 * time.Minute)
//	maxPassenger := 30 // 每车最大装载人数
//	passengers := make([]string, 0, maxPassenger)
//	for {
//		passenger := GetNewPassenger() // 获取一个新乘客
//		if passenger != "" {
//			passengers = append(passengers, passenger)
//		} else {
//			time.Sleep(1 * time.Second)
//		}
//		select {
//		case <-ticker.C: // 时间到，发车
//			Launch(passengers)
//			passengers = []string{}
//		default:
//			if len(passengers) >= maxPassenger { // 时间没到，车已座满，发车
//				Launch(passengers)
//				passengers = []string{}
//			}
//		}
//	}
//}

func main() {
	TickerDemo()
	//TickerLaunch()
}
