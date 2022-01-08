package calculator

import (
	"errors"
)

func Plus(a, b float64) float64 {
	return a+b
}

func Minus(a, b float64) float64 {
	return a-b
}

func Multiply(a, b float64) float64 {
	return a*b
}

func Divide(a, b float64) (float64, error) {
	if b==0 {
		return 0, errors.New("cannot divide by 0")
	}
	return a/b, nil
}

func Factorial(n uint) uint {
	if n <= 1 {
		return 1
	}

	return n * Factorial(n-1)
}