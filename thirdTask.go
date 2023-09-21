package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

type CustomShape struct {
	Name     string
	CalcArea func() float64
}

func (cs CustomShape) Area() float64 {
	return cs.CalcArea()
}

type ShapeManager struct {
	Shapes []Shape
}

func (sm ShapeManager) TotalArea() float64 {
	total := 0.0
	for _, shape := range sm.Shapes {
		total += shape.Area()
	}
	return total
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 3}
	triangle := Triangle{Base: 6, Height: 4}
	customShape := CustomShape{
		Name: "Custom",
		CalcArea: func() float64 {
			return 12.34
		},
	}

	shapeManager := ShapeManager{Shapes: []Shape{circle, rectangle, triangle, customShape}}

	for _, shape := range shapeManager.Shapes {
		fmt.Printf("Area of %T: %.2f\n", shape, shape.Area())
	}

	fmt.Printf("Total Area: %.2f\n", shapeManager.TotalArea())
}
