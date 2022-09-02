package mypackage

import "fmt"

// La visibilidad de los paquetes es sumamente importante para mantener un codigo ergonomico y escalable. En GO para saber si un elemento es exported o unexported es dependiendo de como este escrito. Si empieza con mayuscula, se exportará y podras acceder o modificarlo; si es miniscula no se exportará por lo que asumira el rol de elemento no visible o privado, y te arrojara error.

type PlayerPublic struct {
	X, Y  int
	Name  string
	money int
}

type playerPrivate struct {
	x, y int
	name string
}

// PrintMessage imprimir un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}

func printMessage1(text string) {
	fmt.Println(text)
}
