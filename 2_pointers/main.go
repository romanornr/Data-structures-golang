package main

import "fmt"

func main() {
	x := 1
	y := &x

	fmt.Println("Value of x: ", x)
	fmt.Println("The addr of x: ", &x)
	fmt.Println("The value of y is: ", y)

	*y = 100
	fmt.Println("The value of x is: ", x)
	fmt.Println("The addr of x is: ", &x)
	fmt.Println("The value of y is: ", y)
}
