package main

import "fmt"

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(m)

	// len
	fmt.Println(len(m))

	// get value from existing key
	fmt.Println(m["one"]) // 1

	// get value from non-existing key >> default value
	fmt.Println(m["four"]) // 0

	// add
	m["four"] = 4
	fmt.Println(m)

	// delete
	delete(m, "two")
	fmt.Println(m) // if key exists, delete

	delete(m, "two")
	fmt.Println(m) // if key do not exist, do nothing

	// check if key exist
	val1, ok1 := m["one"]
	fmt.Println(val1, ok1) // if key exist, ok = true

	val2, ok2 := m["two"]
	fmt.Println(val2, ok2) // if key do not exist, ok = false
}
