export CGO_ENABLED := 0

.PHONY: build
build: 
	go build -o bin/taxcalc ./cmd/taxcalc 

.PHONY: run-server
run-server: 
	@chmod +x ./bin/taxcalc
	./bin/taxcalc -server