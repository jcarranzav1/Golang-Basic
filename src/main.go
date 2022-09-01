package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\n************ Arrays ***********\n")

	// La diferencia principal entre los arrays es que estos tienen una longitud fija e invariable y deben declarase especifiandola.

	var text [2]string
	text[0] = "Hello"
	text[1] = "World"
	fmt.Println(text[0], text[1])
	fmt.Println(text)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	fmt.Println("\n************ Slices ***********\n")

	// Un array tiene un tamaño fijo. En cambio, un slice es una vista flexible y de tamaño dinámico de los elementos de un array. En la práctica, los slices son mucho más comunes que los arrays.

	// Un corte se forma especificando dos índices, un límite inferior y otro superior, separados por dos puntos:
	// a[low high]

	var primeSlice []int = primes[1:4]
	fmt.Println(primeSlice)

	fmt.Println("\n************ Slices are like references to arrays ***********\n")

	// Un slice no almacena ningún dato, sólo describe una sección de un array subyacente.
	// Si se cambian los elementos de uN SLICE, se modifican los elementos correspondientes del array subyacente.
	// Otras slices que comparten el mismo array subyacente verán esos cambios.

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	fmt.Println("\n************ Slice literals ***********\n")
	// 		Un slice literal es lo mismo que un array pero sin agregar longitud.

	// Esto es un array literal:
	arrayLiteral := [3]bool{true, true, false}
	fmt.Println(arrayLiteral)

	// Y esto creata el mismo array de arriva y luego contruye un slice que lo referencia:

	sliceLiteral := []bool{true, true, false}
	fmt.Println(sliceLiteral)

	fmt.Println("\n************ Slice defaults ***********\n")
	// en los slices, puedes omitir los limites superiores e inferiores para utilizar sus valores predeterminados en su lugar

	s1 := []int{2, 3, 5, 7, 11, 13}

	s1 = s1[1:4]
	fmt.Println(s1)

	s1 = s1[:2]
	fmt.Println(s1)

	s1 = s1[1:]
	fmt.Println(s1)

	fmt.Println("\n************ Slice length and capacity ***********\n")

	/*
		Un slice tiene tanto una longitud como una capacidad.
		La longitud de un slice es el número de elementos que contiene.

		La capacidad de un slice es el número de elementos del array subyacente, contando desde el primer elemento de la porción.
	*/

	slice2 := []int{2, 3, 5, 7, 11, 13}
	printSlice(slice2)

	// Slice the slice to give it zero length.
	slice2 = slice2[:0]
	printSlice(slice2)

	// Extend its length.
	slice2 = slice2[:4]
	printSlice(slice2)

	// Drop its first two values.
	slice2 = slice2[2:]
	printSlice(slice2)

	fmt.Println("\n************ Nil slices ***********\n")
	// el valor zero de lo slices es nil. Un nil slice tiene un longitud y capacidad de 0, ademas que no tiene un array subyacente

	var nilSlice []int
	fmt.Println(nilSlice, len(nilSlice), cap(nilSlice))
	if nilSlice == nil {
		fmt.Println("nil!")
	}

	fmt.Println("\n************ Creating a slice with make ***********\n")
	// Los slices se pueden crear con la función incorporada make; así es como se crean array de tamaño dinámico.
	// La función make asigna un array a cero y devuelve un slice que hace referencia a ese array.
	// make([]type, len, cap)

	make1 := make([]int, 5)
	printSlice(make1)

	make2 := make([]int, 0, 5)
	printSlice(make2)

	make3 := make2[:2]
	printSlice(make3)

	make4 := make3[2:5]
	printSlice(make4)

	fmt.Println("\n************ Slice of slices ***********\n")
	// Los slices pueden contener cualquier tipo, incluso otros slices

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	fmt.Println("\n************ Appending to a slices ***********\n")
	// Es habitual añadir nuevos elementos a un slice, por lo que Go proporciona una función append incorporada.
	// func append(s []T, vs ...T) []T

	// El primer parámetro s de append es un slice de tipo T, y el resto son valores T a añadir al slice.

	// El valor resultante de append es una porción que contiene todos los elementos de la porción original más los valores proporcionados.

	var sliceAppend []int
	printSlice(sliceAppend)

	// append works on nil slices.
	sliceAppend = append(sliceAppend, 0)
	printSlice(sliceAppend)

	// The slice grows as needed.
	sliceAppend = append(sliceAppend, 1)
	printSlice(sliceAppend)

	// We can add more than one element at a time.
	sliceAppend = append(sliceAppend, 2, 3, 4)
	printSlice(sliceAppend)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
