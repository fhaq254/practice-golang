package main

import "fmt"

func test() {
	fmt.Println("test")
}

func operate(x int, y int) (int, int) {
	return x + y, x - y
}

func operate2(x int, y int) (sum int, diff int) {
	sum = x + y
	diff = x - y
	return
}

func main() {
	test()

	sum1, diff1 := operate(9, 4)
	fmt.Println(sum1, diff1)

	sum2, diff2 := operate(9, 4)
	fmt.Println(sum2, diff2)
}
