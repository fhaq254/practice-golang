// https://tutorialedge.net/golang/go-channels-tutorial/

package main

import (
	"fmt"
	"math/rand"
)

func CalculateValue(chValue chan int) {
	calValue := rand.Intn(10)
	fmt.Println("Calculated Value: ", calValue)
	chValue <- calValue // sends a value to a channel
}

func main() {
	fmt.Println("Channel tutorial")

	chValue := make(chan int) // setup channel
	defer close(chValue)      // teardown channel

	go CalculateValue(chValue)
	fmt.Println("", chValue) // memory address >> pointer

	mainValue := <-chValue // receives a value from a channel
	fmt.Println("", mainValue)

	// Notice how when we execute this program,
	// it doesn’t immediately terminate.
	// This is because the act of sending to and receiving from a channel are blocking.
	// Our main() function 'blocks' until it receives a value from our channel.

}
