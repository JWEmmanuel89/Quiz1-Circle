package main

import (
	"math"
)

// #6. Create structure type Circle
type Circle struct {
	Radius float64
}

// #7. Create method Area of type Circle
func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}
