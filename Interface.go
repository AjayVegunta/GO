package main

import (
	"fmt"
	"math"
)

type shapes interface {
	area() float64
}
type circle struct {
	radius float64
}
type rectangle struct {
	w, l float64
}

func (a circle) area() float64 {
	return math.Pi * a.radius * a.radius
}
func (b rectangle) area() float64 {
	return b.l * b.w
}
func getArea(s shapes) float64 {
	return s.area()
}
func main() {
	a := circle{6}
	b := rectangle{10, 20}
	fmt.Println("the area of circle is", getArea(a))
	fmt.Println("the area of rectangle is", getArea(b))
}
