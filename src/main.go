package main

import (
	"fmt"
	"sync"
)

/*
	En los programas hasta ahora creados, las líneas aparecen de forma secuencial. Una línea no se ejecuta hasta que la anterior se complete, lo que es un comportamiento deseado en muchos casos.

	Sin embargo también podemos toparnos con casos en los cuales una actividad no bloquea otra, al menos no si el tiempo que tomará en completarse es suficiente como para generar un bloqueo significativo. Algunas de estas acciones pueden ser tareas como una solicitud de lectura desde la red o la lectura de datos desde un archivo.

	Go utiliza la concurrencia, que consiste en separar múltiples tareas de forma individual sin que estas bloqueen unas a las otras. Otra aproximación de otros lenguajes es el paralelismo.

	La concurrencia es el poder ejecutar diferentes partes de un algoritmo, programa, solucion o tarea de forma independiente una de la otra, para que sin importar el orden en el que ejecutemos estas partes el resultado siempre sea igual.


	Segun  Rob Pike creador de GO:
	La concurrencia consiste en tratar muchas cosas a la vezen cambio el paralelismo consiste en hacer muchas cosas a la vez. No es lo mismo, pero está relacionado.

	La concurrencia tiene que ver con la estructura, el paralelismo con la ejecución.

	La meta de la concurrencia no es el paralelismo, si no es tener una buena estructura.

	La concurrencia proporciona una forma de estructurar una solución para resolver un problema que puede (pero no necesariamente) ser paralelizable.

	La concurrencia es una forma de estructurar un programa dividiéndolo en piezas que pueden ejecutarse de forma independiente. Este es el modelo de GO.

	revisar: https://go.dev/talks/2012/waza.slide#28

*/

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

// la funcion main esta corriendo dentro de una go rutine, que una ves que termina de ejecutarse muere.

func main() {

	var wg sync.WaitGroup
	fmt.Println("Hello")
	wg.Add(1)

	go say("World", &wg) // agregando go, crearas una gorutine el cual se ejecutará de forma concurrente
	wg.Wait()

	wg.Add(1)

	go func(text string, wg *sync.WaitGroup) { // funcion anonima
		defer wg.Done()
		fmt.Println(text)
	}("Bye", &wg)

	wg.Wait()
}

/*
	func main() {
		say("Hello")
		go say("World")


	}

	cuando ejecutamos lo de arriba, veremos que World no se imprimirá, y esto es porque main se corre dentro una gorutina y una ves termine muere. Pues main termina a Hello porque la World se ejecuta de forma concurrente y no queda en el mismo hilo de ejecucion con main.

	¿Como lo solucionamos?

	Una accion basica, pero recomendable, es agregando un sleep.

	func main() {
		say("Hello")
		go say("World")
		time.Sleep(time.Second * 1)

	}

	Pero esta forma no es optima, porque estamos agregaando tiempo muerto

	Hay una forma algo avanzado, usando wait group

	Un WaitGroup espera a que una colección de goroutines termine su ejecución.
	Para esto se una la WaitGroup.Add() ( wg.add(1) en el ejemplo de la clase).
	El número entero indica el número de goroutines que debe esperar para finalizar la ejecución de la goroutine principal.

	Cada vez que una goroutine termina su ejecución, llama el método Done(). Esto hace que el contador del WaitGroup se reduzca.
	Cuando el contador llegue a zero la rutina principal continuará su ejecución.

	vLa función wait() bloquea la rutina principal hasta que todas las demás rutinas del grupo hayan terminado.


*/
