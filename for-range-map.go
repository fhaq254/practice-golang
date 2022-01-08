package main

import "fmt"

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Println("key and value")
	for key, value := range m {
		fmt.Printf("%s : %d \n", key, value)
	}

	fmt.Println("key only")
	for key := range m {
		fmt.Printf("%s \n", key)
	}

	fmt.Println("value only")
	for _, value := range m {
		fmt.Printf("%d \n", value)
	}
}
