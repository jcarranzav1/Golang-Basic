package main

import (
	"fmt"

	pk "JD/Golang/Curso_Basico/src/mypackage"
)

// structs es un tipo definido por el usuario que permite agrupar/combinar elementos de tipos posiblemente diferentes en un único tipo. Es muy similar al concepto de clases en POO.
// struct es una colección de campos.

type Player struct {
	x, y    int // podemos hacerlo mas compacto combinando los campos del mismo tipo
	name    string
	species string
	alive   bool
}

func main() {
	var player1 pk.PlayerPublic
	// var player2 pk.playerPrivate
	player1.X = 20
	player1.Name = "juan"

	// player1.money = 100

	fmt.Println(player1)

	pk.PrintMessage("Hola Platzi")

	// pk.printMessage1("Hola")
}
