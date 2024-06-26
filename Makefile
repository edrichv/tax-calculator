export CGO_ENABLED := 0

.PHONY: build
build: 
	go build -o bin/taxcalc ./cmd/taxcalc 

.PHONY: run-server
run-server: build
	@chmod +x ./bin/taxcalc
	./bin/taxcalc -server

.PHONY: gofmt
gofmt: 
	gofmt -s -w -l ./..

.PHONY: clean
clean:
	git clean -xdf

.PHONY: test
test:
	go test -v ./...