.PHONY = all build check lint fmt vet cyclo spellcheck test doc clean clean-lib

GOPATH := $(CURDIR)/lib/:$(CURDIR)
export GOPATH
GOBIN = $(CURDIR)/lib/bin

GOLINT = $(GOBIN)/golint
PIGEON = $(GOBIN)/pigeon
GOCYCLO = $(GOBIN)/gocyclo
MISSPELL = $(GOBIN)/misspell

SRC = src/**/*.go
BUILD = src/gowacc/wacc.go

all: pigeon build
	go build gowacc

build: $(SRC) $(BUILD)

src/gowacc/wacc.go: src/grammar/bootstrap.peg src/grammar/wacc.peg src/grammar/*.peg
	cat $^ | $(PIGEON) > $@

check: spellcheck fmt vet lint cyclo tests

lint: golint
	@echo "\n== Lint Checking =="
	$(GOLINT) $(SRC)

fmt:
	@echo "\n== Formatting =="
	gofmt -w $(SRC)

vet:
	@echo "\n== Vetting =="
	go tool vet $(SRC)

cyclo: gocyclo
	@echo "\n== Cyclo =="
	$(GOCYCLO) $(SRC)

spellcheck: misspell
	@echo "\n== Spell Checking =="
	$(MISSPELL) $(SRC)

golint:
	@go get github.com/golang/lint/golint

pigeon:
	@go get github.com/ic-projects/pigeon

gocyclo:
	@go get github.com/fzipp/gocyclo

misspell:
	@go get github.com/client9/misspell/cmd/misspell

tests:
	@echo "\n== To view the gowacc tests, visit http://localhost:18000/ =="
	ruby test/testserver.rb .

docs:
	@echo "\n== To view the gowacc documentation, visit http://localhost:8080/pkg/gowacc/ =="
	godoc -http=:8080 -goroot=$(CURDIR)

clean:
	rm -rf $(BUILD)
	rm -rf gowacc
	rm -rf *.s

clean-lib:
	rm -rf lib
