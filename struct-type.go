package main

import "fmt"

type Point struct {
	x int32
	y int32
}

type Circle struct {
	radius float64
	center Point
}

func main() {
	var p1 Point = Point{1, 2}
	fmt.Println("p1.x >>", p1.x)
	fmt.Println("p1.y >>", p1.y)

	var c1 Circle = Circle{5.5, p1}
	fmt.Println("c1.radius >>", c1.radius)
	fmt.Println("c1.center.x >>", c1.center.x)
	fmt.Println("c1.center.y >>", c1.center.y)
}
