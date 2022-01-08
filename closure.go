// https://tutorialedge.net/golang/go-closures-tutorial/

package main

import (
	"fmt"
)

func createReducer(value int, reducer int) func() int {
	return func() int {
		value -= reducer
		return value
	}
}

func main() {
	fmt.Println("Create Reducer1 start from 10")
	reducer1 := createReducer(10, 1)

	fmt.Println("Create Reducer2 start from 20")
	reducer2 := createReducer(20, 2)

	fmt.Println("Call Reducer1")
	fmt.Println(reducer1())
	fmt.Println(reducer1())
	fmt.Println(reducer1())
	fmt.Println(reducer1())

	fmt.Println("Call Reducer2")
	fmt.Println(reducer2())
	fmt.Println(reducer2())
	fmt.Println(reducer2())
	fmt.Println(reducer2())
}
