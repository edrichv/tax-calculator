package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/edrichv/tax-calculator/pkg/calculator"
	"github.com/gorilla/mux"
)

type server struct {
	port uint16
	calc calculator.Calculator
}

type Server interface {
	Run()
}

func (s *server) handleSalary(w http.ResponseWriter, r *http.Request) {
	salary, err := strconv.ParseFloat(mux.Vars(r)["salary"], 64)
	if err == nil {
		tax := s.calc.Tax(salary)
		w.Write([]byte(calculator.FormatRands(tax)))
	}
}

func (s *server) Run() {
	r := mux.NewRouter()
	r.HandleFunc("/salary/{salary}", s.handleSalary)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), r))
}
