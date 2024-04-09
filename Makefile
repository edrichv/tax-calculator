export CGO_ENABLED := 0

.PHONY: build
build: 
	go build -o bin/taxcalc ./cmd/taxcalc 