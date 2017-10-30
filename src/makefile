.PHONY = all build lint golint peg clean

GOPATH := $(CURDIR)/vendor/:$(GOPATH)
export GOPATH
GOBIN := $(CURDIR)/vendor/bin/:$(GOBIN)
PATH := $(GOBIN):$(PATH)

GOLINT = golint
PIGEON = pigeon

SRC = main.go 
BUILD = wacc.go

all: build
	go build -o wacc

build: $(SRC) $(BUILD)

wacc.go: grammar/wacc.peg
	$(PIGEON) $^ > $@

lint: golint
	@echo "==Lint Checking=="
	$(GOLINT) $(SRC)

golint:
	go get github.com/golang/lint/golint

pigeon:
	go get github.com/pointlander/peg

clean:
	rm wacc $(BUILD)
