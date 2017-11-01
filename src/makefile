.PHONY = all build lint golint pigeon clean check spellcheck fmt vet lint cyclo

PATHDOWN = $(subst $(notdir ${CURDIR}),,${CURDIR})
GOPATH := $(CURDIR)/vendor/:$(PATHDOWN):$(GOPATH)
export GOPATH
GOBIN = $(CURDIR)/vendor/bin

GOLINT = $(GOBIN)/golint
PIGEON = $(GOBIN)/pigeon
GOCYCLO = $(GOBIN)/gocyclo
MISSPELL = $(GOBIN)/misspell

SRC = main.go ast/ast.go
BUILD = wacc.go

all: pigeon build
	go build -o wacc

build: $(SRC) $(BUILD)

wacc.go: grammar/wacc.peg
	$(PIGEON) $^ > $@

check: spellcheck fmt vet lint cyclo

lint: golint
	@echo "== Lint Checking =="
	$(GOLINT) $(SRC)

fmt:
	@echo "== Formatting =="
	gofmt -s $(SRC)

vet:
	@echo "== Vetting =="
	go tool vet $(SRC)

cyclo: gocyclo
	@echo "== Cyclo =="
	$(GOCYCLO) $(SRC)

spellcheck: misspell
	@echo "== Spell Checking =="
	$(MISSPELL) $(SRC)

golint:
	go get github.com/golang/lint/golint

pigeon:
	go get github.com/mna/pigeon

gocyclo:
	go get github.com/fzipp/gocyclo

misspell:
	go get github.com/client9/misspell/cmd/misspell

clean:
	rm $(BUILD)
	rm wacc

clean-vendor:
	rm -rf vendor
