package main

import "fmt"

func test1(x int) int {
	return x * 2
}

func test2(input func(int) int) {
	fmt.Println(input(7))
}

func returnFunc(x string) func() {
	return func() { fmt.Println(x) }
}

func main() {
	// anonymous function
	test := func() {
		fmt.Println("Hello")
	}
	test()

	// function as parameter
	test2(test1)

	// return function
	printHi := returnFunc("Hiiii")
	printHi()
}
