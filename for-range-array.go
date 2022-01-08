package main

import "fmt"

func main() {
	var a []int = []int{1, 3, 5, 7, 11, 13}

	fmt.Println("index and value")
	for index, value := range a {
		fmt.Printf("%d : %d \n", index, value)
	}

	fmt.Println("index only")
	for index := range a {
		fmt.Printf("%d \n", index)
	}

	fmt.Println("value only")
	for _, value := range a {
		fmt.Printf("%d \n", value)
	}
}
