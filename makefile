.PHONY = all build check lint clean clean-vendor

GOPATH := $(CURDIR)/src/vendor/:$(CURDIR)
export GOPATH
GOBIN = $(CURDIR)/src/vendor/bin

GOLINT = $(GOBIN)/golint
PIGEON = $(GOBIN)/pigeon
GOCYCLO = $(GOBIN)/gocyclo
MISSPELL = $(GOBIN)/misspell

SRC = src/main.go src/ast/ast.go
BUILD = src/wacc.go

all: pigeon build
	cd src && go build -o wacc

build: $(SRC) $(BUILD)

src/wacc.go: src/grammar/wacc.peg
	$(PIGEON) $^ > $@

check: spellcheck fmt vet lint cyclo

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

clean:
	rm -rf $(BUILD)
	rm -rf src/wacc

clean-vendor:
	rm -rf src/vendor
