package main

import "fmt"

func main() {

	for x := 0; x < 10; x += 1 {
		output := ""
		if x < 5 {
			output = fmt.Sprintf("x = %d is less than 5", x)
		} else if x == 5 {
			output = fmt.Sprintf("x = %d is equal to 5", x)
		} else {
			output = fmt.Sprintf("x = %d is more than 5", x)
		}
		fmt.Println(output)
	}
}
