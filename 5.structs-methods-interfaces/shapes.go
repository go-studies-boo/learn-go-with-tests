package structsmethodsinterfaces

import "math"

type Rectangle struct {
	Width 	float64
	Height 	float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base 	float64
	Height 	float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Circle) Area() float64 {
	return math.Pi * r.Radius * r.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}