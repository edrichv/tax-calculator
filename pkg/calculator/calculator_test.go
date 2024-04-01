package calculator

import (
	"fmt"
	"testing"
)

func TestMockCalculator(t *testing.T) {
	calc := NewMock()
	salary := float64(50)
	fmt.Println(calc.Tax(salary))
	fmt.Println(calc.InBracket(salary))
}

func NewMock() Calculator {
	return &calculator{
		[]bracket{
			{
				0, 10, 0, 10,
			},
			{
				11, 30, 0.5, 20,
			},
			{
				31, 50, 0.8, 30,
			},
		},
	}
}
