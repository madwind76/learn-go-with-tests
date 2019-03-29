package shapes

import "math"

// Rectangle ...
type Rectangle struct {
	Width  float64
	Height float64
}

// Area ...
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle ...
type Circle struct {
	Radius float64
}

// Area ...
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle ...
type Triangle struct {
	Base   float64
	Height float64
}

// Area ...
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

// Shape ...
type Shape interface {
	Area() float64
}

// Perimter ...
func Perimter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area ...
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
