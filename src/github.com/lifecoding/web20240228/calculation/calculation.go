package calculation

import interfaces "web20240228/interfaces"

func TotalArea(shapes []interfaces.Shape) float64 {
	totalArea := 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}
