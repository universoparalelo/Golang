package main

import "fmt"

func main() {
	// array vacio
	var evenNum [5]int
	numerosPar := [5]int{1, 2, 3, 4, 5}

	fmt.Print(evenNum)
	fmt.Print(numerosPar, "\n")

	// recorrer un arreglo con un for
	for i, value := range numerosPar {
		fmt.Println(value, i)
	}

	var numSlice = []int{1, 2, 3, 6, 5, 4}
	slice := numSlice[2:5]
	fmt.Print(slice, "\n")
}
