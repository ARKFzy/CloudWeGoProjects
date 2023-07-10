//%T打印该参数的类型，%v打印值

package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func printArea(s Shape) {
	fmt.Println("s.Area() =", s.Area())
}

func main() {
	c := Circle{radius: 1}
	printArea(c)

	var s Shape = c
	if Circle, ok := s.(Circle); ok {
		fmt.Println(Circle.radius)
	}
}
