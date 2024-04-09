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
	inputFormat  string
	outputFormat string
	server       bool
}

const (
	monthly  = "monthly"
	annually = "annually"
)

func flags() {
	flag.StringVar(&config.taxTablePath, "taxTablePath", "configs/taxTableSA.json", "The path to the tax table in json format")
	flag.Float64Var(&config.salary, "salary", 0, "The salary to calculate tax for")
	flag.StringVar(&config.inputFormat, "inputFormat", "monthly", "The output format - One of [monthly, annually]")
	flag.StringVar(&config.outputFormat, "outputFormat", "monthly", "The input format - One of [monthly, annually]")
	flag.BoolVar(&config.server, "server", false, "Specify this field to run a server on localhost:8080")
	flag.Parse()
}

func checkFlags() error {
	switch config.inputFormat {
	case monthly, annually:
		break
	default:
		return fmt.Errorf("input format not one of options")
	}
	switch config.outputFormat {
	case monthly, annually:
		break
	default:
		return fmt.Errorf("output format not one of options")
	}
	return nil
}

func formatInput(input float64) float64 {
	if config.inputFormat == monthly {
		input *= 12
	}
	return input
}

func formatOutput(output float64) string {
	if config.outputFormat == monthly {
		output /= 12.0
	}
	return calculator.FormatRands(output)
}

func runCli(calc calculator.Calculator) {
	salary := formatInput(config.salary)
	output := formatOutput(calc.Tax(salary))
	fmt.Println(output)
}

func runServer(calc calculator.Calculator) {
	srv := &server{
		port: 8080,
		calc: calc,
	}
	srv.Run()
}

func main() {
	flags()
	err := checkFlags()
	if err != nil {
		panic(err)
	}
	calc := calculator.NewFromJson(config.taxTablePath)
	if config.server {
		runServer(calc)
	} else {
		runCli(calc)
	}
}
