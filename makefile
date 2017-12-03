# ***************** GOPATH CONFIG ****************

GOPATH := $(CURDIR)/lib/:$(CURDIR)
export GOPATH

PIGEON = $(CURDIR)/lib/bin/pigeon
GOLINT = $(CURDIR)/lib/bin/gometalinter

SRC = src/**/*.go
GRAMMAR = src/grammar/bootstrap.peg src/grammar/wacc.peg src/grammar/*.peg

# ***************** BUILDING ****************

.PHONY: all

all: $(SRC)
	go build gowacc

src/gowacc/wacc.go: $(GRAMMAR)
	@go get github.com/ic-projects/pigeon
	cat $^ | $(PIGEON) > $@

# ***************** LINTING and TESTING ****************

.PHONY: check lint fmt tests

check: lint fmt tests

lint:
	@go get github.com/alecthomas/gometalinter
	@$(GOLINT) --install --update
	@echo "\n== Linting =="
	$(GOLINT) --enable-all --skip=lib ./...

fmt:
	@echo "\n== Formatting =="
	gofmt -s -w $(SRC)

tests:
	@echo "\n== Testing at http://localhost:18000/ =="
	ruby test/testserver.rb .

# ***************** DOCUMENTATION ****************

.PHONY: docs

docs:
	@go get golang.org/x/tools/cmd/godoc
	@echo "\n== Generating docs at http://localhost:8080/pkg/gowacc/ =="
	godoc -http=:8080 -goroot=$(CURDIR)

# ***************** CLEANING ****************

.PHONY: clean clean-lib

clean:
	rm -rf $(BUILD)
	rm -rf gowacc
	rm -rf *.s

clean-lib:
	rm -rf lib
