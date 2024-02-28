package main

import (
	"fmt"
	"web20240228/calculation"
	"web20240228/circle"
	interfaces "web20240228/interfaces"
	"web20240228/rectangle"
)

func main() {
	rectangle := rectangle.Rectangle{Width: 5, Height: 5}
	circle := circle.Circle{Radius: 7}
	totalArea := calculation.TotalArea([]interfaces.Shape{&rectangle, &circle})
	fmt.Printf("%.2f\n", totalArea)
	fmt.Printf("%.2g\n", totalArea)
}
