package main

import (
	pkPC "JD/Golang/Curso_Basico/src/pcpackage"
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() { //creamos el metodo ping para el struct pc
	fmt.Println(myPC.brand, "Pong")

}

func (myPC *pc) duplicateRAM() { //hacer esto me recuerda mucho cuando crear un metodo donde modifique el valor de un propertu usando this
	myPC.ram = myPC.ram * 2
}

func main() {

	// con & accedemos a la posicion de memoria de una variable, mientra que con * accedemos al valor de la posicion de memoria.
	// Por eso con & se usa para crear punteros de una variable, y con * es para entregar el valor que tiene ese puntero.

	a := 50
	b := &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	fmt.Println("\n**************** pointer with structs ***********************")

	pc1 := pc{ram: 16, disk: 1000, brand: "Asus"}

	pc1.ping()

	fmt.Println(pc1)
	pc1.duplicateRAM()

	fmt.Println(pc1)
	pc1.duplicateRAM()

	fmt.Println(pc1)

	fmt.Println("\n**************** EXERCISE ***********************")

	var pc2 pkPC.PC

	pc2.SetSpecs(64, 1000, "Asus")
	pc2.PrintSpecs()

	pc2.DuplicateRAM()

	fmt.Println("RAM: ", pc2.GetRam())

}
