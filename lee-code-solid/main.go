package main

import "fmt"

/*
里式替换原则
所有引用基类（父类）的地方，必须能够透明地使用其子类的对象

就是父类能做的事情，子类也应该能做，而且要做得更好、更灵活。
如果子类不能完全替换父类，那么在代码的使用阶段就会出现问题。
*/

type (
	Shape interface {
		Area() float64
	}

	Rectangle struct {
		width, height float64
	}

	Circle struct {
		radius float64
	}
)

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * 3.14
}

func main() {
	/*
		我们将 Rectangle 和 Circle 结构体都实现为了 Shape 接口的一种，
		同时也不再需要手动判断形状类型来调用对应的计算面积的方法。
		这也遵守了里氏替换原则
	*/

	s := []Shape{
		Rectangle{100, 200},
		Circle{100},
	}
	for i, v := range s {
		fmt.Println("面积：", i, v.Area())
	}
}
