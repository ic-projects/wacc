.PHONY = all build lint golint peg clean

GOPATH := $(CURDIR)/vendor/:$(GOPATH)
export GOPATH
BIN = $(CURDIR)/vendor/bin
GOLINT = $(BIN)/golint
PEG = $(BIN)/peg

SRC = main.go 
BUILD = wacc.go

all: peg build
	go build -o wacc

build: $(SRC) $(BUILD)

wacc.go: grammar/wacc.peg
	$(PEG) grammar/wacc.peg
	mv grammar/wacc.peg.go wacc.go

lint: golint
	@echo "==Lint Checking=="
	$(BIN)/golint $(SRC)

golint:
	go get github.com/golang/lint/golint

peg:
	go get github.com/pointlander/peg

clean:
	rm wacc $(BUILD)
