PACKAGES_NOSIMULATION=$(shell go list ./... | grep -v '/simulation')
BINDIR ?= $(GOPATH)/bin
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

REPO_ROOT := $(shell git rev-parse --show-toplevel)

GENERATE ?= 1

export GO111MODULE = on

LEDGER_ENABLED ?= true

########################################
### Build tags

build_tags = netgo

ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq ($(OS),Windows_NT)
    TARGET_BIN = Windows-AMD64
    TARGET_BUILD = build-windows
else
	UNAME_S = $(shell uname -s)
	ifeq ($(UNAME_S),Darwin)
		TARGET_BIN = Darwin-AMD64
		TARGET_BUILD = build-darwin
	else
		TARGET_BIN = Linux-AMD64
		TARGET_BUILD = build-linux
	endif
endif



build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.ServerName=cnd \
		  -X github.com/cosmos/cosmos-sdk/version.ClientName=cndcli \
		  -X github.com/cosmos/cosmos-sdk/version.Name=commercionetwork \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

all: tools build lint test

# The below include contains the tools and runsim targets.
include contrib/devtools/Makefile

########################################
### Install

install: git-hooks go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/cnd
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/cncli

########################################
### Build

build: go.sum generate
ifeq ($(OS),Windows_NT)
	go build -mod=readonly -o ./build/cnd.exe $(BUILD_FLAGS) ./cmd/cnd
	go build -mod=readonly -o ./build/cncli.exe $(BUILD_FLAGS) ./cmd/cncli
else
	go build -mod=readonly -o ./build/cnd $(BUILD_FLAGS) ./cmd/cnd
	go build -mod=readonly -o ./build/cncli $(BUILD_FLAGS) ./cmd/cncli
endif



build-darwin: go.sum generate
	env GOOS=darwin GOARCH=amd64 go build -mod=readonly -o ./build/Darwin-AMD64/cncli $(BUILD_FLAGS) ./cmd/cncli
	env GOOS=darwin GOARCH=amd64 go build -mod=readonly -o ./build/Darwin-AMD64/cnd $(BUILD_FLAGS) ./cmd/cnd

build-linux: go.sum generate
	env GOOS=linux GOARCH=amd64 go build -mod=readonly -o ./build/Linux-AMD64/cncli $(BUILD_FLAGS) ./cmd/cncli
	env GOOS=linux GOARCH=amd64 go build -mod=readonly -o ./build/Linux-AMD64/cnd $(BUILD_FLAGS) ./cmd/cnd

build-local-linux: go.sum generate
	env GOOS=linux GOARCH=amd64 go build -mod=readonly -o ./build/cncli $(BUILD_FLAGS) ./cmd/cncli
	env GOOS=linux GOARCH=amd64 go build -mod=readonly -o ./build/cnd $(BUILD_FLAGS) ./cmd/cnd

build-windows: go.sum generate
	env GOOS=windows GOARCH=amd64 go build -mod=readonly -o ./build/Windows-AMD64/cncli.exe $(BUILD_FLAGS) ./cmd/cncli
	env GOOS=windows GOARCH=amd64 go build -mod=readonly -o ./build/Windows-AMD64/cnd.exe $(BUILD_FLAGS) ./cmd/cnd

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

git-hooks:
	@git config --local core.hooksPath $(REPO_ROOT)/.githooks/

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: git-hooks go.mod
	@echo "--> Ensure dependencies have not been modified"
	go mod verify

lint:
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s
	go mod verify

generate:
ifeq ($(GENERATE),1)
	go generate ./...
endif

.PHONY: git-hooks

########################################
### Testing

test: test_unit

test_unit:
	@VERSION=$(VERSION) go test -mod=readonly $(PACKAGES_NOSIMULATION) -tags='ledger test_ledger_mock'

.PHONY: lint test test_unit go-mod-cache

build-docker-cndode:
	$(MAKE) -C contrib/localnet


localnet-start: localnet-stop build-local-linux
	@if ! [ -f build/node0/cnd/config/genesis.json ]; then docker run --rm -v $(CURDIR)/build:/cnd:Z commercionetwork/cndnode testnet --v 4 -o . --starting-ip-address 192.168.10.2 --keyring-backend=test ; fi
	@if ! [ -f build/nginx/nginx.conf ]; then cp -r contrib/localnet/nginx build/nginx; fi
	docker-compose up

localnet-reset: localnet-stop $(TARGET_BUILD)
	@for node in 0 1 2 3; do build/$(TARGET_BIN)/cnd unsafe-reset-all --home ./build/node$$node/cnd; done


localnet-stop:
	docker-compose down

clean:
	rm -rf build/

.PHONY: localnet-start localnet-stop build-docker-cndode clean localnet-reset
