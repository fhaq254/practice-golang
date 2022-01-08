// https://medium.com/rungo/string-formatting-in-go-dd752ff7f60

// fmt.Print - fmt.Sprint
// fmt.Println - fmt.Sprintln
// fmt.Printf - fmt.Sprintf

package main

import "fmt"

func main() {
	fmt.Print("fmt.Print \n")
	fmt.Print("1 ", 2, true, 3, 4, nil, "\n")

	fmt.Print("fmt.Sprint \n")
	strSprint := fmt.Sprint("1 ", 2, true, 3, 4, nil, "\n")
	fmt.Print(strSprint)
}
