package main

import "fmt"

type Shape interface {
	getArea() int
}

type square struct {
	width int
}

type rectangle struct {
	width  int
	height int
}

func main() {
	square := square{
		width: 20,
	}

	rectangle := rectangle{
		width:  20,
		height: 50,
	}

	logArea(square)
	logArea(rectangle)
}

func logArea(s Shape) {
	fmt.Println("Area is", s.getArea())

	if _, ok := s.(rectangle); ok {
		fmt.Println("This is a rectangle")
	} else {
		fmt.Println("This is a square")
	}

	switch s := s.(type) {
	case square:
		fmt.Println("This is a square and the area is", s.getArea())
	case rectangle:
		fmt.Println("This is a rectangle and the area is", s.getArea())
	}
}

func (r square) getArea() int {
	return r.width * r.width
}

func (r rectangle) getArea() int {
	return r.height * r.width
}
