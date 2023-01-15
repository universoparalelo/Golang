package main

import "fmt"

func main() {
	var nombre string = "Celeste"
	var edad int = 20
	var (
		n1 = 10
		n2 = 45
	)
	var altura float32 = 1.65
	var puertaAbierta bool = true

	fmt.Println("Mi nombre es", nombre, "y tengo", edad, "años")
	fmt.Println(n1, ",", n2)
	fmt.Println("Mi altura es", altura, "y la puerta esta en", puertaAbierta)
	fmt.Println("n1 + n2 =", n1+n2)

	// shorthand
	isbool := true
	edad2 := 89

	fmt.Println(isbool, edad2)

	// formatting
	// https://pkg.go.dev/fmt#hdr-Printing
	fmt.Printf("%.2f \n", altura)
	fmt.Printf("%T \n", nombre) //muestra el tipo de dato
	fmt.Printf("%v \n", nombre)
	fmt.Printf("%t \n", puertaAbierta)
	fmt.Printf("%d \n", edad)
	fmt.Printf("%b \n", 20) // convierte el numero en binario
}
