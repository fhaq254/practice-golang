// https://medium.com/rungo/string-formatting-in-go-dd752ff7f60

package main

import "fmt"

func main() {
	b := true
	i := 123
	f := 987.654
	c := 3 + 7i
	str := "text"
	arr := []int{1, 2, 3}
	m := map[string]int{"one": 1, "two": 2}
	s := struct {
		name string
		age  int
	}{"John", 26}

	fmt.Print("Standard Print with manual newline \n")
	fmt.Print("1", 2, true, 3, 4, nil, "\n")
	fmt.Print("\n")

	fmt.Println("Standard Print with auto newline")
	fmt.Println("1", 2, true, 3, 4, nil)
	fmt.Println()

	fmt.Printf("String interpolation \n")
	fmt.Printf("%%v - default format in golang \n")
	fmt.Printf("%v \n", b)   // boolean
	fmt.Printf("%v \n", i)   // integer
	fmt.Printf("%v \n", f)   // float
	fmt.Printf("%v \n", c)   // complex number
	fmt.Printf("%v \n", str) // string
	fmt.Printf("%v \n", arr) // array
	fmt.Printf("%v \n", m)   // map
	fmt.Printf("%v \n", s)   // struct
	fmt.Printf("%v \n", c)   // channel
	fmt.Printf("%v \n", &s)  // pointer
	fmt.Printf("\n")

	fmt.Printf("String interpolation \n")
	fmt.Printf("%%#v - syntax representation \n")
	fmt.Printf("%#v \n", b)   // boolean
	fmt.Printf("%#v \n", i)   // integer
	fmt.Printf("%#v \n", f)   // float
	fmt.Printf("%#v \n", c)   // complex number
	fmt.Printf("%#v \n", str) // string
	fmt.Printf("%#v \n", arr) // array
	fmt.Printf("%#v \n", m)   // map
	fmt.Printf("%#v \n", s)   // struct
	fmt.Printf("%#v \n", c)   // channel
	fmt.Printf("%#v \n", &s)  // pointer
	fmt.Printf("\n")

	fmt.Printf("String interpolation \n")
	fmt.Printf("%%T - type format \n")
	fmt.Printf("%T \n", b)   // boolean
	fmt.Printf("%T \n", i)   // integer
	fmt.Printf("%T \n", f)   // float
	fmt.Printf("%T \n", c)   // complex number
	fmt.Printf("%T \n", str) // string
	fmt.Printf("%T \n", arr) // array
	fmt.Printf("%T \n", m)   // map
	fmt.Printf("%T \n", s)   // struct
	fmt.Printf("%T \n", c)   // channel
	fmt.Printf("%T \n", &s)  // pointer
	fmt.Printf("\n")

	fmt.Printf("String interpolation \n")
	fmt.Printf("%%d - base10 format \n")
	fmt.Printf("%%b - base2 format \n")
	fmt.Printf("%%0 - base8 format \n")
	fmt.Printf("%%h - base16 format \n")
	fmt.Printf("%%f - decimal format \n")
	fmt.Printf("%%e - scientific notation format \n")
}
