package main

import "time"

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.requests = append(this.requests, t)

	// 移除超出时间范围的请求
	for len(this.requests) > 0 && this.requests[0] < t-3000 {
		this.requests = this.requests[1:]
	}

	// 返回队列中的请求数
	return len(this.requests)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

func main() {
	o := Constructor()

	println(o.Ping(time.Now().Second()))
	println(o.Ping(time.Now().Second()))
	println(o.Ping(time.Now().Second()))

}
