package main

import "fmt"

func changeByFunc(str string) {
	str = "changed by function"
}

func changeByPointer(str *string) {
	*str = "changed by pointer"
}

func main() {
	x := 7
	fmt.Println("x", x)

	// pointer - memory address (&)
	y := &x
	fmt.Println("y", y)

	// assign value to pointer
	*y = 8
	fmt.Println("x", x)

	text := "original"
	fmt.Println(text)

	changeByFunc(text) // change only inside function
	fmt.Println(text)

	changeByPointer(&text) // change value in memory address directly
	fmt.Println(text)
}
