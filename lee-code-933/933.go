package main

type RecentCounter struct {
	Buf []int
	Len int
}

func Constructor() RecentCounter {
	return RecentCounter{
		Buf: make([]int, 0),
		Len: 0,
	}
}

func (this *RecentCounter) Ping(t int) int {

	if this.Len > 0 && t-this.Buf[this.Len-1] <= 3000 {

		index := 0
		for i := 0; i < this.Len; i++ {
			index = i
			if t-this.Buf[i] <= 3000 {

				goto StopForLoop
			}
		}

	StopForLoop:

		this.Buf = this.Buf[index:]
		this.Buf = append(this.Buf, t)
		this.Len = len(this.Buf)

		return len(this.Buf)
	} else {

		this.Buf = make([]int, 0)
		this.Len = 0
		this.Buf = append(this.Buf, t)
		this.Len++

		return this.Len
	}
}

func main() {
	obj := Constructor()
	nums := []int{1, 100, 3001, 3002}

	for _, v := range nums {
		obj.Ping(v)
	}

}
