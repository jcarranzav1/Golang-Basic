package pcpackage

import "fmt"

type PC struct {
	ram, disk int
	brand     string
}

func (myPC PC) PrintSpecs() {
	fmt.Printf("La marca de la PC es %s\n", myPC.brand)
	fmt.Printf("La cantidad de memoria ram es de:  %d\n", myPC.ram)
	fmt.Printf("La cantidad de disco es:  %d\n", myPC.disk)
}

func (myPC *PC) DuplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func (myPC *PC) SetSpecs(ram int, disk int, brand string) {
	myPC.ram = ram
	myPC.disk = ram
	myPC.brand = brand
}

func (myPC PC) GetRam() int {
	return myPC.ram
}
