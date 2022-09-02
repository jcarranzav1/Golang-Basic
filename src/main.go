package main

import "fmt"

func main() {
	// Un map asigna claves a valores, parecido a diccionario en python o un hash map
	//MAPS

	//Declaration
	var map_1 = map[string]int32{
		"Car":      50000,
		"House":    20000,
		"Computer": 1000,
	}

	fmt.Println(map_1)

	//Another type of declaration
	map_2 := make(map[int]int)
	map_2[10] = 1
	map_2[12] = 4
	map_2[8] = 5
	map_2[11] = 2

	//Displaying a map
	for i, v := range map_2 {
		fmt.Println(i, "=", v) //There is no order at the moment of displaying a map
	}

	//Knowing if a value exist
	//If we display map_1["Car"] it will show the value
	value := map_1["Car"]
	fmt.Println("Car", value)

	//But what if it doesn't exist
	value = map_1["Vehicle"]
	fmt.Println("Vehicle", value) //It will show zero

	//If we want to know if it exists, it must be done with the second returing value
	value, ok := map_1["Vehicle"]
	//it will return a bool depending on if exists or not
	fmt.Println("Vehicle exist:", ok, "  value:", value)

	fmt.Println("\n******** Add new key/value pair and update ************")

	map_1["PS5"] = 9000
	map_1["House"] = 500000

	fmt.Println(map_1)

	fmt.Println("\n******** Delete  key/value pair  ***********")
	delete(map_1, "PS5")
	fmt.Println(map_1)

	fmt.Println("\n******** Modficacion del map  ***********")
	//  Como sabemos los mapas son de tipo referencial. Entonces, cuando asignamos un mapa existente a una nueva variable, ambos mapas aún se refieren a la misma estructura de datos subyacente. Entonces, cuando actualicemos un mapa, se reflejará en otro mapa.

	new_map := map_1

	new_map["Car"] = 1200
	fmt.Println("map_1:", map_1, " new_map:", new_map)

}
