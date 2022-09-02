package main

import "fmt"

// structs es un tipo definido por el usuario que permite agrupar/combinar elementos de tipos posiblemente diferentes en un único tipo. Es muy similar al concepto de clases en POO.
// struct es una colección de campos.

type player struct {
	x, y    int // podemos hacerlo mas compacto combinando los campos del mismo tipo
	name    string
	species string
	alive   bool
}

func main() {
	player1 := player{10, 10, "Juan", "human", true}
	fmt.Println("Player 1:", player1)

	fmt.Println("\n***************** other way *************************")
	var player2 player

	player2.name = "Diego"
	fmt.Println("Player 2:", player2)

}
