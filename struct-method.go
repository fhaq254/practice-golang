package main

import "fmt"

type Student struct {
	name  string
	score []int
	age   int
}

func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func main() {
	s1 := Student{
		"Tim",
		[]int{70, 90, 80, 85},
		19,
	}
	fmt.Println(s1.getAge())

	s1.setAge(22)
	fmt.Println(s1.getAge())
}
