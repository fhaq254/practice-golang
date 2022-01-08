// https://tutorialedge.net/golang/the-go-init-function/

// init()
// regardless of how many times that package is imported,
// the init() function will only be called once.

package main

import "fmt"

var name string

func init() {
	fmt.Println("This will get called on main initialization")
	name = "Elliot"
}

func main() {
	fmt.Println("My wonderful Golang program")
	fmt.Printf("Name: %s\n", name)
}
