package main

import (
	"fmt"
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

// #8 Create method Perimeter of type Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// Create variable with vaule for type Circle
	cir1 := Circle{
		Radius: 6,
	}
	// Call and test methods of Circle
	fmt.Println(cir1.Area())
	fmt.Println(cir1.Perimeter())
}
