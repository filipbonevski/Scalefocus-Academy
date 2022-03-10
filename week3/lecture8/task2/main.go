package main

import (
	"fmt"
	"math"
)

type Shapes []Shape

func (s Shapes) LargestArea() float64 {
	var maxArea float64
	maxArea = 0
	for _, shape := range s {
		if shape.Area() > maxArea {
			maxArea = shape.Area()
		}
	}
	return maxArea
}

type Square struct {
	side float64
}

func (r Square) Area() float64 {
	return r.side * r.side
}

type Circle struct {
	radius float64
}

func (r Circle) Area() float64 {
	return math.Pi * r.radius * r.radius
}

type Shape interface {
	Area() float64
}

func main() {
	var a Shape = Square{side: 6}
	var b Shape = Circle{radius: 3}
	var c Shape = Square{side: 4}
	var d Shape = Circle{radius: 1}

	var shapes Shapes
	shapes = append(shapes, a)
	shapes = append(shapes, b)
	shapes = append(shapes, c)
	shapes = append(shapes, d)

	fmt.Println(shapes.LargestArea())
}
