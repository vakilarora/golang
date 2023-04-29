package main

import "fmt"

func main() {

	var a, b = 98, 7
	c := 98

	//Arithmetic Operators
	add, sub, mul, rem, div := a+b, a-b, a*b, a%b, a/b

	fmt.Printf("\nadd: %d, sub: %d, mul: %d, rem: %d, div: %d\n", add, sub, mul, rem, div)

	//Logical Operators
	or1 := ((a == b) || (a == c))
	and1 := ((a == b) && (a == c))

	fmt.Print(or1, and1)
	fmt.Println("")

	//Relational Operators
	fmt.Println("\nGreater than")
	fmt.Print(a > c)
	fmt.Println("\nGreater than or equal to")
	fmt.Print(a >= c)
	fmt.Println("\nEqual to")
	fmt.Print(a == c)
	fmt.Println("\nLess than")
	fmt.Print(a < c)
	fmt.Println("\nLess than or equal to")
	fmt.Print(a <= c)
	fmt.Println("")
}
