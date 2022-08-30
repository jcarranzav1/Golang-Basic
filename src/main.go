package main // un package es el nombre de la carpeta en donde esta guardado. En este caso como es el archivo principal, será main

import (
	"fmt"
)

// Una instrucción defer añade la invocación de la función después de la palabra clave defer en una pila. Todas las invocaciones de la pila en cuestión se invocan cuando la funcion donde se añadieron termina. Debido a que las invocaciones se disponen en una pila, se llaman en el orden “último en entrar” y “primero en salir” (LIFO).

func main() {
	defer fmt.Println("Hola") // se invocara al final de la funcion main
	defer fmt.Println("entre segundo")
	fmt.Println("Mundo")

	//otro ejemplo
	fmt.Println("Inicio")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Fin")

	// Continue y break
	for i := 0; i < 10; i++ {

		// Continue
		if i == 2 {
			continue
		}

		// Break
		if i == 8 {
			fmt.Println("Break")
			break
		}

		fmt.Println(i)
	}
}
