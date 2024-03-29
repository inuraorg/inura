GITCOMMIT ?= $(shell git rev-parse HEAD)
GITDATE ?= $(shell git show -s --format='%ct')
VERSION := v0.0.0

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X github.com/inuraorg/inura/inura-program/version.Version=$(VERSION)
LDFLAGSSTRING +=-X github.com/inuraorg/inura/inura-program/version.Meta=$(VERSION_META)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

COMPAT_DIR := temp/compat

inura-program: \
	inura-program-host \
	inura-program-client \
	inura-program-client-mips

inura-program-host:
	env GO111MODULE=on GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) go build -v $(LDFLAGS) -o ./bin/inura-program ./host/cmd/main.go

inura-program-client:
	env GO111MODULE=on GOOS=$(TARGETOS) GOARCH=$(TARGETARCH) go build -v $(LDFLAGS) -o ./bin/inura-program-client ./client/cmd/main.go

inura-program-client-mips:
	env GO111MODULE=on GOOS=linux GOARCH=mips GOMIPS=softfloat go build -v $(LDFLAGS) -o ./bin/inura-program-client.elf ./client/cmd/main.go
	# verify output with: readelf -h bin/inura-program-client.elf
	# result is mips32, big endian, R3000

clean:
	rm -rf bin "$(COMPAT_DIR)"

test:
	go test -v ./...

verify-goerli: inura-program-host inura-program-client
	env GO111MODULE=on go run ./verify/cmd/goerli.go --l1 $$L1URL --l2 $$L2URL

capture-goerli-verify: inura-program-host inura-program-client
	rm -rf "$(COMPAT_DIR)/goerli" "$(COMPAT_DIR)/goerli.tar.bz"
	env GO111MODULE=on go run ./verify/cmd/goerli.go --l1 $$L1URL --l2 $$L2URL --datadir "$(COMPAT_DIR)/goerli"
	tar jcf "$(COMPAT_DIR)/goerli.tar.bz" -C "$(COMPAT_DIR)" goerli

capture-chain-test-data: capture-goerli-verify

run-goerli-verify: inura-program-host inura-program-client
	mkdir -p "$(COMPAT_DIR)"
	curl -L -o "$(COMPAT_DIR)/goerli.tar.bz" https://github.com/ethereum-optimism/chain-test-data/releases/download/2023-10-11/goerli.tar.bz
	tar jxf "$(COMPAT_DIR)/goerli.tar.bz" -C "$(COMPAT_DIR)"
	./bin/inura-program `cat "$(COMPAT_DIR)/goerli/args.txt"`

.PHONY: \
	inura-program \
	clean \
	test \
	lint
