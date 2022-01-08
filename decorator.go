// https://tutorialedge.net/golang/go-decorator-function-pattern-tutorial/

package main

import (
	"fmt"
)

func coreFunc() {
	fmt.Println("## Hello from Core Function ##")
}

func decoratorFunc(anyFunc func()) {
	fmt.Println(".. Start from Decorator Function ..")
	anyFunc()
	fmt.Println(".. End of Decorator Function ..")
}

func main() {
	decoratorFunc(coreFunc)
}
