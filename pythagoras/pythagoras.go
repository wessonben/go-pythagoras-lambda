package pythagoras

import "math"

// GetHypotenuse - return triangle hypotenuse.
func GetHypotenuse(a float64, b float64) float64 {

	if a <= 0 || b <= 0 {
		return 0
	}

	return math.Sqrt((a * a) + (b * b))
}

// GetArea - return triangle area.
func GetArea(a float64, b float64) float64 {

	if a <= 0 || b <= 0 {
		return 0
	}

	rect := a * b
	area := rect / 2
	return area
}

//GetPerimeter - return triangle perimeter.
func GetPerimeter(a float64, b float64) float64 {

	if a <= 0 || b <= 0 {
		return 0
	}

	c := GetHypotenuse(a, b)
	return (a + b + c)
}
