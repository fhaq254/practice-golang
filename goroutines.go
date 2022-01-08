// https://tutorialedge.net/golang/concurrency-with-golang-goroutines/

package main

import (
	"fmt"
	"time"
)

func compute(limit int, multiplier int) {
	for i := 1; i <= limit; i++ {
		time.Sleep(time.Duration(multiplier) * time.Second)
		fmt.Println(multiplier * i)
	}
}

func main() {
	fmt.Println("Goroutine tutorial")

	// sync function
	// compute(10)
	// compute(10)

	// make async function by add 'go'
	go compute(10, 1)
	go compute(10, 2)

	fmt.Println("Keep process alive for goroutines to run")
	time.Sleep(25 * time.Second)
	fmt.Println("Process end")
}
