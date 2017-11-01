.PHONY = all build lint golint pigeon clean

PATHDOWN = $(subst $(notdir ${CURDIR}),,${CURDIR})
GOPATH := $(CURDIR)/vendor/:$(PATHDOWN):$(GOPATH)
export GOPATH
GOBIN = $(CURDIR)/vendor/bin

GOLINT = $(GOBIN)/golint
PIGEON = $(GOBIN)/pigeon

SRC = main.go
BUILD = wacc.go

all: pigeon build
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
	go get github.com/mna/pigeon

clean:
	rm $(BUILD)
	rm wacc
