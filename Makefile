SHELL := /bin/bash -eu -o pipefail
SRC := $(shell find . -path ./vendor -prune -o -name '*.go')
COVERAGE_PROFILE := results/coverage.out
COVERAGE_HTML := results/coverage.html

.PHONY: all
all: install

.PHONY: install
install: get test vet
	go install -v ./...

.PHONY: get
get:
	go get -d -u github.com/gonum/blas/blas64
	go get -d -u github.com/gonum/matrix/mat64

.PHONY: test
test: $(COVERAGE_PROFILE)

.PHONY: html
html: $(COVERAGE_HTML)

.PHONY: benchmark
benchmark:
	go test -benchmem -cpuprofile=profile.out -run='^$$' github.com/nfisher/goalgo/mat -bench='^Benchmark_DotLarge$$' -benchtime=20s
	#go test -benchmem -run=^$$ github.com/nfisher/goalgo/mat -bench=^Benchmark_Dot$$ -benchtime=20s

results:
	mkdir -p results

$(COVERAGE_PROFILE): results $(SRC)
	go test -v -coverprofile=$(COVERAGE_PROFILE) ./...

$(COVERAGE_HTML): $(COVERAGE_PROFILE)
	go tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)

.PHONY: vet
vet:
	go vet -all ./...
