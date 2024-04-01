package calculator

import (
	"encoding/json"
	"os"
)

type Calculator interface {
	Tax(float64) float64
	InBracket(salary float64) bracket
}

type calculator struct {
	brackets []bracket
}

type bracket struct {
	From    int
	To      int
	Flat    float64
	Scaling float64
}

func (c *calculator) Tax(salary float64) float64 {
	bracket := c.InBracket(salary)
	return bracket.Flat + bracket.Scaling*(salary-float64(bracket.From))
}

func (c *calculator) InBracket(salary float64) bracket {
	if salary <= float64(c.brackets[0].From) {
		return c.brackets[0]
	}

	for _, bracket := range c.brackets {
		if salary > float64(bracket.From) && salary <= float64(bracket.To) {
			return bracket
		}
	}

	return c.brackets[len(c.brackets)-1]
}

func New(brackets []bracket) Calculator {
	return &calculator{brackets: brackets}
}

func NewFromJson(jsonString string) Calculator {
	content, err := os.ReadFile(jsonString)
	if err != nil {
		panic("Unable to read file")
	}
	brackets := []bracket{}
	json.Unmarshal(content, &brackets)
	return &calculator{
		brackets: brackets,
	}
}
