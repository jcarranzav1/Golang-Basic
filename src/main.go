package main

import "fmt"

/*
Por defecto, los canales no tienen búfer, lo que indica que solo aceptarán envíos (chan <-) si hay una recepción correspondiente (<- chan) que esté lista para recibir el valor enviado. Los canales almacenados permiten aceptar un número limitado de valores sin un receptor correspondiente para esos valores. Es posible crear un canal con un buff. Los canales con búfer se bloquean solo cuando el búfer está lleno. Del mismo modo, la recepción de un canal con búfer se bloquea solo cuando el búfer estará vacío.
*/
func main() {
	fmt.Println("\n**************** Buffered channel *****************")
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 si agregamos este se generara un bloqueo, porque excede la capacidad del buffer

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println("\n**************** Range and close channel *****************")

	/*
		Un emisor puede cerrar un canal para indicar que no se enviarán más valores. Los receptores pueden comprobar si un canal se ha cerrado asignando un segundo parámetro a la expresión de recepción: después de

		v, ok := <-ch

		ok es falso si no hay más valores que recibir y el canal está cerrado.

		El bucle for i := rango c recibe valores del canal repetidamente hasta que se cierra.

		Nota: Sólo el emisor debe cerrar un canal, nunca el receptor. Enviar en un canal cerrado causará un panic.
		Nota 2: Los canales no son como los archivos; normalmente no es necesario cerrarlos. El cierre sólo es necesario cuando el receptor debe saber que no hay más valores en camino, como para terminar un bucle de rango
	*/

	ch2 := make(chan int, 20)
	go fibonacci(cap(ch2), ch2)

	for i := range ch2 { //IMPRIME TODOS LOS VALORES DEL CANAL
		fmt.Println(i)
	}
	fmt.Println("\n**************** SELECT *****************")

	/*
		La sentencia select permite que una goroutine espere en múltiples operaciones de comunicación.

		Un select se bloquea hasta que uno de sus casos puede ejecutarse, entonces ejecuta ese caso. Elige uno al azar si hay varios listos.

		Ejemplo:
		// go func()... se lanza de manera concurrente:
		// la línea fmt.Println(<-c) queda en espera hasta que haya algo que leer de <-c.
		// Al ser lanzado concurrentemente, mientras dicha línea espera, se va a la ejecución de fibonacci(c, quit),
		// que guardará valores en c <- x TANTAS VECES SE PRODUZCA LA ESPERA EN LA LINEA fmt.Println(<-c).
		// Una vez sale del bucle, pasa a esperar case <-quit, que será satisfecho con quit <- 0, por lo que imprimirá "quit" y finalizará}

		// el programa.
	*/

	ch3 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch3)
		}
		quit <- 0
	}()
	fibonacci2(ch3, quit)

}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
