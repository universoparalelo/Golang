package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d, ", i)
	}
	fmt.Println("")

	j := 0
	for j <= 10 {
		fmt.Printf("%d,", j)
		j++
	}
	fmt.Println("")

	for i := 1; i <= 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

	//if statements
	var age int = 17

	if age < 18 {
		fmt.Println("You can't vote yet")
	} else {
		fmt.Println("You can vote")
	}

	switch age {
	case 16:
		fmt.Println("You can't vote")
	case 18:
		fmt.Println("You can vote")
	default:
		fmt.Println("I dont know what to tell you")
	}

}
