package main

import "fmt"

func main() {
	var name string
	fmt.Println("string >> ", name) // empty string

	var number1 uint8
	fmt.Println("unsigned int8 >> ", number1) // 0

	var number2 int8
	fmt.Println("signed int8 >> ", number2) // 0

	var number3 float32
	fmt.Println("float32 >> ", number3) // 0

	var boolean bool
	fmt.Println("bool >> ", boolean) // false
}
