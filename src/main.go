package main

import "fmt"

// Los canales son los elementos que nos van a permitir  enviar y recibir valores entre gorutinas (funciones que se esta ejecutando de manera concurrente)
// los canales son tipados.

/*
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
Por defecto, los envíos y las recepciones se bloquean hasta que el otro lado esté listo. Esto permite que las goroutines se sincronicen sin bloqueos explícitos o variables de condición.

El código de ejemplo suma los números en una porción, distribuyendo el trabajo entre dos goroutines. Una vez que ambas goroutines han completado su cálculo, calcula el resultado final.
*/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func listen(c <-chan int) { //hagamos que la funcion listen solo escuche (recived only)
	value := <-c
	fmt.Println("listen:", value)

	// c <- 31 //sale error porque esta funcion solo puede recibir.

	fmt.Println("la funcion listen it's over")

}

func write(c chan<- int) { //hagamos que la funcion write solo envie (send only)
	c <- 7
	// j := <-c  sale error porque esta funcion solo puede enviar.
	// fmt.Println(j)
	fmt.Println("write:", 7)
	fmt.Println("la funcion write it's over")
}

func main() {

	fmt.Println("\n**************** Firmas de la funcion *****************")
	// Es necesario agregar firma en la funcion para que realizen las acciones que deben hacer. Listen solo tiene que recibir y write debe enviar.
	c1 := make(chan int)
	go listen(c1)
	write(c1)

	fmt.Println("\n**************** Second example *****************")

	s := []int{7, 2, 8, -9, 4, 0}

	c2 := make(chan int)

	go sum(s[:len(s)/2], c2)
	go sum(s[len(s)/2:], c2)

	x, y := <-c2, <-c2

	fmt.Println(x, y)

}
