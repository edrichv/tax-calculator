package main

import (
	"flag"
	"fmt"

	"github.com/edrichv/tax-calculator/pkg/calculator"
)

var config Config

type Config struct {
	taxTablePath string
	salary       float64
}

func flags() {
	flag.StringVar(&config.taxTablePath, "taxTablePath", "configs/taxTableSA.json", "The path to the tax table in json format")
	flag.Float64Var(&config.salary, "salary", 0, "The salary to calculate tax for")
	flag.Parse()
}

func main() {
	flags()
	calc := calculator.NewFromJson(config.taxTablePath)
	fmt.Println(calc.Tax(config.salary))
}
