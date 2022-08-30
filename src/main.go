package main // un package es el nombre de la carpeta en donde esta guardado. En este caso como es el archivo principal, será main

import (
	"fmt"
	"math"
)

func getPerimeter(sides uint, length uint) uint {
	return sides * length
}

func getArea(sides uint, length uint) float64 {
	//So, angle t = 180/n
	// Now, tan(t) = a/2*h
	// So, h = a/(2*tan(t))
	// Area of each triangle = (base * height)/2 = a * a/(4*tan(t))
	// So, area of the polygon,

	area := float64((length * length * sides)) / (4 * math.Tan(math.Pi/float64(sides)))
	return area
}

func main() {

	const sides uint = 10
	const length uint = 9

	perimeter := getPerimeter(sides, length)
	area := getArea(sides, length)

	fmt.Printf("The perimeter is: %d and area is: %f", perimeter, area)

}
