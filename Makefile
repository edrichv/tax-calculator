export CGO_ENABLED := 0

.PHONY: build
build: 
	go build ./cmd/taxcalc