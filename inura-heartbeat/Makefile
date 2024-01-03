GITCOMMIT ?= $(shell git rev-parse HEAD)
GITDATE ?= $(shell git show -s --format='%ct')
VERSION := v0.0.0

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.Version=$(VERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

inura-heartbeat:
	env GO111MODULE=on go build -v $(LDFLAGS) -o ./bin/inura-heartbeat ./cmd

clean:
	rm bin/inura-heartbeat

test:
	go test -v ./...

.PHONY: \
	clean \
	inura-heartbeat \
	test \
	lint
