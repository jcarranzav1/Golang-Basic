package main // un package es el nombre de la carpeta en donde esta guardado. En este caso como es el archivo principal, será main

import (
	"fmt"
	"log"
	"strconv"
)

func isEven(number int) bool {
	return number%2 == 0
}

func logIn(user string, password string) bool {
	if user == "juandiego" && password == "Juandiego02" {
		return true
	}
	return false
}

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// With and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad, AND")
	}

	// With or
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad, OR")
	}

	// Convertir texto a número

	value, err := strconv.Atoi("53")

	//  nil es un identificador predeclarado que representa el valor cero para un puntero, un canal, una función, una interfaz, un mapa o un tipo de trozo.
	if err != nil {
		log.Fatal(err) // va a imprimir el error y despues terminara el codigo
	}

	fmt.Println("Value:", value)

	isNumberEven := isEven(3)
	isLoggin := logIn("juandiego", "Juandiego02")

	fmt.Printf("¿El numero es par? : %t\n", isNumberEven)
	fmt.Printf("¿Es correcto los datos? : %t\n", isLoggin)
}
