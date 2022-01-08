package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlus(t *testing.T) {
	assert.Equal(t, Plus(1,2), 3.0)
	assert.Equal(t, Plus(1,-1), 0.0)
}

func TestMinus(t *testing.T) {
	assert.Equal(t, Minus(1,2), -1.0)
	assert.Equal(t, Minus(1,-1), 2.0)
}

func TestMutiply(t *testing.T) {
	assert.Equal(t, Multiply(1,2), 2.0)
	assert.Equal(t, Multiply(1,-1), -1.0)
}

func TestDivide(t *testing.T) {
	t.Run("check divide by positive number", func(t *testing.T){
		result, err := Divide(5,2)
		assert.NoError(t, err)
		assert.Equal(t, result, 2.5)
	})

	t.Run("check divide by 0", func(t *testing.T){
		result, err := Divide(5,0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cannot divide by 0")
		assert.Equal(t, 0.0, result)
	})
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		name string
		input uint
		output uint
	}{
		{
			name: "N is 0",
			input: 0,
			output: 1,
		},
		{
			name: "N is 1",
			input: 1,
			output: 1,
		},
		{
			name: "N is 3",
			input: 3,
			output: 6,
		},
		{
			name: "N is 10",
			input: 10,
			output: 3628800,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := Factorial(tt.input); result != tt.output {
				t.Errorf("Factorial() = %v, output %v", result, tt.output)
			}
		})
	}
}