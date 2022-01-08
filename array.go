package main

import "fmt"

func main() {
	// create array
	var arr [5]int
	fmt.Println("default array of int", arr)

	// assign value
	arr[0] = 100
	arr[4] = 500
	fmt.Println("post-assigned array of int >>", arr)

	// create & assign
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("pre-assigned array of int >>", arr2)

	// len & indexing
	sum := 0
	for i := 0; i < len(arr2); i++ {
		sum += arr2[i]
	}
	fmt.Println("sum value in arr2 >>", sum)

	// 2D array
	arr2D := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("2D array >>", arr2D)
	fmt.Println("first index of second array >>", arr2D[1][0])
}
