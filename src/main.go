package main

import (
	"fmt"
	"math"
	"reflect"
)

// Si los structs que tenemos en el código tienen métodos que hacen algo en común (Cálculos, obtener data, etc), es posible ejecutar éstos métodos usando una interfaz, de esta forma evitamos hacer código por cada struct.
// Esta es una forma de crear interfaces de forma implicita
type Figure interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.radius
}

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func (s Square) Perimeter() float64 {
	return 4 * s.side
}

func Calcular(f Figure) {
	fmt.Println("El área del ", reflect.TypeOf(f).Name(), "es:", f.Area())
	fmt.Println("El perímetro del ", reflect.TypeOf(f).Name(), "es:", f.Perimeter())
}

func LessArea(s1, s2 Figure) Figure {
	if s1.Area() < s2.Area() {
		return s1
	}
	return s2
}

func main() {
	mySquare := Square{side: 2}
	myCircle := Circle{radius: 4}

	Calcular(mySquare)
	Calcular(myCircle)

	minArea := LessArea(myCircle, mySquare)

	fmt.Println("The smallest area of the two figures is", reflect.TypeOf(minArea).Name())

	fmt.Println("\n*********** empty interface ****************")
	// Una interfaz vacía puede contener valores de cualquier tipo. Las interfaces vacías son utilizadas por el código que maneja valores de tipo desconocido.

	var value interface{}
	value = "hola"

	fmt.Println("Value: ", value)

	value = 12

	fmt.Println("Value: ", value)

	fmt.Println("\n*********** Lista de interfaces ****************")

	// Lista de interfaces

	/*  Usualmente en muchos lenguajes de programacion mucho mas flexibles  a una misma lista puedes agregar diferente tipo de datos, por ejemplo: [1, "bo", true]. En GO no puedes hacer eso con los slices, porque tienes que indicar que tipo de dato vas a ingresar

	Entonces, ¿como podemos simular eso en GO?

	creamos un slice, lo llamaremos myInterface, y este va a instanciar  un slice de interfaces vacias. Como sabemos las interfaces vacias aceptan valores de cualquier tipo. En otras palabras creamos un slice donde cada elemento de este es un slice vacio.

	*/

	myInterface := []interface{}{"Hola", 1, 2.53, true}

	fmt.Println(myInterface...) //con los tres punto imprimimos los elementos sin el array, parecido al spread de js

}
