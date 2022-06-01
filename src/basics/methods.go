package basics

import (
	"fmt"
	"math"
)

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

// Methods are similar to fuctions. Method on a type is a way to achieve behaviour similar to classes.
// Methods allow a logical grouping of behaviour related to a type similar to classes.
// Methods with the same name can be defined on different types whereas functions with the same names are not allowed.

func (r Rectangle) area() int {
	return r.length * r.width
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// Method with pointer receiver.
// Changes made inside a method with a pointer receiver is visible to the caller whereas this is not the case in value receiver.
func (c *Circle) changeRadius(newRadius float64) {
	c.radius = newRadius
}

func MethodExp() {
	rect := Rectangle{length: 2, width: 4}
	fmt.Println("area of rectangle:", rect.area())

	circle := Circle{radius: 2}
	fmt.Println("area of circle:", circle.area())

	circle.changeRadius(4)
	fmt.Println("area of circle:", circle.area())
}
