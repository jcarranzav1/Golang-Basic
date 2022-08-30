package main // un package es el nombre de la carpeta en donde esta guardado. En este caso como es el archivo principal, será main

import "fmt"

func printMessage(message string) {
	fmt.Println(message)
}

func printAllMessages(params ...string) {
	var message string

	for _, s := range params {
		message += s + " "
	}

	fmt.Printf("Messages: %v\n", message)
}

func sum(a int, b int) int {
	return a + b
}

func sumAndRest(a, b int) (c, d int) {
	return a + b, a - b
}

func main() {
	printMessage("hola mundo")
	printAllMessages("hola", "mundo", "quiero comer", "ceviche")

	value := sum(2, 4)
	fmt.Println("Value:", value)

	_, restValue := sumAndRest(2, 4)
	fmt.Println("Value 2:", restValue)
}
