package main // un package es el nombre de la carpeta en donde esta guardado. En este caso como es el archivo principal, será main

import "fmt"

func main() {
	const helloMessage string = "hello"
	const worldMessage string = "world"

	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Platzi"
	cursos := 500
	// Con valores seguros
	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	// Cuando no sabemos que tipo de dato es:
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%v tiene más de %v cursos\n", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos:
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("cursos: %T\n", cursos)

}
