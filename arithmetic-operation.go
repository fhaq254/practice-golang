package main

import "fmt"

func main() {
	// explicit type declaration
	var num1 int8 = 9
	var num2 int8 = 4

	// implicit type declaration
	answer1 := num1 / num2
	fmt.Println(answer1) // int8

	answer2 := float32(num1) / float32(num2)
	fmt.Println(answer2) //float32

	// compile error >> code cannot be run
	// answer3 := float32(num1) / float64(num2)
	// answer4 := float32(num1) / int32(num2)

	// runtime error >> code can be run, but it will crash
	// answer5 := 5 / 0
}
