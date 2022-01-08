// https://tutorialedge.net/golang/concurrency-with-golang-goroutines/
// https://tutorialedge.net/golang/go-waitgroup-tutorial/

package main

import (
	"fmt"
	"sync"
	"time"
)

func compute(limit int, multiplier int, wg *sync.WaitGroup) {
	for i := 1; i <= limit; i++ {
		time.Sleep(time.Duration(multiplier) * time.Second)
		fmt.Println(multiplier * i)
	}
	wg.Done()
}

func main() {
	fmt.Println("Goroutine Tutorial")

	var wg sync.WaitGroup

	wg.Add(1)
	go compute(10, 1, &wg)

	wg.Add(1)
	go compute(10, 2, &wg)

	fmt.Println("Keep process alive for goroutines to run")
	wg.Wait()
	fmt.Println("Process end")
}
