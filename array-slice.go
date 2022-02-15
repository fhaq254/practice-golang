package main

import "fmt"

func main() {
	var d [5]string = [5]string{"A", "B", "C", "D", "E"}
	fmt.Println(d)

	// empty array
	var s []string
	fmt.Println(s)

	// append
	s = append(s, "Z")
	fmt.Println(s)

	// slice
	s = x[1:3]
	fmt.Println(s) // ["B", "C"]

}
