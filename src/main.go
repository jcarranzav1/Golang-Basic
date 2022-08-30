package main // un package es el nombre de la carpeta en donde esta guardado. En este caso como es el archivo principal, será main

import "fmt"

func main() {
	fmt.Println("-------------For condicional-------------")
	// For condicional
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("-------------For while-------------")
	// For while
	counter := 0
	for counter <= 5 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("-------------For forever with break-------------")
	//For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever > 12 {
			break
		}
	}

	fmt.Println("-------------Continue-------------")
	//exit loop
	sum := 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		sum += i
	}
	fmt.Println(sum)

	fmt.Println("-------------For Range-------------")
	//For-each range loop
	arreglo := [6]int{0, 1, 4, 6, 10, 9}
	fmt.Println("Arreglo:", arreglo)

	fmt.Println("Primer ejemplo")
	for i, j := range arreglo {
		fmt.Printf("indice i: %d tiene como valor #%d\n", i, j)
	}

	fmt.Println("Segundo ejemplo")
	for i := range arreglo {
		fmt.Printf("Valor de i: %d\n", i)
	}

	fmt.Println("Tercer ejemplo")
	for _, j := range arreglo {
		fmt.Printf("Valor de i: %d\n", j)
	}

}
