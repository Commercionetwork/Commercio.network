PACKAGES_NOSIMULATION=$(shell go list ./... | grep -v '/simulation')
BINDIR ?= $(GOPATH)/bin

export GO111MODULE = on

include Makefile.ledger

export GO111MODULE = on

all: tools build lint test

# The below include contains the tools and runsim targets.
include contrib/devtools/Makefile

########################################
### Install

install: go.sum
	go install -mod=readonly -tags "$(build_tags)" ./cmd/cnd
	go install -mod=readonly -tags "$(build_tags)" ./cmd/cncli

########################################
### Build

build: go.sum
ifeq ($(OS),Windows_NT)
	go build -mod=readonly -o ./build/cnd.exe -tags "$(build_tags)" ./cmd/cnd/main.go
	go build -mod=readonly -o ./build/cncli.exe -tags "$(build_tags)" ./cmd/cncli/main.go
else
	go build -mod=readonly -o ./build/cnd -tags "$(build_tags)" ./cmd/cnd/main.go
	go build -mod=readonly -o ./build/cncli -tags "$(build_tags)" ./cmd/cncli/main.go
endif


build-darwin: go.sum
	env GOOS=darwin GOARCH=amd64 go build -mod=readonly -o ./build/Darwin-AMD64/cncli -tags "$(build_tags)" ./cmd/cncli/main.go
	env GOOS=darwin GOARCH=amd64 go build -mod=readonly -o ./build/Darwin-AMD64/cnd -tags "$(build_tags)" ./cmd/cnd/main.go

build-linux: go.sum
	env GOOS=linux GOARCH=amd64 go build -mod=readonly -o ./build/Linux-AMD64/cncli -tags "$(build_tags)" ./cmd/cncli/main.go
	env GOOS=linux GOARCH=amd64 go build -mod=readonly -o ./build/Linux-AMD64/cnd -tags "$(build_tags)" ./cmd/cnd/main.go

build-windows: go.sum
	env GOOS=windows GOARCH=amd64 go build -mod=readonly -o ./build/Windows-AMD64/cncli.exe -tags "$(build_tags)" ./cmd/cncli/main.go
	env GOOS=windows GOARCH=amd64 go build -mod=readonly -o ./build/Windows-AMD64/cnd.exe -tags "$(build_tags)" ./cmd/cnd/main.go

build-all: go.sum
	make build-darwin
	make build-linux
	make build-windows

prepare-release: go.sum build-all
	rm -f ./build/Darwin-386.zip ./build/Darwin-AMD64.zip
	rm -f ./build/Linux-386.zip ./build/Linux-AMD64.zip
	rm -f ./build/Windows-386.zip ./build/Windows-AMD64.zip
	zip -jr ./build/Darwin-AMD64.zip ./build/Darwin-AMD64/cncli ./build/Darwin-AMD64/cnd
	zip -jr ./build/Linux-AMD64.zip ./build/Linux-AMD64/cncli ./build/Linux-AMD64/cnd
	zip -jr ./build/Windows-AMD64.zip ./build/Windows-AMD64/cncli.exe ./build/Windows-AMD64/cnd.exe

########################################
### Tools & dependencies

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	go mod verify

lint: golangci-lint
	$(BINDIR)/golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify
.PHONY: lint

########################################
### Testing

test: test_unit

test_unit:
	@VERSION=$(VERSION) go test -mod=readonly $(PACKAGES_NOSIMULATION) -tags='ledger test_ledger_mock'