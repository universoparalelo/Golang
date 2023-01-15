package main

import "fmt"

func main() {
	// creando un map de strings
	// studentAge := make(map[string]int)

	// studentAge["Arrya"] = 23
	// studentAge["Celeste"] = 27
	// studentAge["Ludmila"] = 22
	// studentAge["Bianca"] = 54
	// studentAge["Miley"] = 40
	// studentAge["Constanza"] = 32
	// studentAge["Laura"] = 20

	// fmt.Println(studentAge)

	animales := map[string]map[string]string{

		"mamifero": {
			"vaca":  "vacuno",
			"cerdo": "bobino",
		},

		"oviparo": {
			"serpiente": "reptil",
			"pelicano":  "ave",
		},
	}
	fmt.Println(animales["oviparo"])

	//another one
	nested := map[int]map[string]string{
		1: {
			"a": "Apple",
			"b": "Banana",
			"c": "Coconut",
		},
		2: {
			"a": "Tea",
			"b": "Coffee",
			"c": "Milk",
		},
		3: {
			"a": "Italian Food",
			"b": "Indian Food",
			"c": "Chinese Food",
		},
	}

	fmt.Println(nested)

	// llamada a la funcion sumar
	fmt.Println("5 + 4 =", sumar(5, 4))
}

// funcion para sumar
func sumar(a, b int) int {
	return a + b
}
