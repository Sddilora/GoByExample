package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func school(s shape) {
	fmt.Println("Şeklin alanı:", s.area())
}

func Demo3() {

	r := Rectangle{8, 9}
	school(r)

	c := Circle{5}
	school(c)
}
