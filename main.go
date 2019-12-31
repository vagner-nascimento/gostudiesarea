package gostudiesarea

import "math"

// Pi is a known math const
const Pi = 3.1416

// Circ calculates the circumference of a circle
func Circ(rain float64) float64 {
	return math.Pow(rain, 2) * Pi
}

// Rect calculate the area of a rectangle
func Rect(base, height float64) float64 {
	return base * height
}

func _EqTriangle(base, height float64) float64 { // This function is not accessible out of this package
	return (base * height) / 2
}
