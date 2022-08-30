package main // un package es el nombre de la carpeta en donde esta guardado. En este caso como es el archivo principal, será main

import "fmt"

/*
Declaremos la funcion main, la funcion principal que se ejecutará en el codigo
*/

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// Declaración de variables enteras

	base := 12 // con los : hacemos que una variable que no ha sido declarada, declare y le asigne el valor
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	// son los valores por defecto que tiene un tipo de  variable cuando no le asignas un valor.
	// Recuerda cuando no defines el valor, no indica que es null, no lo olvides !
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Printf("El area del cuadrado es %d y el tipo es %T \n ", areaCuadrado, areaCuadrado)

	var (
		i int
		h bool
		s string
	)

	i = 1
	h = true
	s = "un texto cualquiera"
	// los strings solo se pueden declarar con comillas dobles, no con las simples

	fmt.Println(i)
	fmt.Println(h)
	fmt.Println(s)

	x1, y2, z3 := 1, 2, 3
	fmt.Println(x1)
	fmt.Println(y2)
	fmt.Println(z3)
}
