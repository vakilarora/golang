package main

import "fmt"

func main() {
	//declare and assign the value seperately
	var x1 int
	x1 = 10

	//declare and assign the value in same line
	var x2 float64 = 3.142

	//shorthand declaration

	x3 := "Vakil"

	fmt.Print("x1: ", x1)
	fmt.Print("\n")
	fmt.Printf("x2: %0.2f", x2)
	fmt.Print("\n")
	fmt.Println("x3:", x3)
}
