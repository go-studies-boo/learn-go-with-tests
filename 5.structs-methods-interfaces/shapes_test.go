package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// his is quite different to interfaces in most other programming languages.
// Normally you have to write code to say My type Foo implements interface Bar.
//
// In Go interface resolution is implicit. 
// If the type you pass in matches what the interface is asking for, it will compile.
type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {
	// anonymous struct
	// Kent beck: The test speaks to us more clearly, as if it were an assertion of truth, not a sequence of operations
	areaTests := []struct{
		name string
		shape Shape
		hasArea float64
	} {
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.192653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	// Table driven tests
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.name, got, tt.hasArea)
			}
		})
	}
}