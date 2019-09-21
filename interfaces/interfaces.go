package interfaces

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width float64
	height float64
	isSquare bool
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func (rect Rectangle) Area() float64 {
	return rect.width * rect.height
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

func (circle Circle) Length() float64 {
	return 2 * math.Pi * circle.radius
}

func main() {
	rect1 := Rectangle{
		width:    10.5,
		height:   5.67,
		isSquare: false,
	}
	rect1Area := rect1.Area()
	fmt.Println("Rectangle area:",rect1Area)

	rect2 := Rectangle{
		width:    10,
		height:   10,
		isSquare: true,
	}
	rect2Area := rect2.Area()
	fmt.Println("Rectangle area:",rect2Area)

	circle1 := Circle{
		radius: 3.33,
	}
	circle1Area := circle1.Area()
	fmt.Println("Circle area:",circle1Area)

	var shape1 Shape
	shape1 = rect2
	fmt.Println(shape1)

	var shape2 Shape
	shape2 = circle1
	fmt.Println(shape2)



}
